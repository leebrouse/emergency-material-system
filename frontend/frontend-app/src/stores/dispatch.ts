import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export interface DispatchRequest {
    id: number
    region: string
    urgency: 'High' | 'Medium' | 'Critical'
    items: string // Simplified for demo
    status: 'Pending' | 'Approved' | 'In Transit' | 'Delivered'
    coordinates?: [number, number] // Destination
    warehouse?: string // Assigned source
    estimatedTime?: string
    driver?: string
}

export const useDispatchStore = defineStore('dispatch', () => {
    const requests = ref<DispatchRequest[]>([
        { id: 101, region: 'East Zone (Shelter 4)', urgency: 'High', items: '200 Tents', status: 'Pending', coordinates: [116.481028, 39.989643] }, // Example Coords
        { id: 102, region: 'North Zone (Hospital)', urgency: 'Medium', items: '500 Water Packs', status: 'Approved', coordinates: [116.410003, 39.905694], warehouse: 'Central Hub', estimatedTime: '45 mins' },
        { id: 103, region: 'Hospital A', urgency: 'Critical', items: 'Blood Plasma', status: 'In Transit', coordinates: [116.322056, 39.897445], warehouse: 'Medical Depot', estimatedTime: '12 mins', driver: 'Unit-04' },
    ])

    const activeDispatches = computed(() => requests.value.filter(r => r.status === 'In Transit'))
    const pendingRequests = computed(() => requests.value.filter(r => r.status === 'Pending'))

    function approveRequest(id: number, warehouse: string, eta: string) {
        const req = requests.value.find(r => r.id === id)
        if (req) {
            req.status = 'Approved'
            req.warehouse = warehouse
            req.estimatedTime = eta
        }
    }

    function startDispatch(id: number, driver: string) {
        const req = requests.value.find(r => r.id === id)
        if (req) {
            req.status = 'In Transit'
            req.driver = driver
        }
    }

    return { requests, activeDispatches, pendingRequests, approveRequest, startDispatch }
})
