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
            const demands: Demand[] = res.data.data || []

            requests.value = demands.map(d => {
                let urgency: 'High' | 'Medium' | 'Critical' = 'Medium'
                // Map backend urgency (L1/L2/L3) or legacy priority to frontend urgency
                const p = (d.priority || d.urgency || 'L2').toUpperCase()
                if (p === 'HIGH' || p === 'L2') urgency = 'High'
                else if (p === 'CRITICAL' || p === 'L3') urgency = 'Critical'
                else if (p === 'L1') urgency = 'Medium'

                let status: DispatchRequest['status'] = 'Pending'
                // Map status
                const s = (d.status || '').toLowerCase()
                if (s === 'approved') status = 'Approved'
                else if (s === 'dispatched' || s === 'intransit') status = 'In Transit'
                else if (s === 'completed' || s === 'delivered') status = 'Delivered'

                // Map items
                let itemsStr = d.description || ''
                if (d.items && d.items.length > 0) {
                    itemsStr = d.items.map(i => `${i.quantity}件`).join(', ')
                } else if (d.quantity && d.material_id) {
                    // If structure is flat (from new createRequest)
                    itemsStr = `物资ID:${d.material_id} x ${d.quantity}`
                }

                return {
                    id: d.id,
                    region: d.location || d.target_area || 'Unknown',
                    urgency,
                    items: itemsStr,
                    status,
                    coordinates: [116.481028, 39.989643] as [number, number] // Default coords
                }
            })
        } catch (error) {
            console.error('Failed to fetch demands', error)
            requests.value = []
        } finally {
            isLoading.value = false
        }
    }

    async function approveRequest(id: number, warehouse: string, eta: string) {
        try {
            await dispatchApi.updateDemandStatus(id, 'approve', `Selected warehouse: ${warehouse}, ETA: ${eta}`)
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

    async function createRequest(data: { material_id: number, quantity: number, urgency: 'L1' | 'L2' | 'L3', region: string }) {
        isLoading.value = true
        try {
            await dispatchApi.createRequest({
                material_id: data.material_id,
                quantity: data.quantity,
                urgency_level: data.urgency,
                target_area: data.region,
                description: 'Generated from frontend'
            })
            await fetchDemands()
        } catch (error) {
            console.error('Failed to create request', error)
        } finally {
            isLoading.value = false
        }
    }

    async function getSuggestion(id: number) {
        try {
            const res = await dispatchApi.getSuggestion(id)
            return res.data
        } catch (error) {
            console.error('Failed to get suggestion', error)
            return []
        }
    }

    return { requests, activeDispatches, pendingRequests, fetchDemands, approveRequest, startDispatch, createRequest, getSuggestion, isLoading }
})
