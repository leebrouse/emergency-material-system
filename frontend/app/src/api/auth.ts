import apiClient from './client'

export interface LoginRequest {
    username: string
    password: string
}

export interface LoginResponse {
    token: string
    refresh_token: string
    expires_in: number
}

export const authApi = {
    login: (data: LoginRequest) => apiClient.post<LoginResponse>('/auth/login', data),
    logout: () => apiClient.post('/auth/logout'),
}
