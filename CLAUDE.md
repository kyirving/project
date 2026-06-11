# CLAUDE.md

此文件为 Claude Code（claude.ai/code）在本仓库中工作时提供指导。

## 构建与运行

```bash
# 运行应用
go run main.go --config_dir ./config --config_file config.yaml

# 构建
go build -o bin/app main.go

# 运行测试（暂无）
go test ./...
```

## 架构

Go 1.26 项目（模块名 `app`），使用 **Gin** 作为 HTTP 框架，**GORM** 作为 MySQL ORM，**Zap** 记录结构化日志，**Viper** 加载 YAML 配置。

采用**整洁架构**，在 `router/` 中通过依赖注入串联三层：

```
handler（HTTP）→ service（业务逻辑）→ repository（数据访问）
```

### 目录职责

- **`main.go`** — 入口：加载配置，创建 logger、DB、idgen、jwt 等基础设施，注册路由，启动 HTTP 服务并优雅关闭。
- **`config/`** — 基于 Viper 加载配置。CLI 参数 `--config_dir` 和 `--config_file` 指定 YAML 文件路径。结构体标签：顶层字段使用 `json`，嵌套结构体使用 `mapstructure`。
- **`components/`** — 基础设施（不依赖 `internal/`）：
  - `database/mysql.go` — GORM 连接 + 连接池配置。
  - `logger/logger.go` — Zap + Lumberjack 日志切割，异步缓冲写入，JSON 格式，ISO8601 时间戳，时间戳键名为 `@timestamp`。
  - `idgen/idgen.go` — 雪花 ID 生成器，定义 `Generator` 接口，默认使用 `bwmarrin/snowflake` 实现。在 `main.go` 中全局初始化一次，后续 `service` 层直接调用 `NextID()`。
  - `jwt/jwt.go` — JWT 签发与解析（HS256），包含 `CustomClaims` 和 `JwtManager`。
- **`internal/`** — 业务逻辑：
  - `model/` — 纯数据结构体（对应数据库表），不包含任何业务逻辑或外部依赖。`UserIndex` 是用户名→用户 ID+分片的索引表。
  - `repository/` — 数据访问层，封装 `*gorm.DB`。负责分表路由、事务、SQL 执行。**不要将请求 struct 或 service DTO 传给 repository**，始终传 `model.User` 或原始值。
  - `service/` — 业务逻辑层，通过构造函数注入 repository 和 components。负责调用 idgen 生成 ID、组装 model、调用 jwt 签发 token、业务校验。
  - `handler/` — Gin 处理器。`BaseHandler` 提供标准化 JSON 响应 `{code, message, data}`。各 Handler 通过嵌入 `BaseHandler` 继承方法。
  - `middleware/` — Gin 中间件（logger、exception、timeout、limiter、CORS）。
- **`router/`** — 路由定义。`router.go` 创建 Gin 引擎并挂载中间件，将路由组注册委托给 `user.go` 等文件。在此完成依赖注入：创建 components → repo → service → handler 链。
- **`utils/`** — 跨层共用的工具：
  - `biz_error/` — `BizError` 类型，用于业务错误（用户存在、密码短等），service 返回、handler 判断类型后返回自定义消息给前端。
  - `file/` — DirExists、FileExists。
  - `request/` — HTTP 请求结构体。

## 关键约定

### 分层传参

```
handler:  c.ShouldBindJSON(&req)         → 绑定带 json/binding tag 的请求 struct
handler:  svc.DoSomething(req.Field1, ...) → 逐字段传给 service（不要传请求 struct）
service:  组装 model.User{}              → 转成数据库模型
repo:     db.Create(user)                → 只接收 model
```

- handler 的请求 struct（含 `json` 和 `binding` tag）**不能**传到 service 或 repository 层。
- service 层的 DTO **不能**传到 repository 层。
- model 只做数据结构定义，不参与分表计算、ID 生成等业务逻辑。

### 雪花 ID 与分表

- 雪花 ID 在 `service` 层通过 `idgen.NextID()` 生成。
- 16 张分表：`user_0` ~ `user_15`。
- 分表算法使用 XOR 折叠（因为雪花 ID 低 12 位是序列号，跨毫秒归零，直接 `% 16` 分布不均）：
  ```go
  func TableName(id uint64) string {
      return fmt.Sprintf("user_%d", (id^(id>>32))%16)
  }
  ```
- 分表逻辑全部在 `repository` 层内部，service 无感。
- `model.UserIndex` 是索引表（固定表名 `user_index`），存 `UserID`、`UserName`、`IndexValue`（分片后缀），用于按用户名快速定位数据在哪张分表。

### 错误处理

- `utils/biz_error.BizError` 用于业务错误（Code + Message），service 返回，handler 用 `errors.As` 判断。
- `BaseHandler` 提供三个响应方法：
  - `Success(c, data)` — 成功，code=0。
  - `Error(c, code)` — 系统错误，使用预设的 `respMsg` 映射，不暴露细节给前端。
  - `BizError(c, bizErr)` — 业务错误，将自定义 code 和 message 返回前端。
- handler 中的典型错误分支：
  ```go
  result, err := h.svc.DoSomething(...)
  if err == nil {
      h.Success(c, result)
      return
  }
  var bizErr *biz_error.BizError
  if errors.As(err, &bizErr) {
      h.BizError(c, bizErr)   // 业务错误 → 透传自定义 message
  } else {
      h.logger.Error.Error("xxx", zap.Error(err))
      h.Error(c, INTERNAL_ERROR)  // 系统错误 → 只记日志，返回通用错误
  }
  ```

### 配置与注入

- 配置结构体字段：顶层使用 `json` 标签，嵌套结构体使用 `mapstructure` 标签。
- Logger 使用 `zap.Sugar()` 进行格式化输出。生产日志为 JSON 格式，包含 `@timestamp`、`level`、`caller`、`msg` 等字段。
- 数据库连接池：最大打开连接数 25，最大空闲连接数 10，连接最大存活时间 5 分钟。
- 基础设施（db、idgen、jwt、logger）在 `main.go` 中创建，通过 `router` 的依赖注入链传入 service。
