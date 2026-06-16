# Project

基于 Go 语言构建的用户认证服务，采用整洁架构（Clean Architecture），提供用户注册、登录及 JWT 鉴权能力，内置雪花 ID 生成与分表路由，支持 Kubernetes 部署。

## 技术栈

| 类别         | 组件                                                                 |
| ------------ | -------------------------------------------------------------------- |
| HTTP 框架    | [Gin](https://github.com/gin-gonic/gin)                              |
| ORM          | [GORM](https://github.com/go-gorm/gorm) + MySQL                      |
| 配置管理     | [Viper](https://github.com/spf13/viper)                              |
| 日志         | [Zap](https://github.com/uber-go/zap) + [Lumberjack](https://github.com/natefinch/lumberjack) |
| ID 生成      | [Snowflake](https://github.com/bwmarrin/snowflake)                   |
| JWT          | [golang-jwt](https://github.com/golang-jwt/jwt)                      |
| 容器编排     | Kubernetes + Helm                                                    |

## 项目结构

```
.
├── main.go                  # 入口：加载配置、初始化组件、启动服务并优雅关闭
├── config/                  # Viper 配置加载
│   ├── config.go
│   └── config.yaml          # 应用配置文件
├── components/              # 基础设施组件（不依赖 internal）
│   ├── database/            # GORM 数据库连接 + 建表 SQL
│   ├── idgen/               # 雪花 ID 生成器
│   ├── jwt/                 # JWT 签发与解析
│   └── logger/              # Zap 日志（异步缓冲、按天切割）
├── internal/                # 业务逻辑
│   ├── model/               # 数据模型（User、UserIndex）
│   ├── repository/          # 数据访问层（分表路由、事务）
│   ├── service/             # 业务逻辑层
│   ├── handler/             # Gin HTTP 处理器
│   └── middleware/          # 中间件（日志、异常、超时、限流、CORS）
├── router/                  # 路由注册与依赖注入
├── utils/                   # 工具包
│   ├── biz_error/           # 业务错误定义
│   ├── file/                # 文件工具
│   └── request/             # 请求结构体
├── deploy/                  # Kubernetes 部署配置 + Helm Chart
├── Dockerfile               # 多阶段构建
└── CLAUDE.md                # Claude Code 工作指引
```

## 架构分层

```
handler (HTTP) → service (业务逻辑) → repository (数据访问)
```

**依赖注入链**：`main.go` 创建 DB、Logger、IDGen 等基础设施 → `router` 组装 handler → service → repository 链 → 注册路由 → 启动服务。

### 关键约定

- **请求结构体**（`json` + `binding` tag）只在 handler 层使用，不传入 service 或 repository。
- **model** 只做数据结构定义，不含业务逻辑。
- **repository** 封装分表路由和事务，service 层无感。
- 业务错误使用 `BizError`，service 返回，handler 判断后返回对应消息给前端。
- JSON 响应统一格式：`{ code, message, data }`。

## 快速开始

### 环境要求

- Go 1.26+
- MySQL 8.0+

### 1. 初始化数据库

执行 `components/database/app.sql` 建表语句，创建 `go_user` 数据库及 `user_index`、`user_0` ~ `user_15` 共 17 张表。

### 2. 修改配置

编辑 `config/config.yaml`，填入实际的数据库连接信息：

```yaml
DB:
  HOST: your-mysql-host
  PORT: 3306
  USER: your-username
  PASSWORD: your-password
  DB_NAME: go_user
```

### 3. 运行

```bash
# 直接运行
go run main.go --config_dir ./config --config_file config.yaml

# 构建
go build -o bin/app main.go
./bin/app --config_dir ./config --config_file config.yaml
```

服务默认监听 `:8080`，可通过配置文件中的 `APP.PORT` 修改。

## API 接口

### 用户注册

```
POST /api/user/register
Content-Type: application/json

{
    "username": "foo",
    "password": "123456"
}
```

### 用户登录

```
POST /api/user/login
Content-Type: application/json

{
    "username": "foo",
    "password": "123456"
}
```

## 数据分片

使用雪花算法生成全局唯一用户 ID，通过 XOR 折叠算法将用户数据均匀分布到 16 张分表（`user_0` ~ `user_15`）：

```go
func TableName(id uint64) string {
    return fmt.Sprintf("user_%d", (id^(id>>32))%16)
}
```

`user_index` 表存储用户名到用户 ID 及分片后缀的映射，实现按用户名快速定位数据所在分表。

## Docker 部署

```bash
# 构建镜像
docker build -t app:latest .

# 运行容器
docker run -d -p 8080:8080 \
  -v $(pwd)/config:/app/config \
  app:latest
```

## Kubernetes 部署

使用 Helm Chart 部署到 Kubernetes 集群：

```bash
# 安装
helm install app ./deploy/mychart -n <namespace> --create-namespace

# 升级
helm upgrade app ./deploy/mychart -n <namespace>
```

Helm Chart 包含 HPA 自动伸缩、健康检查探针、资源限制等生产级配置。
