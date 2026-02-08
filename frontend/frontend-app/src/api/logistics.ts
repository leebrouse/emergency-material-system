import apiClient from './client'

export interface TrackingPoint {
    latitude: number
    longitude: number
    address?: string
    timestamp?: string
}

export interface Tracking {
    id: number
    order_id: number
    vehicle_id: string
    status: string
    current_location: TrackingPoint
}

export const logisticsApi = {
    getTracking: (order_id: number) =>
        apiClient.get(`/logistics/tracking/${order_id}`),

    updateLocation: (id: number, data: { latitude: number, longitude: number, address?: string }) =>
        apiClient.post(`/logistics/tracking/${id}/location`, data)
}
