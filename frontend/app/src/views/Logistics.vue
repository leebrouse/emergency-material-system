<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed, shallowRef } from 'vue'
import { useDispatchStore } from '@/stores/dispatch'
import { logisticsApi } from '@/api/logistics'
import AMapLoader from '@amap/amap-jsapi-loader'

const dispatchStore = useDispatchStore()
const mapContainer = ref<HTMLDivElement | null>(null)
const activeDrivers = computed(() => dispatchStore.activeDispatches)

// AMap instance
const map = shallowRef<any>(null)
const markers = ref<any[]>([])

// Your AMap API Key - Configure in .env file
const AMAP_KEY = import.meta.env.VITE_AMAP_KEY || 'YOUR_AMAP_KEY'

const initAMap = async () => {
    try {
        const AMap = await AMapLoader.load({
            key: AMAP_KEY,
            version: '2.0',
            plugins: ['AMap.Scale', 'AMap.ToolBar', 'AMap.Marker']
        })

        if (mapContainer.value) {
            map.value = new AMap.Map(mapContainer.value, {
                zoom: 12,
                center: [116.397428, 39.90923], // Beijing center
                mapStyle: 'amap://styles/darkblue'
            })

            map.value.addControl(new AMap.Scale())
            map.value.addControl(new AMap.ToolBar())

            // Add HQ marker
            const hqMarker = new AMap.Marker({
                position: [116.397428, 39.90923],
                title: '指挥中心',
                content: `<div style="background: #10B981; width: 16px; height: 16px; border-radius: 4px; transform: rotate(45deg); border: 2px solid white;"></div>`
            })
            map.value.add(hqMarker)

            updateMarkers(AMap)
        }
    } catch (error) {
        console.error('Failed to load AMap', error)
    }
}

const updateMarkers = (AMap: any) => {
    // Clear existing markers
    markers.value.forEach(m => map.value?.remove(m))
    markers.value = []

    activeDrivers.value.forEach(driver => {
        if (driver.coordinates && map.value) {
            const marker = new AMap.Marker({
                position: driver.coordinates,
                title: driver.driver || `运输任务 #${driver.id}`,
                content: `
                    <div style="position: relative;">
                        <div style="background: #3B82F6; width: 12px; height: 12px; border-radius: 50%; border: 2px solid white; animation: pulse 1.5s infinite;"></div>
                        <div style="position: absolute; top: 16px; left: 50%; transform: translateX(-50%); background: rgba(0,0,0,0.8); color: white; padding: 2px 6px; border-radius: 4px; font-size: 10px; white-space: nowrap;">${driver.driver || '运输中'}</div>
                    </div>
                `
            })
            map.value.add(marker)
            markers.value.push(marker)
        }
    })
}

onMounted(() => {
    dispatchStore.fetchDemands()
    initAMap()
})

onUnmounted(() => {
    map.value?.destroy()
})
</script>

<template>
    <div class="h-full w-full flex flex-col relative">
        <!-- Control Panel -->
        <div class="absolute top-4 left-4 z-10 bg-white/90 backdrop-blur-md p-4 rounded-xl shadow-lg border border-gray-100 max-w-sm">
             <h1 class="text-xl font-bold text-gray-800 mb-2 flex items-center">
                 <el-icon class="mr-2 text-blue-600"><Location /></el-icon>
                 物流指挥中心
             </h1>
             <div class="space-y-3 max-h-[60vh] overflow-y-auto pr-2">
                 <div v-if="activeDrivers.length === 0" class="text-sm text-gray-500 italic py-4 text-center">暂无活动运输任务</div>
                 <div v-for="driver in activeDrivers" :key="driver.id" 
                      class="border-l-4 border-blue-500 pl-3 py-2 hover:bg-blue-50 transition-colors rounded-r bg-white/50">
                     <p class="font-bold text-sm text-gray-800">
                         {{ driver.driver || `任务 #${driver.id}` }} 
                         <span class="text-xs font-normal text-gray-500 ml-1">→ {{ driver.region }}</span>
                     </p>
                     <p class="text-xs text-gray-500 flex items-center mt-1">
                         <el-icon class="mr-1"><Timer /></el-icon> 预计 {{ driver.estimatedTime || '计算中...' }}
                     </p>
                     <p class="text-xs text-gray-400 mt-1">
                         物资: {{ driver.items }}
                     </p>
                 </div>
             </div>
        </div>

        <!-- Legend -->
        <div class="absolute bottom-4 left-4 z-10 bg-white/90 backdrop-blur-md p-3 rounded-lg shadow-lg border border-gray-100 text-xs">
            <div class="flex items-center gap-4">
                <div class="flex items-center">
                    <div class="w-3 h-3 bg-emerald-500 rotate-45 mr-2"></div>
                    <span>指挥中心</span>
                </div>
                <div class="flex items-center">
                    <div class="w-3 h-3 bg-blue-500 rounded-full mr-2"></div>
                    <span>运输车辆</span>
                </div>
            </div>
        </div>

        <!-- Map Container -->
        <div ref="mapContainer" class="flex-1 w-full h-full rounded-xl overflow-hidden shadow-inner border border-gray-200">
            <!-- AMap will render here -->
            <div v-if="!map" class="w-full h-full bg-slate-900 flex items-center justify-center text-gray-400">
                <div class="text-center">
                    <el-icon class="text-emerald-500 animate-spin mb-4" :size="40"><Loading /></el-icon>
                    <p class="text-lg font-mono">正在加载地图...</p>
                    <p class="text-sm text-gray-600 mt-2">请确保已配置高德地图API Key</p>
                </div>
            </div>
        </div>
    </div>
</template>

<style>
@keyframes pulse {
    0%, 100% { transform: scale(1); opacity: 1; }
    50% { transform: scale(1.5); opacity: 0.5; }
}
</style>
