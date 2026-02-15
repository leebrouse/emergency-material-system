import apiClient from './client'

export interface DemandItem {
    material_id: number
    quantity: number
}

export interface Demand {
    id: number
    location?: string
    priority?: string
    status: string
    description: string
    items?: DemandItem[]
    created_at?: string
    // New backend fields
    material_id?: number
    quantity?: number
    urgency?: string
    target_area?: string
}

export const dispatchApi = {
    getDemands: (params?: { page: number, page_size: number, status?: string }) =>
        apiClient.get('/dispatch/requests', { params }),

    updateDemandStatus: (id: number, action: 'approve' | 'reject', remark?: string) =>
        apiClient.post(`/dispatch/requests/${id}/audit`, { action, remark }),

    createOrder: (data: { request_id: number, allocations: { inventory_id: number, quantity: number }[] }) =>
        apiClient.post('/dispatch/tasks', data),

    createRequest: (data: {
        material_id: number,
        quantity: number,
        urgency_level: 'L1' | 'L2' | 'L3',
        target_area: string,
        description?: string
    }) => apiClient.post('/dispatch/requests', data),

    getSuggestion: (id: number) =>
        apiClient.get(`/dispatch/requests/${id}/allocation-suggestion`),

    getOrders: (params?: { page: number, page_size: number, status?: string }) =>
        apiClient.get('/dispatch/tasks', { params })
}
