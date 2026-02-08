import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { stockApi, type Material as ApiMaterial } from '@/api/stock'

export interface Material extends ApiMaterial {
    status: 'High' | 'Low' | 'Critical'
}

export const useInventoryStore = defineStore('inventory', () => {
    const materials = ref<Material[]>([])
    const isLoading = ref(false)

    const alertCount = computed(() => {
        return materials.value.filter(m => m.status === 'Critical' || m.status === 'Low').length
    })

    const criticalItems = computed(() => {
        return materials.value.filter(m => m.status === 'Critical')
    })

    async function fetchMaterials() {
        isLoading.value = true
        try {
            const res = await stockApi.getMaterials({ page: 1, page_size: 100 })
            const apiMaterials: any[] = res.data.materials || [] // Adjust based on actual response structure

            // Transform to internal model with status
            materials.value = apiMaterials.map(m => {
                const qty = m.quantity || 0
                let status: 'High' | 'Low' | 'Critical' = 'High'
                if (qty < 10) status = 'Critical'
                else if (qty < 50) status = 'Low'

                return {
                    ...m,
                    status
                }
            })
        } catch (error) {
            console.error('Failed to fetch materials', error)
            // Fallback mock data for demo if backend is down
            materials.value = [
                { id: 1, name: '医用口罩 (N95)', category: '医疗物资', quantity: 5000, unit: '个', description: '防护等级N95', status: 'High' },
                { id: 2, name: '瓶装饮用水', category: '生活物资', quantity: 1200, unit: '箱', description: '500ml*24', status: 'High' },
                { id: 3, name: '救灾帐篷', category: '安置设备', quantity: 45, unit: '顶', description: '防水加厚', status: 'Low' },
                { id: 4, name: '便携发电机', category: '工具设备', quantity: 5, unit: '台', description: '柴油', status: 'Critical' },
                { id: 5, name: '急救箱', category: '医疗物资', quantity: 15, unit: '个', description: '含常用急救药品', status: 'Critical' },
            ]
        } finally {
            isLoading.value = false
        }
    }

    return { materials, alertCount, criticalItems, fetchMaterials, isLoading }
})
