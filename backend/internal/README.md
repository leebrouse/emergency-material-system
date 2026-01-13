# 微服务启动框架

本目录包含了应急物资管理系统的所有微服务，每个服务都实现了完整的REST API和gRPC架构。

## 📁 服务架构

每个服务都遵循以下标准架构：

```
service/
├── main.go              # 服务启动入口
├── handler/             # REST API处理器
├── service/             # 业务逻辑层
├── repository/          # 数据访问层
├── model/               # 数据模型
├── rpc/                 # gRPC服务实现
└── go.mod               # Go模块配置
```

## 🚀 服务列表

### 1. Auth Service (认证服务)
- **REST API端口**: 8081
- **gRPC端口**: 9091
- **功能**: 用户登录、登出、令牌管理

```bash
cd internal/auth
go run main.go
```

### 2. Stock Service (物资库存服务)
- **REST API端口**: 8082
- **gRPC端口**: 9092
- **功能**: 物资管理、库存管理

```bash
cd internal/stock
go run main.go
```

### 3. Dispatch Service (调度服务)
- **REST API端口**: 8083
- **gRPC端口**: 9093
- **功能**: 需求申报、调度管理

```bash
cd internal/dispatch
go run main.go
```

### 4. Statistics Service (统计服务)
- **REST API端口**: 8084
- **功能**: 数据统计分析

```bash
cd internal/statistics
go run main.go
```

### 5. Logistics Service (物流服务)
- **REST API端口**: 8085
- **功能**: 物流追踪管理

```bash
cd internal/logistics
go run main.go
```

## 🔧 技术栈

- **Web框架**: Gin
- **RPC框架**: gRPC
- **数据库**: GORM + MySQL
- **配置管理**: Viper
- **日志**: Zap
- **链路追踪**: OpenTelemetry

## 📚 API文档

### REST API

每个服务的REST API都遵循以下模式：

```
GET  /health                    # 健康检查
GET  /api/v1/{service}/...      # 业务API
POST /api/v1/{service}/...      # 创建资源
PUT  /api/v1/{service}/...      # 更新资源
```

### gRPC API

gRPC服务使用Protocol Buffers定义接口，支持双向流式RPC。

## 🗄️ 数据库

每个服务使用独立的数据库连接，通过GORM进行数据操作。数据库迁移文件位于 `migrations/` 目录。

## ⚙️ 配置

服务配置通过 `internal/common/config/global.yaml` 统一管理，包括：

- 数据库连接
- 服务端口
- JWT密钥
- 日志级别

## 🔐 认证与授权

- 使用JWT进行身份认证
- 支持角色-based访问控制
- API密钥验证

## 📊 监控与日志

- 结构化日志输出
- 请求链路追踪
- 健康检查端点
- Prometheus指标收集

## 🚀 部署

### 本地开发

1. 启动MySQL数据库
2. 运行数据库迁移
3. 启动各个微服务

### Docker部署

使用 `deploy/docker-compose.yaml` 进行容器化部署。

## 🧪 测试

每个服务都包含单元测试和集成测试。运行测试：

```bash
go test ./...
```

## 📝 注意事项

1. 确保MySQL数据库已启动
2. 检查配置文件中的端口是否冲突
3. gRPC服务需要先于REST API启动
4. 注意JWT密钥的安全性
5. 生产环境请启用TLS
