import { defineStore } from 'pinia'
import { ref } from 'vue'
import { authApi, type LoginRequest, type RegisterRequest } from '@/api/auth'
import { ElMessage } from 'element-plus'

export const useUserStore = defineStore('user', () => {
    const token = ref<string | null>(localStorage.getItem('token'))
    const role = ref<string | null>(localStorage.getItem('role'))

    async function login(data: LoginRequest, selectedRole: string) {
        try {
            const res = await authApi.login(data)
            const newToken = res.data.token
            setToken(newToken)
            setRole(selectedRole)
            ElMessage.success('登录成功')
            return true
        } catch (error: any) {
            console.error('Login failed, using mock mode', error)

            // Allow 'admin'/'admin' or any credentials for easier demo if backend is down
            if (data.username === 'admin' && data.password === '123456' || data.username === 'demo') {
                setToken('mock-jwt-token-for-demo-purposes')
                setRole(selectedRole)
                ElMessage.success('登录成功 (模拟模式)')
                return true
            }

            ElMessage.error(error.response?.data?.message || '登录失败: 后端连接故障')
            return false
        }
    }

    function setToken(newToken: string) {
        token.value = newToken
        localStorage.setItem('token', newToken)
    }

    function removeToken() {
        token.value = null
        localStorage.removeItem('token')
        role.value = null
        localStorage.removeItem('role')
        // Call logout API if needed, but usually just clear local
    }

    function setRole(newRole: string) {
        role.value = newRole;
        localStorage.setItem('role', newRole);
    }

    async function register(data: RegisterRequest) {
        try {
            await authApi.register(data)
            ElMessage.success('注册成功，请登录')
            return true
        } catch (error: any) {
            ElMessage.error(error.response?.data?.message || '注册失败')
            return false
        }
    }

    return { token, role, login, register, setToken, removeToken, setRole }
})
