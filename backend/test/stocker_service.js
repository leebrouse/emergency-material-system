import http from 'k6/http';
import { check, group, sleep } from 'k6';

/**
 * Stock Service Functional Test (Improved k6 Version)
 * 
 * This test strictly follows the ServerInterface defined in 
 * backend/internal/common/genopenapi/stock/server.gen.go
 * 
 * Paths tested:
 * - POST /api/v1/stock/materials
 * - GET  /api/v1/stock/materials
 * - GET  /api/v1/stock/materials/{id}
 * - POST /api/v1/stock/inbound
 * - GET  /api/v1/stock/inventory
 * - POST /api/v1/stock/outbound
 * - POST /api/v1/stock/transfer
 * - GET  /api/v1/stock/stats
 */

const BASE_URL = __ENV.BASE_URL || 'http://localhost:8082/api/v1';

export const options = {
    thresholds: {
        http_req_failed: ['rate<0.01'], // <1% errors
    },
};

export default function () {
    const headers = { 'Content-Type': 'application/json' };
    let materialId;

    group('Material Management', function () {
        // 1. Create Material
        console.log("Step 1: POST /stock/materials");
        const materialPayload = JSON.stringify({
            name: "Medical Oxygen Tank",
            category: "Gas",
            specs: "10L",
            unit: "Cylinder",
            batch_num: "OX-2024-001",
            description: "High pressure oxygen for emergency"
        });

        const createRes = http.post(`${BASE_URL}/stock/materials`, materialPayload, { headers });

        const isCreated = check(createRes, {
            'status is 201': (r) => r.status === 201,
            'has id': (r) => r.json().id !== undefined,
        });

        if (!isCreated) {
            console.error(`Status: ${createRes.status}, Body: ${createRes.body}`);
            return;
        }
        materialId = createRes.json().id;
        console.log(`✅ Material created with ID: ${materialId}`);

        // 2. List Materials
        console.log("Step 2: GET /stock/materials");
        const listRes = http.get(`${BASE_URL}/stock/materials?page=1&page_size=10`, { headers });
        check(listRes, {
            'list returns 200': (r) => r.status === 200,
            'data is array': (r) => Array.isArray(r.json().data),
        });

        // 3. Get Specific Material
        console.log(`Step 3: GET /stock/materials/${materialId}`);
        const getRes = http.get(`${BASE_URL}/stock/materials/${materialId}`, { headers });
        check(getRes, {
            'detail returns 200': (r) => r.status === 200,
            'name matches': (r) => r.json().name === "Medical Oxygen Tank",
        });
    });

    if (!materialId) return;

    group('Inventory Operations', function () {
        // 4. Inbound
        console.log("Step 4: POST /stock/inbound");
        const inboundPayload = JSON.stringify({
            material_id: materialId,
            location: "Room-101",
            quantity: 100,
            operator_id: 1,
            remark: "Procured from Supplier A"
        });
        const inboundRes = http.post(`${BASE_URL}/stock/inbound`, inboundPayload, { headers });
        check(inboundRes, { 'inbound status is 200': (r) => r.status === 200 });

        // 5. Check Inventory
        console.log("Step 5: GET /stock/inventory");
        const invRes = http.get(`${BASE_URL}/stock/inventory`, { headers });
        const invData = invRes.json().data || [];
        const entry = invData.find(i => i.material_id === materialId && i.location === "Room-101");
        check(entry, { 'inventory entry created': (e) => e !== undefined && e.quantity === 100 });

        // 6. Outbound
        console.log("Step 6: POST /stock/outbound");
        const outboundPayload = JSON.stringify({
            material_id: materialId,
            location: "Room-101",
            quantity: 30,
            operator_id: 2,
            remark: "Deployment to ER"
        });
        const outboundRes = http.post(`${BASE_URL}/stock/outbound`, outboundPayload, { headers });
        check(outboundRes, { 'outbound status is 200': (r) => r.status === 200 });

        // 7. Transfer
        console.log("Step 7: POST /stock/transfer");
        const transferPayload = JSON.stringify({
            material_id: materialId,
            from_location: "Room-101",
            to_location: "Ambulance-01",
            quantity: 20,
            operator_id: 1,
            remark: "Ready for transport"
        });
        const transferRes = http.post(`${BASE_URL}/stock/transfer`, transferPayload, { headers });
        check(transferRes, { 'transfer status is 200': (r) => r.status === 200 });
    });

    group('Statistics', function () {
        // 8. Final Stats
        console.log("Step 8: GET /stock/stats");
        const statsRes = http.get(`${BASE_URL}/stock/stats`, { headers });
        check(statsRes, {
            'stats returns 200': (r) => r.status === 200,
            'is array': (r) => Array.isArray(r.json()),
        });

        // Verify final quantities locally
        const invRes = http.get(`${BASE_URL}/stock/inventory`, { headers });
        const data = invRes.json().data || [];

        const loc101 = data.find(i => i.material_id === materialId && i.location === "Room-101");
        const amb01 = data.find(i => i.material_id === materialId && i.location === "Ambulance-01");

        console.log(`📊 Final Balance - Room-101: ${loc101 ? loc101.quantity : 0}, Ambulance-01: ${amb01 ? amb01.quantity : 0}`);

        check(null, {
            'Room-101 expected 50': () => loc101 && loc101.quantity === 50,
            'Ambulance-01 expected 20': () => amb01 && amb01.quantity === 20,
        });
    });
}
