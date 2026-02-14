# 应急物资管理系统 API 文档

本文档整理了应急物资管理系统各微服务的 API 接口。

## 通用信息

- **基础路径**: `/api/v1`
- **所有服务均包含健康检查接口**: `GET /health`

---

## 1. 身份认证服务 (Auth Service)
- **端口**: REST `8081` / gRPC `9091`
- **基础路径**: `http://localhost:8081/api/v1`

| 接口 | 方法 | 描述 |
| :--- | :--- | :--- |
| `/auth/register` | POST | 用户注册 |
| `/auth/login` | POST | 用户登录，返回 Access Token |
| `/auth/logout` | POST | 用户登出 |
| `/auth/refresh` | POST | 刷新 Token |

### 接口示例

#### 用户注册
**请求:**
```bash
curl -X POST http://localhost:8081/api/v1/auth/register \
-H "Content-Type: application/json" \
-d '{
  "username": "admin",
  "password": "password123",
  "email": "admin@example.com",
  "phone": "13800138000",
  "roles": ["admin"]
}'
```
**响应 (201 Created):**
```json
{
  "status": "ok",
  "message": "user registered successfully"
}
```

#### 用户登录
**请求:**
```bash
curl -X POST http://localhost:8081/api/v1/auth/login \
-H "Content-Type: application/json" \
-d '{
  "username": "admin",
  "password": "123456"
}'
```
**响应 (200 OK):**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expires_in": 86400,
  "token_type": "Bearer"
}
```

---

## 2. 物资库存服务 (Stock Service)
- **端口**: REST `8082` / gRPC `9092`
- **基础路径**: `http://localhost:8082/api/v1`

| 接口 | 方法 | 描述 |
| :--- | :--- | :--- |
| `/stock/materials` | GET | 获取物资列表 |
| `/stock/materials` | POST | 创建新物资 |
| `/stock/materials/:id` | GET | 获取物资详情 |
| `/stock/inventory` | GET | 获取库存列表 |

### 接口示例

#### 获取物资列表
**请求:**
```bash
curl -X GET "http://localhost:8082/api/v1/stock/materials?page=1&page_size=10&search=口罩" \
-H "Authorization: Bearer <your_token>"
```

#### 创建物资
**请求:**
```bash
curl -X POST http://localhost:8082/api/v1/stock/materials \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_token>" \
-d '{
  "name": "医用外科口罩",
  "category": "防护用品",
  "specs": "50只/盒",
  "unit": "盒",
  "description": "三层过滤，符合YY/T 0969标准"
}'
```

---

## 3. 调度指挥服务 (Dispatch Service)
- **端口**: REST `8083` / gRPC `9093`
- **基础路径**: `http://localhost:8083/api/v1`

| 接口 | 方法 | 描述 |
| :--- | :--- | :--- |
| `/dispatch/requests` | POST | 创建物资需求申报 |
| `/dispatch/requests/:id/audit` | POST | 审核需求申报 |

### 接口示例

#### 创建需求申报
**请求:**
```bash
curl -X POST http://localhost:8083/api/v1/dispatch/requests \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_token>" \
-d '{
  "material_id": 1,
  "quantity": 1000,
  "urgency_level": "L1",
  "target_area": "武汉市中心医院",
  "description": "急需医疗防护物资"
}'
```

#### 审核需求申报
**请求:**
```bash
curl -X POST http://localhost:8083/api/v1/dispatch/requests/1/audit \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_token>" \
-d '{
  "action": "approve",
  "remark": "审核通过，立即安排调拨"
}'
```

---

## 4. 物流追踪服务 (Logistics Service)
- **端口**: REST `8085` / gRPC `9095`
- **基础路径**: `http://localhost:8085/api/v1`

| 接口 | 方法 | 描述 |
| :--- | :--- | :--- |
| `/logistics/tracking` | POST | 创建物流追踪记录 |
| `/logistics/tracking/:id` | GET | 获取物流追踪信息及轨迹节点 |
| `/logistics/tracking/:id` | PUT | 更新物流追踪状态 (如 `delivering`, `completed`) |
| `/logistics/tracking/:id/nodes` | POST | 记录物流轨迹节点 (位置、坐标、时间) |

### 接口示例

#### 创建物流追踪
**请求:**
```bash
curl -X POST http://localhost:8085/api/v1/logistics/tracking \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_token>" \
-d '{
  "task_id": 1,
  "carrier": "顺丰速运",
  "tracking_number": "SF1234567890",
  "status": "shipped"
}'
```

---

## 5. 统计分析服务 (Statistics Service)
- **端口**: REST `8084`
- **基础路径**: `http://localhost:8084/api/v1`

| 接口 | 方法 | 描述 |
| :--- | :--- | :--- |
| `/statistics/summary` | GET | 获取核心指标汇总 (总物资、总任务等) |
| `/statistics/reports` | GET | 物资分类统计报表 |
| `/statistics/trends` | GET | 物资需求/调度趋势分析 |

### 接口示例

#### 获取核心指标汇总
**请求:**
```bash
curl -X GET http://localhost:8084/api/v1/statistics/summary \
-H "Authorization: Bearer <your_token>"
```
**响应 (200 OK):**
```json
{
  "total_materials": 150,
  "total_stock": 25000,
  "active_dispatch_tasks": 12,
  "completed_dispatch_tasks": 156,
  "total_logistics_records": 168
}
```

#### 获取物资分类统计报表
**请求:**
```bash
curl -X GET http://localhost:8084/api/v1/statistics/reports \
-H "Authorization: Bearer <your_token>"
```
**响应 (200 OK):**
```json
{
  "categories": [
    { "name": "防护用品", "count": 45, "percentage": 30.0 },
    { "name": "医疗器械", "count": 30, "percentage": 20.0 },
    { "name": "消杀用品", "count": 75, "percentage": 50.0 }
  ]
}
```

#### 获取趋势分析
**请求:**
```bash
curl -X GET http://localhost:8084/api/v1/statistics/trends \
-H "Authorization: Bearer <your_token>"
```
**响应 (200 OK):**
```json
{
  "dispatch_trend": [
    { "date": "2024-01-01", "task_count": 5 },
    { "date": "2024-01-02", "task_count": 8 },
    { "date": "2024-01-03", "task_count": 12 }
  ],
  "demand_trend": [
    { "date": "2024-01-01", "request_count": 10 },
    { "date": "2024-01-02", "request_count": 15 },
    { "date": "2024-01-03", "request_count": 20 }
  ]
}
```

---

## 6. gRPC 接口声明 (服务间调用)

| 服务 | 接口定义文件 | 描述 |
| :--- | :--- | :--- |
| **Stock** | `proto/stock.proto` | 提供库存查询、扣减、增加等接口 |
| **Dispatch** | `proto/dispatch.proto` | 提供调度任务状态变更接口 |
| **Logistics** | `proto/logistics.proto` | 提供物流单创建接口 |
