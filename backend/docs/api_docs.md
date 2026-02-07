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
| `/auth/login` | POST | 用户登录，返回 Access Token |
| `/auth/logout` | POST | 用户登出 |
| `/auth/refresh` | POST | 刷新 Token |

---

## 2. 物资库存服务 (Stock Service)
- **端口**: REST `8082` / gRPC `9092`
- **基础路径**: `http://localhost:8082/api/v1`

| 接口 | 方法 | 描述 |
| :--- | :--- | :--- |
| `/stock/materials` | GET | 获取物资列表 (支持 `page`, `pageSize`, `search` 参数) |
| `/stock/materials` | POST | 创建新物资 |
| `/stock/materials/:id` | GET | 获取物资详情 |
| `/stock/inventory` | GET | 获取库存列表 |
| `/stock/inbound` | POST | 物资入库 |
| `/stock/outbound` | POST | 物资出库 |
| `/stock/transfer` | POST | 物资调拨 |
| `/stock/stats` | GET | 库存统计汇总 |

---

## 3. 调度指挥服务 (Dispatch Service)
- **端口**: REST `8083` / gRPC `9093`
- **基础路径**: `http://localhost:8083/api/v1`

| 接口 | 方法 | 描述 |
| :--- | :--- | :--- |
| `/dispatch/requests` | GET | 获取需求申报列表 (支持 `page`, `pageSize`, `status` 参数) |
| `/dispatch/requests` | POST | 创建物资需求申报 |
| `/dispatch/requests/:id` | GET | 获取需求申报详情 |
| `/dispatch/requests/:id/audit` | POST | 审核需求申报 (`action`: approve/reject) |
| `/dispatch/requests/:id/allocation-suggestion` | GET | 获取库存分配建议 (基于库存服务数据) |
| `/dispatch/tasks` | GET | 获取调度任务列表 |
| `/dispatch/tasks` | POST | 创建调度任务 (确认分配库存) |

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

---

## 5. 统计分析服务 (Statistics Service)
- **端口**: REST `8084`
- **基础路径**: `http://localhost:8084/api/v1`

| 接口 | 方法 | 描述 |
| :--- | :--- | :--- |
| `/statistics/summary` | GET | 获取核心指标汇总 (总物资、总任务等) |
| `/statistics/reports` | GET | 物资分类统计报表 |
| `/statistics/trends` | GET | 物资需求/调度趋势分析 |

---

## 6. gRPC 接口声明 (服务间调用)

| 服务 | 接口定义文件 | 描述 |
| :--- | :--- | :--- |
| **Stock** | `proto/stock.proto` | 提供库存查询、扣减、增加等接口 |
| **Dispatch** | `proto/dispatch.proto` | 提供调度任务状态变更接口 |
| **Logistics** | `proto/logistics.proto` | 提供物流单创建接口 |
