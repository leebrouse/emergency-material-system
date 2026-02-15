import apiClient from './client';
export const stockApi = {
    getMaterials: (params) => apiClient.get('/stock/materials', { params }),
    createMaterial: (data) => apiClient.post('/stock/materials', data),
    updateMaterial: (id, data) => apiClient.put(`/stock/materials/${id}`, data),
    deleteMaterial: (id) => apiClient.delete(`/stock/materials/${id}`),
    getInventory: (material_id) => apiClient.get('/stock/inventory', { params: { material_id } }),
    updateInventory: (data) => {
        const path = `/stock/${data.operation}`;
        // Map frontend fields to backend DTO fields if necessary
        const payload = {
            material_id: data.material_id,
            quantity: data.quantity,
            location: data.location || 'Default-Warehouse',
            from_location: data.from_location,
            to_location: data.to_location,
            remark: 'Operation from frontend'
        };
        return apiClient.post(path, payload);
    }
};
