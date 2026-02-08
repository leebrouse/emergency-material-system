import apiClient from './client'

export interface DemandItem {
    material_id: number
    quantity: number
}

export interface Demand {
    id: number
    location: string
    priority: string
    status: string
    description: string
    items: DemandItem[]
    created_at?: string
}

export const dispatchApi = {
    getDemands: (params?: { page: number, page_size: number, status?: string }) =>
        apiClient.get('/dispatch/demands', { params }),

    updateDemandStatus: (id: number, status: string) =>
        apiClient.put(`/dispatch/demands/${id}/status`, { status }),

    createOrder: (data: { demand_id: number, warehouse_id: string, vehicle_info: string }) =>
        apiClient.post('/dispatch/orders', data),

    getOrders: (params?: { page: number, page_size: number, status?: string }) =>
        apiClient.get('/dispatch/orders', { params })
}
