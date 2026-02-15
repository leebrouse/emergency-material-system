import apiClient from './client'

export interface Summary {
    total_materials: number
    total_requests: number
    pending_requests: number
    completed_requests: number
}

export interface MaterialStat {
    category: string
    count: number
}

export interface TrendPoint {
    date: string
    value: number
}

export const statisticsApi = {
    getSummary: () => apiClient.get<Summary>('/statistics/summary'),
    getReports: () => apiClient.get<MaterialStat[]>('/statistics/reports'),
    getTrends: () => apiClient.get<TrendPoint[]>('/statistics/trends'),
}
