import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore('user', () => {
    const token = ref<string | null>(localStorage.getItem('token'))
    const role = ref<string | null>(localStorage.getItem('role'))

    function setToken(newToken: string) {
        token.value = newToken
        localStorage.setItem('token', newToken)
    }

    function removeToken() {
        token.value = null
        localStorage.removeItem('token')
        role.value = null
        localStorage.removeItem('role')
    }

    function setRole(newRole: string) {
        role.value = newRole;
        localStorage.setItem('role', newRole);
    }

    return { token, role, setToken, removeToken, setRole }
})
