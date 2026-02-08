<script setup lang="ts">
import { onMounted, ref, computed, watchEffect } from 'vue'
import * as echarts from 'echarts'
import { useInventoryStore } from '@/stores/inventory'
import { useDispatchStore } from '@/stores/dispatch'

const inventoryStore = useInventoryStore()
const dispatchStore = useDispatchStore()
const lineChartRef = ref<HTMLElement | null>(null)
const pieChartRef = ref<HTMLElement | null>(null)

// Computed Stats for Cards
const totalStockQuantity = computed(() => {
    return inventoryStore.materials.reduce((acc, curr) => acc + (curr.quantity || 0), 0)
})

const activeDispatchesCount = computed(() => {
    return dispatchStore.activeDispatches.length
})

const pendingRequestsCount = computed(() => {
    return dispatchStore.pendingRequests.length
})

// Function to init charts with responsive resize
let lineChart: echarts.ECharts | null = null
let pieChart: echarts.ECharts | null = null

const initCharts = () => {
    if (lineChartRef.value) {
        lineChart = echarts.init(lineChartRef.value)
        updateLineChart()
        window.addEventListener('resize', () => lineChart?.resize())
    }

    if (pieChartRef.value) {
        pieChart = echarts.init(pieChartRef.value)
        updatePieChart()
        window.addEventListener('resize', () => pieChart?.resize())
    }
}

const updateLineChart = () => {
    if (!lineChart) return
    lineChart.setOption({
        title: { text: '物资消耗趋势（近一周）', left: 'left', textStyle: { fontSize: 14 } },
        tooltip: { trigger: 'axis' },
        grid: { left: '3%', right: '4%', bottom: '3%', containLabel: true },
        xAxis: {
            type: 'category',
            data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
            axisLine: { lineStyle: { color: '#9CA3AF' } }
        },
        yAxis: { type: 'value', splitLine: { lineStyle: { type: 'dashed' } } },
        series: [
            {
                data: [150, 230, 224, 218, 135, 147, 260],
                type: 'line',
                smooth: true,
                showSymbol: false,
                lineStyle: { color: '#10B981', width: 3 },
                areaStyle: {
                        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                        { offset: 0, color: 'rgba(16, 185, 129, 0.4)' },
                        { offset: 1, color: 'rgba(16, 185, 129, 0)' }
                    ])
                }
            }
        ]
    })
}

const updatePieChart = () => {
    if (!pieChart) return

    // Group materials by category
    const categoryMap: { [key: string]: number } = {}
    inventoryStore.materials.forEach(m => {
        categoryMap[m.category] = (categoryMap[m.category] || 0) + (m.quantity || 0)
    })

    const pieData = Object.entries(categoryMap).map(([name, value]) => ({ name, value }))

    pieChart.setOption({
        title: { text: '库存分类占比', left: 'center', textStyle: { fontSize: 14 } },
        tooltip: { trigger: 'item' },
        legend: { bottom: '5%', left: 'center' },
        series: [
            {
                name: '数量',
                type: 'pie',
                radius: ['40%', '70%'],
                avoidLabelOverlap: false,
                itemStyle: {
                    borderRadius: 10,
                    borderColor: '#fff',
                    borderWidth: 2
                },
                label: { show: false, position: 'center' },
                emphasis: {
                    label: { show: true, fontSize: '18', fontWeight: 'bold' }
                },
                data: pieData
            }
        ]
    })
}

onMounted(() => {
    inventoryStore.fetchMaterials()
    dispatchStore.fetchDemands()
    setTimeout(initCharts, 100) // Slight delay to ensure DOM is ready
})

// React to store changes
watchEffect(() => {
    if (inventoryStore.materials.length > 0) {
        updatePieChart()
    }
})
</script>

<template>
    <div class="h-full w-full">
        <h1 class="text-2xl font-bold mb-6 text-gray-800">运营概览</h1>
        
        <!-- Summary Cards -->
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
            <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-100 hover:shadow-md transition-shadow relative overflow-hidden group">
                <h3 class="text-gray-500 text-sm font-medium uppercase tracking-wider">库存总量</h3>
                <p class="text-3xl font-bold text-gray-900 mt-2">{{ totalStockQuantity.toLocaleString() }}</p>
                <div class="w-full bg-gray-100 rounded-full h-1.5 mt-4">
                     <div class="bg-emerald-500 h-1.5 rounded-full" style="width: 70%"></div>
                </div>
            </div>

             <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-100 hover:shadow-md transition-shadow">
                <h3 class="text-gray-500 text-sm font-medium uppercase tracking-wider">待审批需求</h3>
                <p class="text-3xl font-bold text-yellow-600 mt-2">{{ pendingRequestsCount }}</p>
                <span class="text-gray-400 text-xs font-semibold">等待审核</span>
            </div>

             <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-100 hover:shadow-md transition-shadow">
                <h3 class="text-gray-500 text-sm font-medium uppercase tracking-wider">运输中任务</h3>
                <p class="text-3xl font-bold text-blue-600 mt-2">{{ activeDispatchesCount }}</p>
                <span class="text-blue-500 text-xs font-semibold animate-pulse">● 实时追踪</span>
            </div>

             <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-100 hover:shadow-md transition-shadow">
                <h3 class="text-gray-500 text-sm font-medium uppercase tracking-wider">紧急告警</h3>
                <p class="text-3xl font-bold text-red-600 mt-2">{{ inventoryStore.criticalItems.length }}</p>
                <span class="text-red-500 text-xs font-semibold">需立即处理</span>
            </div>
        </div>

        <!-- Charts Grid -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 h-[400px]">
            <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-100 h-full flex flex-col">
                <div ref="lineChartRef" class="flex-1 w-full h-full"></div>
            </div>
            <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-100 h-full flex flex-col">
                <div ref="pieChartRef" class="flex-1 w-full h-full"></div>
            </div>
        </div>
    </div>
</template>
