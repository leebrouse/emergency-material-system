# K6 负载测试

本目录包含用于测试应急物资管理系统后端服务的 K6 负载测试脚本。

## 前置要求

1. 安装 K6：
   ```bash
   # Ubuntu/Debian
   sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys C5AD17C747E3415A3642D57D77C6C491D6AC1D69
   echo "deb https://dl.k6.io/deb stable main" | sudo tee /etc/apt/sources.list.d/k6.list
   sudo apt-get update
   sudo apt-get install k6

   # macOS
   brew install k6

   # 或使用 Docker
   docker pull grafana/k6
   ```

2. 确保所有后端服务已启动：
   - Auth Service (端口 8081)
   - Stock Service (端口 8082)
   - Dispatch Service (端口 8083)
   - Statistics Service (端口 8084)
   - Logistics Service (端口 8085)

## 运行测试

### 基本运行

```bash
cd backend/test
k6 run k6_load_test.js
```

### 自定义基础URL

如果服务运行在不同的主机或端口，可以通过环境变量设置：

```bash
BASE_URL=http://localhost k6 run k6_load_test.js
```

### 使用 Docker 运行

```bash
docker run --rm -i grafana/k6 run - < k6_load_test.js
```

## 测试内容

测试脚本会执行以下测试场景：

1. **健康检查**
   - 检查所有5个服务的健康状态端点

2. **认证流程**
   - 用户登录获取访问令牌
   - 验证令牌有效性

3. **库存服务测试**
   - 获取材料列表
   - 创建新材料
   - 获取单个材料详情
   - 查询库存信息

4. **调度服务测试**
   - 获取调度请求列表
   - 创建新的调度请求
   - 获取单个请求详情
   - 更新请求状态

5. **物流服务测试**
   - 创建物流跟踪记录
   - 获取跟踪信息
   - 更新跟踪状态

6. **统计服务测试**
   - 获取概览统计
   - 获取材料统计
   - 获取请求统计

## 测试配置

当前测试配置：
- **阶段1**: 30秒内逐步增加到10个虚拟用户
- **阶段2**: 保持10个用户1分钟
- **阶段3**: 30秒内增加到20个用户
- **阶段4**: 保持20个用户1分钟
- **阶段5**: 30秒内逐步减少到0

**性能阈值**：
- 95%的请求应该在500ms内完成
- 错误率应该低于10%

## 自定义测试配置

您可以修改 `k6_load_test.js` 中的 `options` 对象来自定义测试：

```javascript
export const options = {
  stages: [
    { duration: '1m', target: 50 },  // 1分钟内增加到50个用户
    { duration: '3m', target: 50 },  // 保持50个用户3分钟
    { duration: '1m', target: 0 },   // 1分钟内减少到0
  ],
  thresholds: {
    http_req_duration: ['p(95)<1000'],  // 95%的请求应该在1秒内完成
    http_req_failed: ['rate<0.05'],      // 错误率应该低于5%
  },
};
```

## 查看测试结果

测试运行时会实时显示：
- 请求总数和成功率
- 响应时间统计（平均值、最小值、最大值、P95、P99）
- 错误率
- 自定义指标

## 故障排查

如果测试失败，请检查：

1. **所有服务是否已启动**
   ```bash
   curl http://localhost:8081/health
   curl http://localhost:8082/health
   curl http://localhost:8083/health
   curl http://localhost:8084/health
   curl http://localhost:8085/health
   ```

2. **服务端口是否正确**
   - 检查 `global.yaml` 配置文件中的端口设置

3. **网络连接**
   - 确保防火墙没有阻止端口访问

4. **查看详细日志**
   - 检查各个服务的控制台输出
   - K6 会在测试开始时显示服务可用性检查结果

## 持续集成

可以将 K6 测试集成到 CI/CD 流程中：

```yaml
# .github/workflows/k6-test.yml
name: K6 Load Test
on: [push, pull_request]
jobs:
  k6-test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Run K6 tests
        run: |
          docker run --rm -i grafana/k6 run - < backend/test/k6_load_test.js
```

