<script setup lang="ts">
import { computed, ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { useInventoryStore } from '@/stores/inventory'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()
const inventoryStore = useInventoryStore()
const isMobileMenuOpen = ref(false)
const isCollapsed = ref(false)

const toggleSidebar = () => {
    isCollapsed.value = !isCollapsed.value
}

const menuItems = computed(() => {
  const role = userStore.role
  const items = [
    { title: '控制台', path: '/dashboard', icon: 'House', roles: ['admin', 'warehouse', 'rescue'] },
    { title: '物资管理', path: '/inventory', icon: 'Box', roles: ['admin', 'warehouse'] }, 
    { title: '调度指挥', path: '/dispatch', icon: 'Van', roles: ['admin', 'rescue'] },
    { title: '物流追踪', path: '/logistics', icon: 'MapLocation', roles: ['admin', 'rescue', 'warehouse'] }, 
  ]
  return items.filter(item => item.roles.includes(role || 'admin'))
})

const handleLogout = () => {
    userStore.removeToken()
    router.push('/login')
}

const alertCount = computed(() => inventoryStore.alertCount)

const getPageTitle = (name: string) => {
    const titles: { [key: string]: string } = {
        'dashboard': '控制台',
        'inventory': '物资管理',
        'dispatch': '调度指挥',
        'logistics': '物流追踪'
    }
    return titles[name] || name?.toUpperCase()
}
</script>

<template>
  <div class="h-screen w-screen grid grid-cols-1 grid-rows-[60px_1fr] transition-all duration-300"
       :class="isCollapsed ? 'md:grid-cols-[80px_1fr]' : 'md:grid-cols-[250px_1fr]'">
    <!-- Sidebar (Desktop) -->
    <aside class="bg-gray-900 text-white row-span-2 hidden md:flex flex-col border-r border-gray-800 transition-all duration-300 overflow-hidden">
      <div class="h-[60px] flex items-center px-4 border-b border-gray-800 overflow-hidden shrink-0"
           :class="isCollapsed ? 'justify-center' : 'justify-between'">
        <h1 v-if="!isCollapsed" class="text-xl font-bold text-emerald-500 tracking-wider truncate">应急物资系统</h1>
        <div v-else class="text-2xl font-bold text-emerald-500">EMS</div>
      </div>
      
      <nav class="flex-1 p-4 space-y-2">
        <router-link 
          v-for="item in menuItems" 
          :key="item.path" 
          :to="item.path"
          class="flex items-center p-3 rounded-xl hover:bg-gray-800/50 transition-all duration-300 group"
          :class="[
            { 'bg-emerald-900/20 text-emerald-400 border border-emerald-500/20': route.path === item.path, 'text-gray-400': route.path !== item.path },
            isCollapsed ? 'justify-center' : ''
          ]"
        >
          <el-icon class="transition-transform group-hover:scale-110 shrink-0" :class="isCollapsed ? '' : 'mr-3'" :size="20">
            <component :is="item.icon" />
          </el-icon>
          <span v-if="!isCollapsed" class="font-medium truncate transition-all duration-300">{{ item.title }}</span>
        </router-link>
      </nav>

      <div class="p-4 border-t border-gray-800 bg-gray-900/50">
        <div class="flex items-center" :class="isCollapsed ? 'justify-center' : ''">
            <div class="w-10 h-10 rounded-full bg-gradient-to-tr from-emerald-500 to-teal-400 flex items-center justify-center text-white font-bold shrink-0 shadow-lg"
                 :class="isCollapsed ? '' : 'mr-3'">
                {{ userStore.role?.[0].toUpperCase() }}
            </div>
            <div v-if="!isCollapsed" class="overflow-hidden">
                <p class="text-sm font-medium text-white truncate">
                    {{ userStore.role === 'admin' ? '系统管理员' : (userStore.role === 'warehouse' ? '仓库管理员' : '救援指挥员') }}
                </p>
                <button @click="handleLogout" class="text-xs text-gray-400 hover:text-emerald-400 transition-colors flex items-center mt-1">
                    <el-icon class="mr-1"><SwitchButton /></el-icon> 退出登录
                </button>
            </div>
        </div>
      </div>
    </aside>

    <!-- Mobile Drawer for Navigation -->
    <el-drawer v-model="isMobileMenuOpen" direction="ltr" size="70%" :with-header="false">
        <div class="bg-gray-900 h-full text-white p-4 flex flex-col">
            <div class="h-[60px] flex items-center justify-center border-b border-gray-800 mb-4">
                <h1 class="text-xl font-bold text-emerald-500">应急物资系统</h1>
            </div>
             <nav class="flex-1 space-y-2">
                <router-link 
                v-for="item in menuItems" 
                :key="item.path" 
                :to="item.path"
                @click="isMobileMenuOpen = false"
                class="flex items-center p-3 rounded-lg hover:bg-gray-800"
                :class="{ 'bg-gray-800 text-emerald-400': route.path === item.path }"
                >
                <el-icon class="mr-3" :size="20"><component :is="item.icon" /></el-icon>
                {{ item.title }}
                </router-link>
            </nav>
            <button @click="handleLogout" class="mt-auto w-full py-3 bg-gray-800 rounded text-center">退出登录</button>
        </div>
    </el-drawer>

    <!-- Header -->
    <header class="bg-white/80 backdrop-blur-md border-b border-gray-100 flex items-center justify-between px-6 shadow-sm z-10 sticky top-0">
      <div class="flex items-center">
          <button @click="isMobileMenuOpen = true" class="md:hidden mr-4 text-gray-600">
              <el-icon :size="24"><Menu /></el-icon>
          </button>
          <button @click="toggleSidebar" class="hidden md:flex mr-4 text-gray-500 hover:text-emerald-600 transition-colors p-1 rounded hover:bg-gray-100">
              <el-icon :size="20">
                  <component :is="isCollapsed ? 'Expand' : 'Fold'" />
              </el-icon>
          </button>
          <h2 class="text-lg font-bold text-gray-800 tracking-tight">{{ getPageTitle(route.name as string) }}</h2>
      </div>
      
      <div class="flex items-center space-x-6">
        <!-- Alert Component -->
        <el-popover placement="bottom" title="系统通知" :width="300" trigger="click">
            <template #reference>
                 <div class="relative cursor-pointer group">
                    <el-icon :size="24" class="text-gray-500 group-hover:text-emerald-600 transition-colors"><Bell /></el-icon>
                    <span v-if="alertCount > 0" class="absolute -top-1 -right-1 bg-red-500 text-white text-[10px] font-bold rounded-full h-4 w-4 flex items-center justify-center animate-pulse border-2 border-white">
                        {{ alertCount }}
                    </span>
                </div>
            </template>
            <div class="space-y-2">
                <p v-if="alertCount === 0" class="text-gray-500 text-sm">暂无新通知</p>
                <template v-else>
                     <div v-for="item in inventoryStore.criticalItems" :key="item.id" class="p-2 bg-red-50 rounded border border-red-100 flex items-start">
                        <el-icon class="text-red-500 mt-1 mr-2"><Warning /></el-icon>
                        <div>
                            <p class="text-xs font-bold text-red-800">库存预警</p>
                            <p class="text-xs text-gray-600">{{ item.name }} 库存告急 (剩余 {{ item.quantity }})</p>
                        </div>
                    </div>
                </template>
            </div>
        </el-popover>

        <div class="h-8 w-8 rounded-full bg-gray-200 overflow-hidden border border-gray-300">
            <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=Felix" alt="User" />
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="bg-slate-50 p-4 md:p-8 overflow-auto relative">
      <router-view v-slot="{ Component }">
        <transition name="fade" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </main>
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
