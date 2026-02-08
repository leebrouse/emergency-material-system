<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'

const username = ref('')
const password = ref('')
const role = ref('admin') // Simple selector for demo
const isLoading = ref(false)
const userStore = useUserStore()
const router = useRouter()

const handleLogin = async () => {
    isLoading.value = true
    try {
        // Simulate API call
        await new Promise(resolve => setTimeout(resolve, 1000))
        
        userStore.setToken('mock-jwt-token-12345')
        userStore.setRole(role.value)
        router.push('/dashboard')
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
          <h2 class="text-3xl font-bold text-white mb-2 text-center tracking-tight">System access</h2>
          <p class="text-gray-400 text-center mb-8 text-sm">Emergency Command Center</p>

          <form @submit.prevent="handleLogin" class="space-y-6">
              <div>
                  <label class="block text-sm font-medium text-gray-300 mb-1">Username</label>
                  <input v-model="username" type="text" class="w-full bg-gray-800/50 border border-gray-600 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-emerald-500 transition-all placeholder-gray-500" placeholder="Enter ID" required />
              </div>

               <div>
                  <label class="block text-sm font-medium text-gray-300 mb-1">Password</label>
                  <input v-model="password" type="password" class="w-full bg-gray-800/50 border border-gray-600 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-emerald-500 transition-all placeholder-gray-500" placeholder="••••••••" required />
              </div>

              <div>
                  <label class="block text-sm font-medium text-gray-300 mb-1">Role Simulation</label>
                  <select v-model="role" class="w-full bg-gray-800/50 border border-gray-600 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-emerald-500 transition-all">
                      <option value="admin">System Admin</option>
                      <option value="warehouse">Warehouse Manager</option>
                      <option value="rescue">Rescue Team</option>
                  </select>
              </div>

              <button 
                type="submit" 
                :disabled="isLoading"
                class="w-full bg-emerald-600 hover:bg-emerald-500 text-white font-bold py-3 rounded-lg shadow-lg transform transition-transform active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center"
              >
                  <span v-if="isLoading" class="animate-spin mr-2 h-5 w-5 border-2 border-white border-t-transparent rounded-full"></span>
                  {{ isLoading ? 'Authenticating...' : 'Initialize Session' }}
              </button>
          </form>
          
          <div class="mt-6 text-center">
              <span class="text-xs text-gray-500">Secure Connection • v2.4.0</span>
          </div>
      </div>
  </div>
</template>
