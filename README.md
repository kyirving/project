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
| 监控         | Prometheus + Grafana                                                 |

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
│   ├── logger/              # Zap 日志（异步缓冲、按天切割）
│   └── metrics/             # Prometheus 指标定义（Counter、Histogram）
├── internal/                # 业务逻辑
│   ├── model/               # 数据模型（User、UserIndex）
│   ├── repository/          # 数据访问层（分表路由、事务）
│   ├── service/             # 业务逻辑层
│   ├── handler/             # Gin HTTP 处理器
│   └── middleware/          # 中间件（日志、异常、超时、限流、CORS、Metrics）
├── router/                  # 路由注册与依赖注入
├── utils/                   # 工具包
│   ├── biz_error/           # 业务错误定义
│   ├── file/                # 文件工具
│   └── request/             # 请求结构体
├── deploy/                  # Kubernetes 部署配置 + Helm Chart
│   ├── configMap.yaml       # 应用配置
│   ├── deployment.yaml      # Deployment 部署
│   ├── service.yaml         # Service (NodePort)
│   └── mychart/             # Helm Chart (含 HPA、探针、资源限制)
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

## 监控（Prometheus + Grafana）

### 内置端点

| 端点       | 用途           | 检查内容                     |
| ---------- | -------------- | ---------------------------- |
| `/health`  | Liveness Probe | 进程可响应 HTTP 即可          |
| `/ready`   | Readiness Probe | DB Ping（失败则摘除流量）     |
| `/metrics` | Prometheus 抓取 | Go runtime + 自定义 HTTP 指标 |

### 采集的指标

- `http_requests_total` — 按 method / path / status 的请求计数
- `http_request_duration_seconds` — 请求耗时分布（Histogram）
- `go_*` / `process_*` — Go runtime 和进程指标（自动采集）

### 本地开发（minikube）

```bash
# 1. 宿主机拉镜像并 load 进 minikube（解决国内拉不下来的问题）
docker pull prom/prometheus
# 使用阿里云镜像，解决国内拉不下来的问题
docker pull registry.cn-hangzhou.aliyuncs.com/chenby/kube-state-metrics:v2.19.0
docker tag registry.cn-hangzhou.aliyuncs.com/chenby/kube-state-metrics:v2.19.0 registry.k8s.io/kube-state-metrics/kube-state-metrics:v2.19.1
docker pull grafana/grafana:latest
minikube image load prom/prometheus
minikube image load grafana/grafana:latest
minikube image load registry.k8s.io/kube-state-metrics/kube-state-metrics:v2.19.1


# 2. 安装 Prometheus（轻量版，去掉不需要的组件）
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install prometheus prometheus-community/prometheus \
  --namespace monitoring --create-namespace \
  --set alertmanager.enabled=false \
  --set pushgateway.enabled=false \
  --set server.persistentVolume.enabled=false \
  --set kube-state-metrics.enabled=false \
  --set prometheus-node-exporter.enabled=false \
  --set server.image.pullPolicy=IfNotPresent

# 3. 安装 Grafana
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
helm install grafana grafana/grafana \
  --namespace monitoring \
  --set image.pullPolicy=IfNotPresent \
  --set adminPassword=admin

# 4. 部署你的应用（确认 values.yaml 中 podAnnotations 已配好）
helm install app ./deploy/mychart -n default

# 5. 端口转发
kubectl port-forward -n monitoring svc/prometheus-server 9090:80 &
kubectl port-forward -n monitoring svc/grafana 3000:80 &
```

### Grafana 配置

1. 打开 `http://localhost:3000`，账号 `admin`，密码 `admin`
2. Administration → Data sources → Add data source → Prometheus
3. URL 填 `http://prometheus-server.monitoring:80`，Save & test
4. Dashboards → Import → ID 填 `11074`（Go 应用模板），选择数据源即可

### PromQL 示例

```promql
# QPS
rate(http_requests_total[1m])

# P99 延迟
histogram_quantile(0.99, rate(http_request_duration_seconds_bucket[1m]))

# 错误率
sum(rate(http_requests_total{status="500"}[1m])) / sum(rate(http_requests_total[1m]))
```

### 生产环境

生产环境推荐 `kube-prometheus-stack`（集成 Prometheus Operator + Grafana + 预置仪表盘），只需在 Helm values 中配置 `podAnnotations`：

```yaml
podAnnotations:
  prometheus.io/scrape: "true"
  prometheus.io/port: "8080"
  prometheus.io/path: "/metrics"
```
