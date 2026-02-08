<script setup lang="ts">
import { ref, computed } from 'vue'
import { useDispatchStore, type DispatchRequest } from '@/stores/dispatch'
import { Select, Van, Checked, Loading } from '@element-plus/icons-vue' // Icons

const dispatchStore = useDispatchStore()
const activeTab = ref('pending')
const dialogVisible = ref(false)
const selectedRequest = ref<DispatchRequest | null>(null)

// Mock Recommendation Logic
const recommendation = ref<{ warehouse: string; eta: string; route: string }>({ warehouse: '', eta: '', route: '' })
const isAnalysisLoading = ref(false)

const handleAudit = (req: DispatchRequest) => {
    selectedRequest.value = req
    dialogVisible.value = true
    isAnalysisLoading.value = true
    
    // Simulate AI Analysis
    setTimeout(() => {
        recommendation.value = {
            warehouse: 'Central Storage Hub (Distance: 12km)',
            eta: '35 mins',
            route: 'Fastest Route via 3rd Ring Road'
        }
        isAnalysisLoading.value = false
    }, 1500)
}

const confirmAllocation = () => {
    if (selectedRequest.value) {
        dispatchStore.approveRequest(selectedRequest.value.id, recommendation.value.warehouse, recommendation.value.eta)
        dialogVisible.value = false
        // Optionally switch tabs or show success message
    }
}

const handleStartDispatch = (req: DispatchRequest) => {
    // Simulate assigning a driver
    dispatchStore.startDispatch(req.id, `Unit-${Math.floor(Math.random() * 100)}`)
}

const getStatusType = (urgency: string) => {
    switch (urgency) {
        case 'Critical': return 'danger'
        case 'High': return 'warning'
        default: return 'info'
    }
}
</script>

<template>
    <div class="h-full flex flex-col space-y-4">
        <h1 class="text-2xl font-bold text-gray-800">Dispatch Operations</h1>

        <el-tabs v-model="activeTab" class="flex-1 bg-white p-4 rounded-xl shadow-sm border border-gray-100">
            <!-- PENDING REQUESTS -->
            <el-tab-pane label="Pending Requests" name="pending">
                <div class="space-y-4">
                    <div v-if="dispatchStore.pendingRequests.length === 0" class="text-center text-gray-400 py-10">
                        No pending requests. All clear.
                    </div>
                    <div v-for="req in dispatchStore.pendingRequests" :key="req.id" 
                         class="border border-gray-200 rounded-lg p-4 flex justify-between items-center hover:bg-gray-50 transition-colors">
                        <div>
                            <div class="flex items-center gap-2 mb-1">
                                <span class="text-lg font-bold text-gray-800">{{ req.region }}</span>
                                <el-tag :type="getStatusType(req.urgency)" size="small" effect="dark">{{ req.urgency }}</el-tag>
                            </div>
                            <p class="text-sm text-gray-500">Request: {{ req.items }}</p>
                        </div>
                        <el-button type="primary" @click="handleAudit(req)">
                            Audit & Allocate
                        </el-button>
                    </div>
                </div>
            </el-tab-pane>

            <!-- APPROVED / READY -->
            <el-tab-pane label="Ready for Transport" name="approved">
                <div class="space-y-4">
                    <div v-for="req in dispatchStore.requests.filter(r => r.status === 'Approved')" :key="req.id"
                         class="border border-green-100 bg-green-50/30 rounded-lg p-4 flex justify-between items-center">
                        <div>
                            <div class="flex items-center gap-2 mb-1">
                                <span class="font-bold">{{ req.region }}</span>
                                <span class="text-xs bg-green-100 text-green-700 px-2 py-0.5 rounded">Allocated</span>
                            </div>
                            <p class="text-xs text-gray-500">Source: {{ req.warehouse }} • ETA: {{ req.estimatedTime }}</p>
                        </div>
                        <el-button type="success" plain @click="handleStartDispatch(req)">
                            <el-icon class="mr-1"><Van /></el-icon> Assign Driver
                        </el-button>
                    </div>
                </div>
            </el-tab-pane>

             <!-- IN TRANSIT -->
            <el-tab-pane label="Active Dispatches" name="active">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                     <div v-for="req in dispatchStore.activeDispatches" :key="req.id"
                         class="border border-blue-100 bg-blue-50/50 rounded-lg p-4 relative overflow-hidden">
                        <div class="absolute right-0 top-0 p-2 opacity-10">
                            <el-icon :size="60" color="#3B82F6"><Loading /></el-icon>
                        </div>
                        <h3 class="font-bold text-blue-900 mb-2">{{ req.region }}</h3>
                        <div class="text-sm space-y-1 text-blue-800">
                            <p>Cargo: {{ req.items }}</p>
                            <p>Driver: <span class="font-mono font-bold">{{ req.driver }}</span></p>
                            <p>ETA: {{ req.estimatedTime }}</p>
                        </div>
                        <div class="mt-4 pt-2 border-t border-blue-200 flex justify-between items-center">
                            <span class="text-xs font-bold text-blue-600 animate-pulse">● Live Tracking</span>
                            <router-link to="/logistics" class="text-xs underline text-blue-600 hover:text-blue-800">View Map</router-link>
                        </div>
                    </div>
                </div>
            </el-tab-pane>
        </el-tabs>

        <!-- Audit Dialog -->
        <el-dialog v-model="dialogVisible" title="System Allocation Analysis" width="600px">
            <div v-if="selectedRequest">
                <div class="mb-6 p-4 bg-gray-50 rounded-lg border border-gray-200">
                    <h4 class="font-bold text-gray-700 mb-2">Request Details</h4>
                    <p><strong>Destination:</strong> {{ selectedRequest.region }}</p>
                    <p><strong>Urgency:</strong> <span :class="`text-${getStatusType(selectedRequest.urgency)}`">{{ selectedRequest.urgency }}</span></p>
                    <p><strong>Items:</strong> {{ selectedRequest.items }}</p>
                </div>

                <div v-if="isAnalysisLoading" class="flex flex-col items-center justify-center py-8">
                     <el-icon class="is-loading text-emerald-500 mb-2" :size="30"><Loading /></el-icon>
                     <p class="text-sm text-gray-500">Calculating optimal routes and checking inventory...</p>
                </div>

                <div v-else class="space-y-4">
                     <div class="bg-emerald-50 border border-emerald-100 p-4 rounded-lg">
                        <h4 class="font-bold text-emerald-800 flex items-center mb-2">
                            <el-icon class="mr-2"><Select /></el-icon> Recommendation: Best Match
                        </h4>
                        <div class="grid grid-cols-2 gap-4 text-sm">
                            <div>
                                <span class="text-gray-500 block">Warehouse</span>
                                <span class="font-medium">{{ recommendation.warehouse }}</span>
                            </div>
                             <div>
                                <span class="text-gray-500 block">Est. Time</span>
                                <span class="font-medium">{{ recommendation.eta }}</span>
                            </div>
                             <div class="col-span-2">
                                <span class="text-gray-500 block">Suggested Route</span>
                                <span class="font-medium">{{ recommendation.route }}</span>
                            </div>
                        </div>
                     </div>
                     
                     <div class="bg-yellow-50 border border-yellow-100 p-4 rounded-lg opacity-60">
                         <h4 class="font-bold text-yellow-800 text-sm mb-1">Alternative Option</h4>
                         <p class="text-xs text-yellow-700">South Depot (Distance: 25km) - High traffic warning</p>
                     </div>
                </div>
            </div>
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="dialogVisible = false">Cancel</el-button>
                    <el-button type="primary" :disabled="isAnalysisLoading" @click="confirmAllocation">
                        Confirm Allocation
                    </el-button>
                </span>
            </template>
        </el-dialog>
    </div>
</template>
