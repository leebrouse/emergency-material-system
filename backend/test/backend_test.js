import http from 'k6/http';
import { check, sleep, group } from 'k6';

const BASE_URLS = {
    auth: 'http://localhost:8081/api/v1',
    stock: 'http://localhost:8082/api/v1',
    dispatch: 'http://localhost:8083/api/v1',
    statistics: 'http://localhost:8084/api/v1',
    logistics: 'http://localhost:8085/api/v1',
};

// 配置测试选项
export const options = {
    thresholds: {
        http_req_duration: ['p(95)<500'], // 95% 的请求应在 500ms 内完成
        http_req_failed: ['rate<0.01'],    // 错误率应低于 1%
    },
    stages: [
        { duration: '30s', target: 1 }, // 启动 5 个并发用户
        { duration: '1m', target: 1 },  // 维持
        { duration: '30s', target: 0 }, // 停止
    ],
};

export default function () {
    let token = '';
    let materialId = 0;
    let inventoryId = 0;
    let requestId = 0;
    let taskId = 0;

    const headers = {
        'Content-Type': 'application/json',
    };

    // 1. 认证登录
    group('Step 1: Auth Login', function () {
        const payload = JSON.stringify({
            username: 'admin',
            password: 'admin123',
        });
        const res = http.post(`${BASE_URLS.auth}/auth/login`, payload, { headers });

        check(res, {
            'login success': (r) => r.status === 200,
            'has access token': (r) => r.json().access_token !== undefined,
        });

        if (res.status === 200) {
            token = res.json().access_token;
            headers['Authorization'] = `Bearer ${token}`;
        }
    });

    if (!token) return; // 登录失败跳过后续

    // 2. 创建物资 (Stock Service)
    group('Step 2: Create Material', function () {
        const materialPayload = JSON.stringify({
            name: `Test Material ${Date.now()}`,
            category: 'Medical',
            specs: 'Standard',
            unit: 'Box',
            description: 'Created by k6 integration test',
        });
        const res = http.post(`${BASE_URLS.stock}/stock/materials`, materialPayload, { headers });

        check(res, {
            'create material success': (r) => r.status === 201,
            'has material id': (r) => r.json().id !== undefined,
        });

        if (res.status === 201) {
            materialId = res.json().id;
        }
    });

    if (!materialId) return;

    // 3. 物资入库 (Stock Service)
    group('Step 3: Inbound Stock', function () {
        const inboundPayload = JSON.stringify({
            material_id: materialId,
            quantity: 1000,
            location: 'Warehouse-A-101',
            operator_id: 1,
            remark: 'Initial stock for testing',
        });
        const res = http.post(`${BASE_URLS.stock}/stock/inbound`, inboundPayload, { headers });

        check(res, {
            'inbound success': (r) => r.status === 200,
        });

        // 获取 inventory_id 供后续分配
        const invRes = http.get(`${BASE_URLS.stock}/stock/inventory`, { headers });
        const items = invRes.json().data || [];
        const item = items.find(i => i.material_id === materialId);
        if (item) {
            inventoryId = item.id;
        }
    });

    // 4. 发起需求申报 (Dispatch Service)
    group('Step 4: Create Demand Request', function () {
        const requestPayload = JSON.stringify({
            material_id: materialId,
            quantity: 50,
            urgency_level: 'L1',
            target_area: 'Central Hospital',
            description: 'Urgent medical supply request',
        });
        const res = http.post(`${BASE_URLS.dispatch}/dispatch/requests`, requestPayload, { headers });

        check(res, {
            'create request success': (r) => r.status === 201,
            'has request id': (r) => r.json().id !== undefined,
        });

        if (res.status === 201) {
            requestId = res.json().id;
        }
    });

    if (!requestId) return;

    // 5. 审核需求 (Dispatch Service)
    group('Step 5: Audit Request', function () {
        const auditPayload = JSON.stringify({
            action: 'approve',
            remark: 'Verified and approved',
        });
        const res = http.post(`${BASE_URLS.dispatch}/dispatch/requests/${requestId}/audit`, auditPayload, { headers });

        check(res, {
            'audit success': (r) => r.status === 200,
        });
    });

    // 6. 创建调度任务 (Dispatch Service)
    group('Step 6: Create Dispatch Task', function () {
        // 先获取分配建议
        const suggestRes = http.get(`${BASE_URLS.dispatch}/dispatch/requests/${requestId}/allocation-suggestion`, { headers });
        const suggestions = suggestRes.json();

        if (suggestions && suggestions.length > 0) {
            const taskPayload = JSON.stringify({
                request_id: requestId,
                allocations: suggestions.map(s => ({
                    inventory_id: s.inventory_id,
                    quantity: s.quantity
                }))
            });
            console.log("taskPayload", taskPayload);
            const res = http.post(`${BASE_URLS.dispatch}/dispatch/tasks`, taskPayload, { headers });

            check(res, {
                'create task success': (r) => r.status === 201,
                'has task id': (r) => r.json().task_id !== undefined,
            });

            if (res.status === 201) {
                taskId = res.json().task_id;
            }
        }
    });

    // 7. 查看统计汇总 (Statistics Service)
    group('Step 7: Get Summary Stats', function () {
        const res = http.get(`${BASE_URLS.statistics}/statistics/summary`, { headers });

        check(res, {
            'get summary success': (r) => r.status === 200,
            'has stats data': (r) => r.json() !== null,
        });
    });

    sleep(1);
}
