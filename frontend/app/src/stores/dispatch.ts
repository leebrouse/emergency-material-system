import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { dispatchApi, type Demand } from '@/api/dispatch'

export interface DispatchRequest {
    id: number
    region: string
    urgency: 'High' | 'Medium' | 'Critical'
    items: string
    status: 'Pending' | 'Approved' | 'In Transit' | 'Delivered'
    coordinates?: [number, number]
    warehouse?: string
    estimatedTime?: string
    driver?: string
}

export const useDispatchStore = defineStore('dispatch', () => {
    const requests = ref<DispatchRequest[]>([])
    const isLoading = ref(false)

    const activeDispatches = computed(() => requests.value.filter(r => r.status === 'In Transit'))
    const pendingRequests = computed(() => requests.value.filter(r => r.status === 'Pending'))

    async function fetchDemands() {
        isLoading.value = true
        try {
            const res = await dispatchApi.getDemands({ page: 1, page_size: 100 })
            const demands: Demand[] = res.data.demands || []

            requests.value = demands.map(d => {
                let urgency: 'High' | 'Medium' | 'Critical' = 'Medium'
                if (d.priority === 'high') urgency = 'High'
                else if (d.priority === 'critical') urgency = 'Critical'

                let status: DispatchRequest['status'] = 'Pending'
                if (d.status === 'approved') status = 'Approved'
                else if (d.status === 'dispatched') status = 'In Transit'
                else if (d.status === 'completed') status = 'Delivered'

                return {
                    id: d.id,
                    region: d.location,
                    urgency,
                    items: d.items?.map(i => `${i.quantity}件`).join(', ') || d.description,
                    status,
                    coordinates: [116.481028, 39.989643] as [number, number] // Default coords
                }
            })
        } catch (error) {
            console.error('Failed to fetch demands', error)
            // Fallback mock data
            requests.value = [
                { id: 101, region: '东部救援点 (4号帐篷区)', urgency: 'High', items: '200顶帐篷', status: 'Pending', coordinates: [116.481028, 39.989643] },
                { id: 102, region: '北部医疗站', urgency: 'Medium', items: '500箱饮用水', status: 'Approved', coordinates: [116.410003, 39.905694], warehouse: '中央物资仓库', estimatedTime: '45分钟' },
                { id: 103, region: '市人民医院A区', urgency: 'Critical', items: '血浆及急救药品', status: 'In Transit', coordinates: [116.322056, 39.897445], warehouse: '医疗物资储备库', estimatedTime: '12分钟', driver: '运输车-04' },
            ]
        } finally {
            isLoading.value = false
        }
    }

    async function approveRequest(id: number, warehouse: string, eta: string) {
        try {
            await dispatchApi.updateDemandStatus(id, 'approved')
            const req = requests.value.find(r => r.id === id)
            if (req) {
                req.status = 'Approved'
                req.warehouse = warehouse
                req.estimatedTime = eta
            }
        } catch (error) {
            console.error('Failed to approve', error)
            // Still update locally for demo
            const req = requests.value.find(r => r.id === id)
            if (req) {
                req.status = 'Approved'
                req.warehouse = warehouse
                req.estimatedTime = eta
            }
        }
    }

    function startDispatch(id: number, driver: string) {
        const req = requests.value.find(r => r.id === id)
        if (req) {
            req.status = 'In Transit'
            req.driver = driver
        }
    }

    return { requests, activeDispatches, pendingRequests, fetchDemands, approveRequest, startDispatch, isLoading }
})
