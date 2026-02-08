import { defineStore } from 'pinia'
import { ref } from 'vue'
import { authApi, type LoginRequest } from '@/api/auth'
import { ElMessage } from 'element-plus'

export const useUserStore = defineStore('user', () => {
    const token = ref<string | null>(localStorage.getItem('token'))
    const role = ref<string | null>(localStorage.getItem('role'))

    async function login(data: LoginRequest, selectedRole: string) {
        try {
            // In a real app, role might come from backend
            const res = await authApi.login(data)
            const newToken = res.data.token
            setToken(newToken)
            setRole(selectedRole)
            ElMessage.success('登录成功')
            return true
        } catch (error: any) {
            console.error('Login failed', error)
            ElMessage.error(error.response?.data?.message || '登录失败')
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

    return { token, role, login, setToken, removeToken, setRole }
})
