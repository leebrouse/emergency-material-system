<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useDispatchStore, type DispatchRequest } from '@/stores/dispatch'
import { Van, Loading, Select } from '@element-plus/icons-vue'

const dispatchStore = useDispatchStore()
const activeTab = ref('pending')
const dialogVisible = ref(false)
const selectedRequest = ref<DispatchRequest | null>(null)

// Mock Recommendation Logic
const recommendation = ref<{ warehouse: string; eta: string; route: string }>({ warehouse: '', eta: '', route: '' })
const isAnalysisLoading = ref(false)

onMounted(() => {
    dispatchStore.fetchDemands()
})

const handleAudit = (req: DispatchRequest) => {
    selectedRequest.value = req
    dialogVisible.value = true
    isAnalysisLoading.value = true
    
    // Simulate AI Analysis
    setTimeout(() => {
        recommendation.value = {
            warehouse: '中央物资仓库 (距离: 12km)',
            eta: '35分钟',
            route: '最优路线: 三环高速'
        }
        isAnalysisLoading.value = false
    }, 1500)
}

const confirmAllocation = async () => {
    if (selectedRequest.value) {
        await dispatchStore.approveRequest(selectedRequest.value.id, recommendation.value.warehouse, recommendation.value.eta)
        dialogVisible.value = false
    }
}

const handleStartDispatch = (req: DispatchRequest) => {
    dispatchStore.startDispatch(req.id, `运输车-${String(Math.floor(Math.random() * 100)).padStart(2, '0')}`)
}

const getStatusType = (urgency: string) => {
    switch (urgency) {
        case 'Critical': return 'danger'
        case 'High': return 'warning'
        default: return 'info'
    }
}

const getUrgencyLabel = (urgency: string) => {
    switch (urgency) {
        case 'Critical': return '紧急'
        case 'High': return '高优先'
        default: return '普通'
    }
}
</script>

<template>
    <div class="h-full flex flex-col space-y-4">
        <h1 class="text-2xl font-bold text-gray-800">调度指挥</h1>

        <el-tabs v-model="activeTab" class="flex-1 bg-white p-4 rounded-xl shadow-sm border border-gray-100">
            <!-- PENDING REQUESTS -->
            <el-tab-pane label="待审批" name="pending">
                <div class="space-y-4">
                    <div v-if="dispatchStore.pendingRequests.length === 0" class="text-center text-gray-400 py-10">
                        暂无待审批需求，一切正常
                    </div>
                    <div v-for="req in dispatchStore.pendingRequests" :key="req.id" 
                         class="border border-gray-200 rounded-lg p-4 flex justify-between items-center hover:bg-gray-50 transition-colors">
                        <div>
                            <div class="flex items-center gap-2 mb-1">
                                <span class="text-lg font-bold text-gray-800">{{ req.region }}</span>
                                <el-tag :type="getStatusType(req.urgency)" size="small" effect="dark">{{ getUrgencyLabel(req.urgency) }}</el-tag>
                            </div>
                            <p class="text-sm text-gray-500">需求物资: {{ req.items }}</p>
                        </div>
                        <el-button type="primary" @click="handleAudit(req)">
                            审核分配
                        </el-button>
                    </div>
                </div>
            </el-tab-pane>

            <!-- APPROVED / READY -->
            <el-tab-pane label="待发货" name="approved">
                <div class="space-y-4">
                    <div v-for="req in dispatchStore.requests.filter(r => r.status === 'Approved')" :key="req.id"
                         class="border border-green-100 bg-green-50/30 rounded-lg p-4 flex justify-between items-center">
                        <div>
                            <div class="flex items-center gap-2 mb-1">
                                <span class="font-bold">{{ req.region }}</span>
                                <span class="text-xs bg-green-100 text-green-700 px-2 py-0.5 rounded">已分配</span>
                            </div>
                            <p class="text-xs text-gray-500">仓库: {{ req.warehouse }} • 预计 {{ req.estimatedTime }}</p>
                        </div>
                        <el-button type="success" plain @click="handleStartDispatch(req)">
                            <el-icon class="mr-1"><Van /></el-icon> 安排发运
                        </el-button>
                    </div>
                    <div v-if="dispatchStore.requests.filter(r => r.status === 'Approved').length === 0" class="text-center text-gray-400 py-10">
                        暂无待发货任务
                    </div>
                </div>
            </el-tab-pane>

             <!-- IN TRANSIT -->
            <el-tab-pane label="运输中" name="active">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                     <div v-for="req in dispatchStore.activeDispatches" :key="req.id"
                         class="border border-blue-100 bg-blue-50/50 rounded-lg p-4 relative overflow-hidden">
                        <div class="absolute right-0 top-0 p-2 opacity-10">
                            <el-icon :size="60" color="#3B82F6"><Loading /></el-icon>
                        </div>
                        <h3 class="font-bold text-blue-900 mb-2">{{ req.region }}</h3>
                        <div class="text-sm space-y-1 text-blue-800">
                            <p>物资: {{ req.items }}</p>
                            <p>运输车辆: <span class="font-mono font-bold">{{ req.driver }}</span></p>
                            <p>预计到达: {{ req.estimatedTime }}</p>
                        </div>
                        <div class="mt-4 pt-2 border-t border-blue-200 flex justify-between items-center">
                            <span class="text-xs font-bold text-blue-600 animate-pulse">● 实时追踪中</span>
                            <router-link to="/logistics" class="text-xs underline text-blue-600 hover:text-blue-800">查看地图</router-link>
                        </div>
                    </div>
                    <div v-if="dispatchStore.activeDispatches.length === 0" class="col-span-2 text-center text-gray-400 py-10">
                        当前无运输中任务
                    </div>
                </div>
            </el-tab-pane>
        </el-tabs>

        <!-- Audit Dialog -->
        <el-dialog v-model="dialogVisible" title="智能分配审核" width="600px">
            <div v-if="selectedRequest">
                <div class="mb-6 p-4 bg-gray-50 rounded-lg border border-gray-200">
                    <h4 class="font-bold text-gray-700 mb-2">需求详情</h4>
                    <p><strong>目的地:</strong> {{ selectedRequest.region }}</p>
                    <p><strong>紧急程度:</strong> <span :class="`text-${getStatusType(selectedRequest.urgency)}`">{{ getUrgencyLabel(selectedRequest.urgency) }}</span></p>
                    <p><strong>物资:</strong> {{ selectedRequest.items }}</p>
                </div>

                <div v-if="isAnalysisLoading" class="flex flex-col items-center justify-center py-8">
                     <el-icon class="is-loading text-emerald-500 mb-2" :size="30"><Loading /></el-icon>
                     <p class="text-sm text-gray-500">正在计算最优路线并检查库存...</p>
                </div>

                <div v-else class="space-y-4">
                     <div class="bg-emerald-50 border border-emerald-100 p-4 rounded-lg">
                        <h4 class="font-bold text-emerald-800 flex items-center mb-2">
                            <el-icon class="mr-2"><Select /></el-icon> 推荐方案: 最优匹配
                        </h4>
                        <div class="grid grid-cols-2 gap-4 text-sm">
                            <div>
                                <span class="text-gray-500 block">发货仓库</span>
                                <span class="font-medium">{{ recommendation.warehouse }}</span>
                            </div>
                             <div>
                                <span class="text-gray-500 block">预计时间</span>
                                <span class="font-medium">{{ recommendation.eta }}</span>
                            </div>
                             <div class="col-span-2">
                                <span class="text-gray-500 block">推荐路线</span>
                                <span class="font-medium">{{ recommendation.route }}</span>
                            </div>
                        </div>
                     </div>
                     
                     <div class="bg-yellow-50 border border-yellow-100 p-4 rounded-lg opacity-60">
                         <h4 class="font-bold text-yellow-800 text-sm mb-1">备选方案</h4>
                         <p class="text-xs text-yellow-700">南部仓库 (距离: 25km) - 交通拥堵预警</p>
                     </div>
                </div>
            </div>
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="dialogVisible = false">取消</el-button>
                    <el-button type="primary" :disabled="isAnalysisLoading" @click="confirmAllocation">
                        确认分配
                    </el-button>
                </span>
            </template>
        </el-dialog>
    </div>
</template>
