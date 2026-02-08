# 应急物资管理系统 API 文档

本文档整理了应急物资管理系统各微服务的 API 接口，已与后端 `.proto` 定义同步。

## 通用信息

- **基础路径**: `/api/v1`
- **认证方式**: Bearer Token (JWT)

---

## 1. 身份认证服务 (Auth Service)
- **基础路径**: `http://localhost:8081/api/v1`

| 接口 | 方法 | 描述 | 请求参数 | 响应字段 |
| :--- | :--- | :--- | :--- | :--- |
| `/auth/login` | POST | 用户登录 | `username`, `password` | `token`, `refresh_token`, `expires_in` |
| `/auth/logout` | POST | 用户登出 | `token` | `success` |
| `/auth/refresh` | POST | 刷新 Token | `refresh_token` | `token`, `refresh_token`, `expires_in` |

---

## 2. 物资库存服务 (Stock Service)
- **基础路径**: `http://localhost:8082/api/v1`

| 接口 | 方法 | 描述 | 请求参数 | 响应字段 |
| :--- | :--- | :--- | :--- | :--- |
| `/stock/materials` | GET | 获取物资列表 | `page`, `page_size`, `keyword` | `materials[]`, `total` |
| `/stock/materials` | POST | 创建新物资 | `name`, `category`, `unit`, `description` | `material` |
| `/stock/materials/:id` | GET | 获取物资详情 | - | `material` |
| `/stock/materials/:id` | PUT | 更新物资 | `name`, `category`, `unit`, `description` | `material` |
| `/stock/inventory` | GET | 获取库存列表 | `material_id` | `inventory` |
| `/stock/inventory` | POST | 更新库存量 | `material_id`, `quantity`, `operation` | `inventory` |

---

## 3. 调度指挥服务 (Dispatch Service)
- **基础路径**: `http://localhost:8083/api/v1`

| 接口 | 方法 | 描述 | 请求参数 | 响应字段 |
| :--- | :--- | :--- | :--- | :--- |
| `/dispatch/demands` | GET | 获取需求列表 | `page`, `page_size`, `status` | `demands[]`, `total` |
| `/dispatch/demands` | POST | 创建需求申报 | `location`, `priority`, `description`, `items[]` | `demand` |
| `/dispatch/demands/:id` | GET | 获取需求详情 | - | `demand` |
| `/dispatch/demands/:id/status` | PUT | 更新需求状态 | `status` | `demand` |
| `/dispatch/orders` | POST | 创建调度订单 | `demand_id`, `warehouse_id`, `vehicle_info` | `order` |
| `/dispatch/orders` | GET | 获取订单列表 | `page`, `page_size`, `status` | `orders[]`, `total` |

---

## 4. 物流追踪服务 (Logistics Service)
- **基础路径**: `http://localhost:8085/api/v1`

| 接口 | 方法 | 描述 | 请求参数 | 响应字段 |
| :--- | :--- | :--- | :--- | :--- |
| `/logistics/tracking` | POST | 创建物流追踪 | `order_id`, `vehicle_id`, `driver_info` | `tracking` |
| `/logistics/tracking/:order_id` | GET | 获取追踪信息 | - | `tracking` |
| `/logistics/tracking/:id/location` | POST | 更新实时位置 | `latitude`, `longitude`, `address` | `tracking` |
| `/logistics/tracking/:order_id/history` | GET | 获取历史轨迹 | - | `points[]` |

---

## 5. 统计分析服务 (Statistics Service)
- **基础路径**: `http://localhost:8084/api/v1`

| 接口 | 方法 | 描述 | 请求参数 | 响应字段 |
| :--- | :--- | :--- | :--- | :--- |
| `/statistics/summary` | GET | 核心指标汇总 | - | - |
| `/statistics/reports` | GET | 分类统计报表 | - | - |
| `/statistics/trends` | GET | 趋势分析 | - | - |

---

## 6. 数据模型说明

### Material (物资)
- `id`: int64
- `name`: string
- `category`: string
- `unit`: string
- `description`: string

### Demand (需求)
- `id`: int64
- `location`: string
- `priority`: string (e.g. low, medium, high)
- `status`: string (e.g. pending, approved, rejected, dispatched)
- `items`: DemandItem[]

### Tracking (物流)
- `id`: int64
- `status`: string (e.g. pending, delivering, completed)
- `current_location`: TrackingPoint