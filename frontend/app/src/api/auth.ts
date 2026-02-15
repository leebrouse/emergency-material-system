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

export interface RegisterRequest {
    username: string
    password: string
    email: string
    phone: string
    roles: string[]
}

export const authApi = {
    login: (data: LoginRequest) => apiClient.post<LoginResponse>('/auth/login', data),
    register: (data: RegisterRequest) => apiClient.post('/auth/register', data),
    logout: () => apiClient.post('/auth/logout'),
}
