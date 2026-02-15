<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'

const username = ref('')
const password = ref('')
const role = ref('admin')
const isLoading = ref(false)
const userStore = useUserStore()
const router = useRouter()

const handleLogin = async () => {
    isLoading.value = true
    try {
        const success = await userStore.login({ username: username.value, password: password.value }, role.value)
        if (success) {
            router.push('/dashboard')
        }
    } catch (error) {
        console.error('Login failed', error)
    } finally {
        isLoading.value = false
    }
}
</script>

<template>
  <div class="h-screen w-full flex items-center justify-center bg-gray-900 overflow-hidden relative">
      <!-- Background Effect -->
      <div class="absolute inset-0 bg-gradient-to-br from-gray-900 via-gray-800 to-black z-0"></div>
      <div class="absolute top-0 left-0 w-full h-full opacity-10 pointer-events-none" style="background-image: url('https://grainy-gradients.vercel.app/noise.svg');"></div>

      <div class="z-10 bg-white/10 backdrop-blur-lg border border-white/20 p-8 rounded-2xl shadow-2xl w-full max-w-md transform transition-all hover:scale-[1.01]">
          <h2 class="text-3xl font-bold text-white mb-2 text-center tracking-tight">系统访问</h2>
          <p class="text-gray-400 text-center mb-8 text-sm">应急物资调度指挥中心</p>

          <form @submit.prevent="handleLogin" class="space-y-6">
              <div>
                  <label class="block text-sm font-medium text-gray-300 mb-1">用户名</label>
                  <input v-model="username" type="text" class="w-full bg-gray-800/50 border border-gray-600 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-emerald-500 transition-all placeholder-gray-500" placeholder="请输入用户名" required />
              </div>

               <div>
                  <label class="block text-sm font-medium text-gray-300 mb-1">密码</label>
                  <input v-model="password" type="password" class="w-full bg-gray-800/50 border border-gray-600 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-emerald-500 transition-all placeholder-gray-500" placeholder="••••••••" required />
              </div>

              <div>
                  <label class="block text-sm font-medium text-gray-300 mb-1">角色模拟</label>
                  <select v-model="role" class="w-full bg-gray-800/50 border border-gray-600 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-emerald-500 transition-all">
                      <option value="admin">系统管理员</option>
                      <option value="warehouse">仓库管理员</option>
                      <option value="rescue">救援指挥员</option>
                  </select>
              </div>

              <button 
                type="submit" 
                :disabled="isLoading"
                class="w-full bg-emerald-600 hover:bg-emerald-500 text-white font-bold py-3 rounded-lg shadow-lg transform transition-transform active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center"
              >
                  <span v-if="isLoading" class="animate-spin mr-2 h-5 w-5 border-2 border-white border-t-transparent rounded-full"></span>
                  {{ isLoading ? '正在认证...' : '登录系统' }}
              </button>
          </form>
          
          <div class="mt-6 text-center flex flex-col items-center gap-2">
              <router-link to="/register" class="text-emerald-400 hover:text-emerald-300 text-sm transition-colors">没有账户？立即注册</router-link>
              <span class="text-[10px] text-gray-500 uppercase tracking-widest opacity-50">安全连接 • v2.4.0</span>
          </div>
      </div>
  </div>
</template>
