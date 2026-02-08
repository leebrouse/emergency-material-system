import axios, { type AxiosInstance, type AxiosRequestConfig } from 'axios'
import { useUserStore } from '@/stores/user'
import { ElMessage } from 'element-plus'
import router from '@/router'

const apiClient: AxiosInstance = axios.create({
    baseURL: import.meta.env.VITE_API_URL || '/api',
    headers: {
        'Content-Type': 'application/json',
    },
})

// Request Interceptor
apiClient.interceptors.request.use(
    (config) => {
        const userStore = useUserStore()
        if (userStore.token) {
            config.headers.Authorization = `Bearer ${userStore.token}`
        }
        return config
    },
    (error) => {
        return Promise.reject(error)
    }
)

// Response Interceptor
apiClient.interceptors.response.use(
    (response) => response,
    (error) => {
        const userStore = useUserStore()
        if (error.response && error.response.status === 401) {
            userStore.removeToken()
            router.push('/login')
            ElMessage.error('Session expired, please login again')
        } else {
            ElMessage.error(error.message || 'Request failed')
        }
        return Promise.reject(error)
    }
)

export default apiClient
