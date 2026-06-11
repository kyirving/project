# 用户分表设计方案

## 当前状态

项目处于骨架阶段，没有任何 User Model、无建表逻辑、无分表代码。使用 Gin + GORM + MySQL，遵循 Handler → Service → Repository 三层架构。

## 设计决策

| 决策项 | 选择 |
|--------|------|
| 分表策略 | Hash 取模 |
| 分表数量 | **256 张（预分表）**，user_0 ~ user_255 |
| 分片键 | `user_id`（雪花算法） |
| 实现方式 | 手动封装 Repository 层 |

**为什么选 256 张（预分表）：**

核心思路是 **用空间换免迁移**。直接一步到位建 256 张表，省去后续扩容的痛苦。

- MySQL 空表开销极小（仅元数据 + .ibd 文件），256 张空表几乎不占磁盘
- 假设每张表存 1000 万用户，256 × 1000万 = 25.6 亿，基本不用再扩
- 路由算法不变，**零数据迁移**
- 退一步说，256 张还不够 → 说明体量该上分库中间件或 TiDB 了，不是分表能解决的

## 方案概览

**核心思路：`user_index` 充当路由表**

```
注册：
  username + password
       │
       ▼
  生成 snowflake user_id
       │
       ├──▶ user_index 表写入 {username → user_id}   ← 单张表，不改动
       │
       └──▶ hash(user_id)%256 → user_N 表写入完整数据

登录：
  username + password
       │
       ▼
  查 user_index 表拿到 user_id
       │
       └──▶ hash(user_id)%256 → user_N 表查密码，校验
```

## 具体实现

### 1. 新增 `internal/model/user.go` — User Model

```go
package model

import "time"

type User struct {
	ID        uint64    `gorm:"column:id;primaryKey" json:"id"`
	Username  string    `gorm:"column:username;type:varchar(64);uniqueIndex;not null" json:"username"`
	Password  string    `gorm:"column:password;type:varchar(128);not null" json:"-"`
	Email     string    `gorm:"column:email;type:varchar(128)" json:"email"`
	Phone     string    `gorm:"column:phone;type:varchar(20)" json:"phone"`
	Status    int8      `gorm:"column:status;default:1" json:"status"` // 1:正常 0:禁用
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// UserIndex 用户名→用户ID路由表，不受分表影响，始终单表
type UserIndex struct {
	ID       uint64 `gorm:"column:id;primaryKey;autoIncrement"`
	Username string `gorm:"column:username;type:varchar(64);uniqueIndex;not null"`
	UserID   uint64 `gorm:"column:user_id;not null;index"`
}
```

### 2. 新增 `internal/repository/user_index.go` — 索引表 Repository

```go
package repository

import (
	"app/internal/model"
	"errors"

	"gorm.io/gorm"
)

type UserIndexRepository struct {
	db *gorm.DB
}

func NewUserIndexRepository(db *gorm.DB) *UserIndexRepository {
	return &UserIndexRepository{db: db}
}

// Resolve 根据 username 获取 user_id（登录时用）
func (r *UserIndexRepository) Resolve(username string) (uint64, error) {
	var idx model.UserIndex
	err := r.db.Where("username = ?", username).First(&idx).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil // 用户不存在
	}
	return idx.UserID, err
}

// Bind 注册时写入映射关系
func (r *UserIndexRepository) Bind(username string, userID uint64) error {
	return r.db.Create(&model.UserIndex{
		Username: username,
		UserID:   userID,
	}).Error
}
```

### 3. 新增 `internal/repository/shard.go` — 分表路由器

```go
package repository

import (
	"fmt"
	"hash/crc32"
)

const (
	UserTablePrefix = "user"
	UserShardCount  = 256
)

// UserTableName 根据 userID 计算分表名
func UserTableName(userID uint64) string {
	idx := crc32.ChecksumIEEE([]byte(fmt.Sprintf("%d", userID))) % UserShardCount
	return fmt.Sprintf("%s_%d", UserTablePrefix, idx)
}
```

**为什么用 CRC32 而不是 `user_id % 16`：** 
- 如果 user_id 使用自增主键，直接取模会导致连续 ID 落在同一张表（因为模数太小），CRC32 打散效果更好
- 如果后续使用雪花算法，直接取模也可行，但 CRC32 更通用

### 4. 新增 `internal/repository/user.go` — UserRepository

```go
package repository

import (
	"app/internal/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create 插入用户，自动路由到对应分表
func (r *UserRepository) Create(user *model.User) error {
	return r.db.Table(UserTableName(user.ID)).Create(user).Error
}

// FindByID 根据 ID 查询
func (r *UserRepository) FindByID(userID uint64) (*model.User, error) {
	var user model.User
	err := r.db.Table(UserTableName(userID)).Where("id = ?", userID).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// Update 更新用户
func (r *UserRepository) Update(user *model.User) error {
	return r.db.Table(UserTableName(user.ID)).
		Where("id = ?", user.ID).Updates(user).Error
}

// Delete 软删除（改 status）
func (r *UserRepository) Delete(userID uint64) error {
	return r.db.Table(UserTableName(userID)).
		Where("id = ?", userID).Update("status", 0).Error
}
```

### 5. 新增建表 SQL — `migrations/001_create_user_tables.sql`

生成 256 张结构相同的表（user_0 ~ user_255），通过脚本批量生成：

```sql
-- user_0（其余 user_1 ~ user_255 结构一致）
CREATE TABLE IF NOT EXISTS `user_0` (
  `id` BIGINT UNSIGNED NOT NULL,
  `username` VARCHAR(64) NOT NULL,
  `password` VARCHAR(128) NOT NULL,
  `email` VARCHAR(128) DEFAULT '',
  `phone` VARCHAR(20) DEFAULT '',
  `status` TINYINT DEFAULT 1,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- 重复 user_1 ~ user_15 ...
```

### 5. 更新 `config/config.go` — 增加分表配置（可选）

```go
type DbConfig struct {
	// ... 现有字段 ...
	UserShardCount  int `json:"user_shard_count"`   // 默认 256
	UserShardPrefix string `json:"user_shard_prefix"` // 默认 "user"
}
```

这样可以后续通过配置调整分表数量，Repository 中的 `UserTableName` 读取配置值即可。

### 6. 改造 `internal/service/oauth.go` — 整合

将 Login/Register 改为调用 `UserRepository`，不再直接用空的 `OAuthRepository`。

## 扩容策略

### 为什么预分表 256 张基本消除了扩容需求

| 每表容量 | 256 张可承载用户数 | 评估 |
|----------|-------------------|------|
| 100 万 | 2.56 亿 | 绝大多数项目天花板 |
| 500 万 | 12.8 亿 | 头部互联网级别 |
| 1000 万 | 25.6 亿 | 超出 MySQL 单表建议上限 |

MySQL 单表建议控制在 1000 万以内，所以 256 张表覆盖的用户量已经远超单库承载能力。**真到这个量级，瓶颈是数据库实例本身（连接数、IO、CPU），需要上分库或 TiDB，而不是继续分表。**

### 万一还是不够：256 → 512 扩容流程

即使走了预分表路线，如果真的需要扩，步骤如下：

```
扩容前：hash(user_id) % 256  →  user_0 ~ user_255
扩容后：hash(user_id) % 512  →  user_0 ~ user_511
```

**四步平滑扩容：**

```
步骤 1 — 双写
  ┌──────────┐
  │ 新请求   │──写──▶ 新路由表 (user_0~511)
  │          │──写──▶ 旧路由表 (user_0~255)   ← 保证旧路由仍有全量数据
  └──────────┘

步骤 2 — 迁移历史数据
  遍历 user_0 ~ user_255 所有记录
  → 按 hash(id)%512 重新计算目标表
  → INSERT INTO user_N ... ON DUPLICATE KEY UPDATE

步骤 3 — 灰度切读
  读请求 → 先查新表 → miss 再查旧表 → 异步补写新表
  逐步增加新路由比例：10% → 50% → 100%

步骤 4 — 完成
  全量读切新路由 → 停止双写 → 清理旧路由逻辑
```

> 这个流程需要运维脚本配合，项目初期不用实现，了解即可。**结论：256 张表基本用不到扩容。**

## 注意事项

1. **跨表查询问题：** 按 username 查用户时无法确定在哪张表。256 张表遍历成本太高 → **强烈建议额外维护一张 `user_index` 表**（username → user_id 映射），注册时写入，登录时先查索引拿到 user_id 再定位分表。或者用 Redis 缓存 username→user_id。
2. **ID 生成策略：** 推荐使用雪花算法（如 `github.com/bwmarrin/snowflake`），提前生成 user_id 再分表写入。
3. **事务：** 单表操作不受影响。跨表操作（如改 username 同时维护索引表）需注意一致性，但本项目暂不需要。

## 文件变更清单

| 操作 | 文件 | 说明 |
|------|------|------|
| 新增 | `internal/model/user.go` | User 结构体 |
| 新增 | `internal/repository/shard.go` | 分表路由计算（256 张） |
| 新增 | `internal/repository/user.go` | UserRepository CRUD |
| 新增 | `migrations/001_create_user_tables.sql` | 建表脚本（user_0 ~ user_255） |
| 新增 | `migrations/002_create_user_index.sql` | username → user_id 索引表 |
| 修改 | `config/config.go` | 增加分表配置项 |
| 修改 | `internal/service/oauth.go` | 改用 UserRepository |
| 修改 | `internal/handler/oauth.go` | 适配新 Service |
| 修改 | `go.mod` | 加 snowflake 依赖 |

## 验证方式

1. 写单元测试验证 `UserTableName` 对不同 ID 的路由结果分布于 0~15
2. `go build ./...` 编译通过
3. 启动服务，调用注册接口，确认数据写入正确分表
