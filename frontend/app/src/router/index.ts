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

router.beforeEach((to, from, next) => {
    const userStore = useUserStore()
    const publicPages = ['/login']
    const authRequired = !publicPages.includes(to.path)

    if (authRequired && !userStore.token) {
        return next({ name: 'login' })
    }

    // Role Base Access Control
    if (userStore.role) {
        // Define restricted routes map
        const restrictedRoutes: { [key: string]: string[] } = {
            'inventory': ['admin', 'warehouse'],
            'dispatch': ['admin', 'rescue'],
            'logistics': ['admin', 'rescue', 'warehouse']
        }

        const routeName = to.name as string
        if (restrictedRoutes[routeName] && !restrictedRoutes[routeName].includes(userStore.role)) {
            // Redirect to authorized default page or error
            return next({ name: 'dashboard' }) // Dashboard is accessible to all roles in our logic
        }
    }

    next()
})

export default router
