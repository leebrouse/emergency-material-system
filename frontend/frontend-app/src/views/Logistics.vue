<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useDispatchStore } from '@/stores/dispatch'
import { Van, Location, Timer } from '@element-plus/icons-vue'

const dispatchStore = useDispatchStore()
const mapContainer = ref<HTMLDivElement | null>(null)
const activeDrivers = computed(() => dispatchStore.activeDispatches)

const mockCoordinates = ref<{ [key: number]: { lat: number, lng: number } }>({})

// Simulate movement
const updatePositions = () => {
    activeDrivers.value.forEach(d => {
        if (!mockCoordinates.value[d.id]) {
            mockCoordinates.value[d.id] = { lat: 39.90, lng: 116.40 } // Start near center
        }
        // Random small movement
        mockCoordinates.value[d.id].lat += (Math.random() - 0.5) * 0.001
        mockCoordinates.value[d.id].lng += (Math.random() - 0.5) * 0.001
    })
}

let interval: number

onMounted(() => {
    interval = setInterval(updatePositions, 2000)
})

</script>

<template>
    <div class="h-full w-full flex flex-col relative">
        <div class="absolute top-4 left-4 z-10 bg-white/90 backdrop-blur-md p-4 rounded-xl shadow-lg border border-gray-100 max-w-sm">
             <h1 class="text-xl font-bold text-gray-800 mb-2 flex items-center">
                 <el-icon class="mr-2 text-blue-600"><Location /></el-icon>
                 Logistics Command
             </h1>
             <div class="space-y-3 max-h-[60vh] overflow-y-auto pr-2">
                 <div v-if="activeDrivers.length === 0" class="text-sm text-gray-500 italic">No active dispatches.</div>
                 <div v-for="driver in activeDrivers" :key="driver.id" 
                      class="border-l-4 border-blue-500 pl-3 py-1 hover:bg-blue-50 transition-colors rounded-r">
                     <p class="font-bold text-sm text-gray-800">{{ driver.driver }} <span class="text-xs font-normal text-gray-500">to {{ driver.region }}</span></p>
                     <p class="text-xs text-gray-500 flex items-center mt-1">
                         <el-icon class="mr-1"><Timer /></el-icon> {{ driver.estimatedTime }} • 
                         <span class="ml-1 font-mono text-blue-600">
                             {{ mockCoordinates[driver.id]?.lat.toFixed(4) }}, {{ mockCoordinates[driver.id]?.lng.toFixed(4) }}
                         </span>
                     </p>
                 </div>
             </div>
        </div>

        <div class="flex-1 bg-slate-900 rounded-xl overflow-hidden shadow-inner relative border border-gray-700">
             <!-- Fake Map Background Grid -->
             <div class="absolute inset-0 opacity-20" 
                  style="background-image: radial-gradient(#4b5563 1px, transparent 1px); background-size: 20px 20px;">
             </div>
             
             <!-- Fake Map Markers -->
             <div v-for="driver in activeDrivers" :key="driver.id"
                  class="absolute transition-all duration-[2000ms] ease-linear flex flex-col items-center"
                  :style="{
                      top: `${50 + (mockCoordinates[driver.id]?.lat - 39.90) * 1000}%`,
                      left: `${50 + (mockCoordinates[driver.id]?.lng - 116.40) * 1000}%`
                  }">
                  <div class="w-3 h-3 bg-blue-500 rounded-full animate-ping absolute"></div>
                  <div class="w-3 h-3 bg-blue-400 rounded-full border-2 border-white relative z-10 shadow-lg"></div>
                  <span class="mt-1 text-[10px] bg-black/70 text-white px-1 rounded backdrop-blur-sm whitespace-nowrap">{{ driver.driver }}</span>
             </div>

             <!-- Center Hub Marker -->
             <div class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 flex flex-col items-center">
                 <div class="w-4 h-4 bg-emerald-500 rounded-sm rotate-45 border-2 border-white shadow-xl"></div>
                 <span class="mt-2 text-xs font-bold text-emerald-400">HQ</span>
             </div>
        </div>
    </div>
</template>
