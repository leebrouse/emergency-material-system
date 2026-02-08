import apiClient from './client'

export interface Material {
    id: number
    name: string
    category: string
    unit: string
    description: string
    quantity?: number // Injected from inventory endpoint usually? Ah docs say inventory is separate endpoint.
}

export interface Inventory {
    material_id: number
    quantity: number
}

export const stockApi = {
    getMaterials: (params?: { page: number, page_size: number, keyword?: string }) =>
        apiClient.get('/stock/materials', { params }),

    getInventory: (material_id?: number) =>
        apiClient.get('/stock/inventory', { params: { material_id } }),

    updateInventory: (data: { material_id: number, quantity: number, operation: string }) =>
        apiClient.post('/stock/inventory', data)
}
