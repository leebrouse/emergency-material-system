# 应急物资管理系统 - 后端服务

## 📋 项目简介

应急物资管理系统后端采用微服务架构，基于 Go 语言开发，提供物资管理、库存管理、需求调度、物流追踪和统计分析等核心功能。

## 🏗️ 项目结构

```
backend/
├── api/                          # API 定义（对外 & 对内）
│   ├── openapi/                  # OpenAPI / Swagger 定义
│   │   ├── auth.yaml             # 认证服务 API
│   │   ├── stock.yaml            # 物资库存服务 API
│   │   ├── dispatch.yaml         # 调度服务 API
│   │   └── statistics.yaml       # 统计服务 API
│   └── proto/                    # gRPC proto 定义
│       ├── auth.proto            # 认证服务 proto
│       ├── stock.proto           # 物资库存服务 proto
│       ├── dispatch.proto        # 调度服务 proto
│       └── logistics.proto       # 物流追踪服务 proto
│
├── cmd/                          # 各微服务启动入口
│   ├── gateway/                  # API Gateway（REST 统一入口）
│   │   └── main.go
│   ├── auth/                     # 认证与权限服务
│   │   └── main.go
│   ├── stock/                    # 物资与库存服务
│   │   └── main.go
│   ├── dispatch/                 # 调度与配送服务
│   │   └── main.go
│   ├── logistics/                # 物流追踪服务
│   │   └── main.go
│   └── statistics/               # 统计分析服务
│       └── main.go
│
├── internal/                     # 核心业务代码（不对外暴露）
│   ├── common/                   # 通用基础设施
│   │   ├── config/               # 配置加载（Viper）
│   │   ├── database/             # 数据库连接（MySQL / GORM）
│   │   ├── logging/              # 日志（zap / logrus）
│   │   ├── middleware/           # Gin 中间件
│   │   ├── metrics/              # Prometheus 指标
│   │   ├── tracing/              # 链路追踪（OpenTelemetry）
│   │   ├── errors/               # 统一错误码
│   │   └── utils/                # 工具函数
│   │
│   ├── auth/                     # 【模块 1】认证与权限
│   │   ├── handler/              # HTTP Handler（Gin）
│   │   ├── service/              # 业务逻辑
│   │   ├── repository/           # 数据访问
│   │   ├── model/                # 数据模型
│   │   └── rpc/                  # gRPC Server
│   │
│   ├── stock/                    # 【模块 2】物资 & 库存
│   │   ├── handler/
│   │   ├── service/
│   │   ├── repository/
│   │   ├── model/
│   │   └── rpc/
│   │
│   ├── dispatch/                 # 【模块 3】需求申报 & 调度
│   │   ├── handler/
│   │   ├── service/
│   │   ├── repository/
│   │   ├── model/
│   │   └── rpc/
│   │
│   ├── logistics/                # 【模块 4】物流轨迹
│   │   ├── handler/
│   │   ├── service/
│   │   ├── repository/
│   │   ├── model/
│   │   └── rpc/
│   │
│   └── statistics/               # 【模块 5】统计分析 & 报表
│       ├── handler/
│       ├── service/
│       ├── repository/
│       └── model/
│
├── deploy/                       # 部署相关
│   ├── docker-compose.yaml       # Docker Compose 配置
│   ├── Dockerfile.gateway        # Gateway 服务 Dockerfile
│   ├── Dockerfile.auth           # Auth 服务 Dockerfile
│   ├── Dockerfile.stock          # Stock 服务 Dockerfile
│   └── k8s/                      # Kubernetes 部署配置（可选）
│
├── migrations/                   # 数据库迁移
│   ├── 001_init.sql
│   ├── 002_stock.sql
│   └── 003_dispatch.sql
│
├── script/                       # 脚本工具
│   ├── genopenapi.sh            # 生成 OpenAPI 文档
│   ├── genproto.sh              # 生成 gRPC 代码
│   ├── lint.sh                  # 代码检查
│   └── migrate.sh               # 数据库迁移
│
├── docs/                         # 文档
│   ├── architecture.md           # 系统架构说明
│   ├── database.md               # ER 设计
│   ├── api_docs.md               # API 说明
│   └── deployment.md             # 部署文档
│
├── go.work                       # Go Workspace 配置
└── README.md                     # 本文件
```

## 🚀 快速开始

### 前置要求

- Go 1.24+ 
- MySQL 8.0+
- Docker & Docker Compose（可选）

### 本地开发

1. **克隆项目**
```bash
git clone <repository-url>
cd emergency-material-system/backend
```

2. **启动依赖服务**
```bash
cd deploy
docker-compose up -d mysql
```

3. **运行数据库迁移**
```bash
./script/migrate.sh
```

4. **启动服务**
```bash
# 启动 Gateway
cd cmd/gateway
go run main.go

# 启动其他服务（新终端窗口）
cd cmd/auth
go run main.go
```

### 使用 Docker Compose

```bash
cd deploy
docker-compose up -d
```

## 🏛️ 架构设计

### 微服务架构

系统采用微服务架构，包含以下服务：

- **Gateway**: API 网关，统一对外提供 REST API
- **Auth**: 认证与权限管理服务
- **Stock**: 物资与库存管理服务
- **Dispatch**: 需求申报与调度服务
- **Logistics**: 物流追踪服务
- **Statistics**: 统计分析服务

### 技术栈

- **Web 框架**: Gin
- **数据库**: MySQL 8.0 + GORM
- **RPC**: gRPC
- **配置管理**: Viper
- **日志**: zap / logrus
- **监控**: Prometheus
- **链路追踪**: OpenTelemetry
- **API 文档**: OpenAPI 3.0 / Swagger

### 通信方式

- **对外**: REST API（通过 Gateway）
- **服务间**: gRPC

## 📝 API 文档

### OpenAPI 文档

各服务的 OpenAPI 定义位于 `api/openapi/` 目录：

- `auth.yaml` - 认证服务 API
- `stock.yaml` - 物资库存服务 API
- `dispatch.yaml` - 调度服务 API
- `statistics.yaml` - 统计服务 API

### gRPC 定义

各服务的 gRPC proto 定义位于 `api/proto/` 目录。

### 接口快速参考

| 服务 | REST 端口 | gRPC 端口 | 主要职责 |
| :--- | :--- | :--- | :--- |
| **Auth** | 8081 | 9091 | 身份认证、JWT 鉴权、角色管理 |
| **Stock** | 8082 | 9092 | 物资元数据、实时库存、入库/出库/调拨 |
| **Dispatch** | 8083 | 9093 | 需求申请、审核、库存分配建议、调度任务创建 |
| **Statistics**| 8084 | - | 数据聚合、多角度统计报表、趋势分析 |
| **Logistics** | 8085 | 9095 | 物流追踪记录、轨迹节点实时上报 |

详细 API 接口列表请参考 [docs/api_docs.md](./docs/api_docs.md)。

### 生成文档

```bash
# 生成 OpenAPI 文档
./script/genopenapi.sh

# 生成 gRPC 代码
./script/genproto.sh
```

## 🗄️ 数据库

### 数据库迁移

数据库迁移文件位于 `migrations/` 目录。

运行迁移：
```bash
./script/migrate.sh
```

## 🛠️ 开发工具

### 代码检查

```bash
./script/lint.sh
```

### 生成代码

```bash
# 生成 gRPC 代码
./script/genproto.sh

# 生成 OpenAPI 文档
./script/genopenapi.sh
```

## 📦 部署

### Docker 部署

详见 `deploy/docker-compose.yaml`

### Kubernetes 部署

K8s 配置文件位于 `deploy/k8s/` 目录。

## 📚 文档

详细文档请参考 `docs/` 目录：

- `architecture.md` - 系统架构说明
- `database.md` - 数据库设计
- `api_docs.md` - API 详细说明 (综合梳理)
- `deployment.md` - 部署指南

## 🔧 配置

配置文件位于 `internal/common/config/global.yaml`

主要配置项：
- 数据库连接
- 服务端口
- 日志级别
- JWT 密钥
- 服务发现配置

## 🤝 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 📄 许可证

[添加许可证信息]

## 👥 团队

[添加团队信息]

## 📞 联系方式

[添加联系方式]

---

## ⚠️ 重要说明

### Go Workspace 配置

本项目采用 Go Workspace 管理多模块，使用 `go.work` 文件进行开发。

**注意**：
- `go.work` 和 `go.work.sum` 文件已添加到 `.gitignore`，**不会被 Git 追踪**
- 原因：`go.work` 是本地开发配置，不同开发者的工作区配置可能不同
- 每个开发者需要根据自己的环境创建 `go.work` 文件
- 如果团队需要统一配置，也可以将 `go.work` 提交到仓库（根据团队约定）

### 创建 go.work 文件

如果项目中没有 `go.work` 文件，可以运行以下命令创建：

```bash
cd backend
go work init
go work use ./internal/common
go work use ./internal/auth
go work use ./internal/stock
go work use ./internal/dispatch
go work use ./internal/logistics
go work use ./internal/statistics

# 任意一次构建或整理
go work sync

```

或者手动创建 `go.work` 文件，参考项目根目录的示例。

