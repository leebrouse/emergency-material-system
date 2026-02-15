import apiClient from './client';
export const statisticsApi = {
    getSummary: () => apiClient.get('/statistics/summary'),
    getReports: () => apiClient.get('/statistics/reports'),
    getTrends: () => apiClient.get('/statistics/trends'),
};
