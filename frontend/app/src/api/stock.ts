import apiClient from './client'

export interface Material {
    id: number
    name: string
    category: string
    specs?: string
    unit: string
    batch_num?: string
    description: string
    quantity?: number
    min_stock?: number
}

export interface Inventory {
    material_id: number
    quantity: number
}

export const stockApi = {
    getMaterials: (params?: { page: number, page_size: number, search?: string }) =>
        apiClient.get('/stock/materials', { params }),

    createMaterial: (data: Partial<Material>) =>
        apiClient.post('/stock/materials', data),

    updateMaterial: (id: number, data: Partial<Material>) =>
        apiClient.put(`/stock/materials/${id}`, data),

    deleteMaterial: (id: number) =>
        apiClient.delete(`/stock/materials/${id}`),

    getInventory: (material_id?: number) =>
        apiClient.get('/stock/inventory', { params: { material_id } }),

    updateInventory: (data: {
        material_id: number,
        quantity: number,
        operation: 'inbound' | 'outbound' | 'transfer',
        location?: string,
        from_location?: string,
        to_location?: string,
        target_warehouse_id?: string
    }) => {
        const path = `/stock/${data.operation}`
        // Map frontend fields to backend DTO fields if necessary
        const payload = {
            material_id: data.material_id,
            quantity: data.quantity,
            location: data.location || 'Default-Warehouse',
            from_location: data.from_location,
            to_location: data.to_location,
            remark: 'Operation from frontend'
        }
        return apiClient.post(path, payload)
    }
}
