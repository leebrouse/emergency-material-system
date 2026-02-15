import axios, { type AxiosInstance } from 'axios'
import { useUserStore } from '@/stores/user'
import { ElMessage } from 'element-plus'
import router from '@/router'

const apiClient: AxiosInstance = axios.create({
    baseURL: import.meta.env.VITE_API_URL || '/api/v1',
    headers: {
        'Content-Type': 'application/json',
    },
})

// Request Interceptor
apiClient.interceptors.request.use(
    (config) => {
        const userStore = useUserStore()

        // Log request
        console.groupCollapsed(`%cAPI Request: ${config.method?.toUpperCase()} ${config.url}`, 'color: #3b82f6; font-weight: bold;')
        console.log('BaseURL:', config.baseURL)
        console.log('Headers:', config.headers)
        if (config.data) console.log('Body:', config.data)
        if (config.params) console.log('Params:', config.params)
        console.groupEnd()

        if (userStore.token) {
            config.headers.Authorization = `Bearer ${userStore.token}`
        }
        return config
    },
    (error) => {
        console.error('%cRequest Error:', 'color: #ef4444; font-weight: bold;', error)
        return Promise.reject(error)
    }
)

// Response Interceptor
apiClient.interceptors.response.use(
    (response) => {
        // Log response
        console.log(
            `%cAPI Response: ${response.status} ${response.config.url}`,
            'color: #10b981; font-weight: bold;',
            response.data
        )
        return response
    },
    (error) => {
        const userStore = useUserStore()

        // Log error details
        console.group(`%cAPI Error: ${error.response?.status || 'Network'} ${error.config?.url}`, 'color: #ef4444; font-weight: bold;')
        console.log('Message:', error.message)
        if (error.response?.data) console.log('Data:', error.response.data)
        console.groupEnd()

        if (error.response && error.response.status === 401) {
            userStore.removeToken()
            router.push('/login')
            ElMessage.error('Session expired, please login again')
        } else {
            ElMessage.error(error.response?.data?.error || error.message || 'Request failed')
        }
        return Promise.reject(error)
    }
)

export default apiClient
