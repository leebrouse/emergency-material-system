import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/login',
            name: 'login',
            component: () => import('@/views/Login.vue'),
        },
        {
            path: '/',
            name: 'home',
            component: () => import('@/layouts/MainLayout.vue'),
            redirect: '/dashboard',
            children: [
                {
                    path: 'dashboard',
                    name: 'dashboard',
                    component: () => import('@/views/Dashboard.vue'),
                    meta: { requiresAuth: true }
                },
                {
                    path: 'inventory',
                    name: 'inventory',
                    component: () => import('@/views/Inventory.vue'),
                    meta: { requiresAuth: true }
                },
                {
                    path: 'dispatch',
                    name: 'dispatch',
                    component: () => import('@/views/Dispatch.vue'),
                    meta: { requiresAuth: true }
                },
                {
                    path: 'logistics',
                    name: 'logistics',
                    component: () => import('@/views/Logistics.vue'),
                    meta: { requiresAuth: true }
                }
            ]
        }
    ]
})

router.beforeEach((to, _, next) => {
    const userStore = useUserStore()
    const publicPages = ['/login']
    const authRequired = !publicPages.includes(to.path)

    // Safety check: if pinia isn't updated yet, try localStorage
    const currentToken = userStore.token || localStorage.getItem('token')

    console.log(`[Router] Navigating to: ${to.path}, Auth Required: ${authRequired}, Token Present: ${!!currentToken}`)

    if (authRequired && !currentToken) {
        console.warn('[Router] No token found, redirecting to login')
        return next({ name: 'login' })
    }

    // Role Base Access Control
    const currentRole = userStore.role || localStorage.getItem('role')
    if (currentRole) {
        const restrictedRoutes: { [key: string]: string[] } = {
            'inventory': ['admin', 'warehouse'],
            'dispatch': ['admin', 'rescue'],
            'logistics': ['admin', 'rescue', 'warehouse']
        }

        const routeName = to.name as string
        if (restrictedRoutes[routeName] && !restrictedRoutes[routeName].includes(currentRole)) {
            console.warn(`Role ${currentRole} unauthorized for ${routeName}, redirecting...`)
            if (to.name !== 'dashboard') {
                return next({ name: 'dashboard' })
            }
        }
    }

    next()
})

export default router
