<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { ElMessage } from 'element-plus'

const username = ref('')
const password = ref('')
const confirmPassword = ref('')
const email = ref('')
const phone = ref('')
const roles = ref(['guest'])
const isLoading = ref(false)

const userStore = useUserStore()
const router = useRouter()

const handleRegister = async () => {
    if (password.value !== confirmPassword.value) {
        ElMessage.error('两次输入的密码不一致')
        return
    }

    isLoading.value = true
    try {
        const success = await userStore.register({
            username: username.value,
            password: password.value,
            email: email.value,
            phone: phone.value,
            roles: roles.value
        })
        if (success) {
            router.push('/login')
        }
    } catch (error) {
        console.error('Registration failed', error)
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

      <div class="z-10 bg-white/10 backdrop-blur-lg border border-white/20 p-8 rounded-2xl shadow-2xl w-full max-w-lg transform transition-all">
          <h2 class="text-3xl font-bold text-white mb-2 text-center tracking-tight">账户注册</h2>
          <p class="text-gray-400 text-center mb-8 text-sm">创建您的应急指挥中心访问权限</p>

          <form @submit.prevent="handleRegister" class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="md:col-span-2">
                  <label class="block text-sm font-medium text-gray-300 mb-1">用户名</label>
                  <input v-model="username" type="text" class="w-full bg-gray-800/50 border border-gray-600 rounded-lg px-4 py-2.5 text-white focus:outline-none focus:ring-2 focus:ring-emerald-500 transition-all" placeholder="请输入用户名" required />
              </div>

               <div>
                  <label class="block text-sm font-medium text-gray-300 mb-1">密码</label>
                  <input v-model="password" type="password" class="w-full bg-gray-800/50 border border-gray-600 rounded-lg px-4 py-2.5 text-white focus:outline-none focus:ring-2 focus:ring-emerald-500 transition-all" placeholder="••••••••" required />
              </div>

              <div>
                  <label class="block text-sm font-medium text-gray-300 mb-1">确认密码</label>
                  <input v-model="confirmPassword" type="password" class="w-full bg-gray-800/50 border border-gray-600 rounded-lg px-4 py-2.5 text-white focus:outline-none focus:ring-2 focus:ring-emerald-500 transition-all" placeholder="••••••••" required />
              </div>

               <div>
                <label class="block text-sm font-medium text-gray-300 mb-1">邮箱</label>
                <input v-model="email" type="email" class="w-full bg-gray-800/50 border border-gray-600 rounded-lg px-4 py-2.5 text-white focus:outline-none focus:ring-2 focus:ring-emerald-500 transition-all" placeholder="example@mail.com" required />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-300 mb-1">电话</label>
                <input v-model="phone" type="text" class="w-full bg-gray-800/50 border border-gray-600 rounded-lg px-4 py-2.5 text-white focus:outline-none focus:ring-2 focus:ring-emerald-500 transition-all" placeholder="请输入电话号码" required />
              </div>

               <div class="md:col-span-2">
                  <label class="block text-sm font-medium text-gray-300 mb-1">初始角色</label>
                  <select v-model="roles[0]" class="w-full bg-gray-800/50 border border-gray-600 rounded-lg px-4 py-2.5 text-white focus:outline-none focus:ring-2 focus:ring-emerald-500 transition-all">
                      <option value="admin">系统管理员</option>
                      <option value="warehouse">仓库管理员</option>
                      <option value="rescue">救援指挥员</option>
                  </select>
              </div>

              <div class="md:col-span-2 mt-4">
                <button 
                  type="submit" 
                  :disabled="isLoading"
                  class="w-full bg-emerald-600 hover:bg-emerald-500 text-white font-bold py-3 rounded-lg shadow-lg transform transition-transform active:scale-95 disabled:opacity-50 flex items-center justify-center"
                >
                    <span v-if="isLoading" class="animate-spin mr-2 h-5 w-5 border-2 border-white border-t-transparent rounded-full"></span>
                    {{ isLoading ? '注册中...' : '注册账户' }}
                </button>
              </div>
          </form>
          
          <div class="mt-6 text-center">
              <router-link to="/login" class="text-emerald-400 hover:text-emerald-300 text-sm transition-colors">已有账户？立即登录</router-link>
          </div>
      </div>
  </div>
</template>
