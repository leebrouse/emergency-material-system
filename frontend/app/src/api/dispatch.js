import apiClient from './client';
export const dispatchApi = {
    getDemands: (params) => apiClient.get('/dispatch/requests', { params }),
    updateDemandStatus: (id, action, remark) => apiClient.post(`/dispatch/requests/${id}/audit`, { action, remark }),
    createOrder: (data) => apiClient.post('/dispatch/tasks', data),
    createRequest: (data) => apiClient.post('/dispatch/requests', data),
    getSuggestion: (id) => apiClient.get(`/dispatch/requests/${id}/allocation-suggestion`),
    getOrders: (params) => apiClient.get('/dispatch/tasks', { params })
};
