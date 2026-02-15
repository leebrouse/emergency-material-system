import apiClient from './client';
export const logisticsApi = {
    getTracking: (order_id) => apiClient.get(`/logistics/tracking/${order_id}`),
    updateLocation: (id, data) => apiClient.post(`/logistics/tracking/${id}/location`, data)
};
