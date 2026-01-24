import http from 'k6/http';
import { check, group, sleep } from 'k6';

/**
 * Dispatch Service Functional Test (k6)
 * 
 * Verifies the full workflow:
 * 1. Material & Stock Initialization (via Stock Service)
 * 2. Demand Request Creation
 * 3. Multi-stage Audit
 * 4. Intelligent Allocation Suggestion
 * 5. Task Generation & Stock Locking
 */

const STOCK_URL = __ENV.STOCK_URL || 'http://localhost:8082/api/v1';
const DISPATCH_URL = __ENV.DISPATCH_URL || 'http://localhost:8083/api/v1';

export const options = {
    vus: 1,
    iterations: 1,
};

export default function () {
    const headers = { 'Content-Type': 'application/json' };
    let materialId;
    let requestId;
    let suggestions = [];

    // --- SETUP: Create Material and Stock ---
    group('0. Setup Stock Environment', function () {
        console.log("Creating test material...");
        const matRes = http.post(`${STOCK_URL}/stock/materials`, JSON.stringify({
            name: `Critical Med-Kit ${Date.now()}`,
            category: "Medical",
            specs: "Type-A",
            unit: "Kit",
            batch_num: "BATCH-2026-X"
        }), { headers });

        if (!check(matRes, { 'Material setup success': (r) => r.status === 201 })) {
            console.error(`Setup Failed: ${matRes.status} - ${matRes.body}`);
            return;
        }
        materialId = matRes.json().id;

        console.log(`Inbounding stock for Material ${materialId}...`);
        const inRes = http.post(`${STOCK_URL}/stock/inbound`, JSON.stringify({
            material_id: materialId,
            location: "HUB-NORTH-01",
            quantity: 100,
            operator_id: 99,
            remark: "Initial inventory setup"
        }), { headers });
        check(inRes, { 'Stock setup success': (r) => r.status === 200 });
    });

    if (!materialId) return;

    // --- STEP 1: Post Dispatch Request ---
    group('1. Demand Request Creation', function () {
        console.log(`Creating demand request for ${materialId}...`);
        const reqPayload = JSON.stringify({
            material_id: materialId,
            quantity: 40,
            urgency_level: "L1",
            target_area: "Zone-Red-05",
            description: "Immediate need for emergency kits"
        });
        const res = http.post(`${DISPATCH_URL}/dispatch/requests`, reqPayload, { headers });
        
        check(res, { 'Request created': (r) => r.status === 201 });
        requestId = res.json().id;
        console.log(`✅ Request ID: ${requestId}`);
    });

    // --- STEP 2: Audit Request ---
    group('2. Audit Process', function () {
        console.log(`Approving request ${requestId}...`);
        const auditRes = http.post(`${DISPATCH_URL}/dispatch/requests/${requestId}/audit`, JSON.stringify({
            action: "approve",
            remark: "Validated by system admin"
        }), { headers });
        
        check(auditRes, { 'Audit successful': (r) => r.status === 200 });
    });

    // --- STEP 3: Get Allocation Suggestion ---
    group('3. Intelligent Allocation', function () {
        console.log(`Fetching suggestion for request ${requestId}...`);
        const res = http.get(`${DISPATCH_URL}/dispatch/requests/${requestId}/allocation-suggestion`, { headers });
        
        check(res, { 
            'Got suggestions': (r) => r.status === 200,
            'Suggestions not empty': (r) => Array.isArray(r.json()) && r.json().length > 0
        });
        
        suggestions = res.json();
        console.log(`🔍 Suggestion: ${JSON.stringify(suggestions)}`);
    });

    // --- STEP 4: Create Dispatch Task ---
    group('4. Dispatch Execution', function () {
        if (!Array.isArray(suggestions)) {
            console.error(`Skipping Task Creation: Suggestions were not returned as an array. Got: ${JSON.stringify(suggestions)}`);
            return;
        }

        console.log(`Generating task for ${requestId}...`);
        const taskPayload = JSON.stringify({
            request_id: requestId,
            allocations: suggestions.map(s => ({
                inventory_id: s.inventory_id,
                quantity: s.quantity
            }))
        });
        const res = http.post(`${DISPATCH_URL}/dispatch/tasks`, taskPayload, { headers });
        
        check(res, { 'Task generated': (r) => r.status === 201 });
        if (res.status === 201) {
            console.log(`🚀 Task ID: ${res.json().task_id} generated. Stock is now locked.`);
        }
    });

    // --- STEP 5: Verification ---
    group('5. Final Integrity Check', function () {
        if (!requestId) return;

        // Verify Request Status
        const reqRes = http.get(`${DISPATCH_URL}/dispatch/requests/${requestId}`, { headers });
        check(reqRes, { 'Status is Dispatching': (r) => r.json().status === "Dispatching" });

        // Verify Stock Locking
        const invRes = http.get(`${STOCK_URL}/stock/inventory`, { headers });
        const stockData = invRes.json().data || [];
        const hubStock = stockData.find(i => i.material_id === materialId && i.location === "HUB-NORTH-01");
        
        check(hubStock, { 
            'Inventory quantity updated (60 left)': (s) => s.quantity === 60,
            'Locked quantity updated (40 locked)': (s) => s.locked_quantity === 40 
        });
    });
}
