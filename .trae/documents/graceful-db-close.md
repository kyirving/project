# 优雅关闭数据库连接

## 当前状态分析

`main.go` 已有 HTTP 服务优雅关闭的基础设施（信号监听 → `srv.Shutdown`），但缺少数据库连接的关闭逻辑：

- [database/mysql.go](file:///Users/wuhao/data/goland/components/project/components/database/mysql.go) — `Connect` 返回 `*gorm.DB`，但没有提供 `Close` 方法，调用方无法直接关闭底层连接池。
- [main.go](file:///Users/wuhao/data/goland/components/project/main.go#L54-L56) — 信号处理中只关闭了 HTTP Server，数据库连接未关闭就退出了。
- `println(db)` 是占位代码，db 尚未实际使用。

## 方案

### 1. 在 database 包添加 `Close` 函数

**文件：** `components/database/mysql.go`

**做什么：** 新增 `Close(db *gorm.DB) error` 函数，封装获取 `*sql.DB` → 调用 `Close()` 的逻辑。

**为什么：** 
- 调用方不需要知道 `gorm.DB → sql.DB` 的内部细节
- 统一关闭入口，后续可扩展（如 flush、metrics 清理等）

### 2. 在 main.go 信号处理中关闭数据库

**文件：** `main.go`

**做什么：**
- 在 `srv.Shutdown` 成功后，调用 `database.Close(db)`
- 同时清理占位代码 `println(db)`

**为什么：**
- 关闭顺序：先停 HTTP 再关 DB（避免关闭 DB 后还有请求进来）
- 利用已有的信号处理通道，不用额外引入复杂度

## 具体改动

### `components/database/mysql.go`

在 `Connect` 函数后新增 `Close` 函数：

```go
// Close 关闭数据库连接池
func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
```

### `main.go`

在 `srv.Shutdown` 之后，`logger.Sugar().Infof("Server exiting")` 之前插入数据库关闭逻辑：

```go
if err := srv.Shutdown(ctx); err != nil {
	logger.Sugar().Errorf("Server Shutdown: %s", err)
}

// 关闭数据库连接
if err := database.Close(db); err != nil {
	logger.Sugar().Errorf("Database Close: %s", err)
}

logger.Sugar().Infof("Server exiting")
```

同时移除 `println(db)` 这行占位代码。

## 验证方式

1. 编译通过：`go build ./...`
2. 启动服务，发送 SIGTERM，确认日志中输出 `Server exiting` 且无数据库关闭报错
