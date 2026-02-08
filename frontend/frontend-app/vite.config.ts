import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue()],
    resolve: {
        alias: {
            '@': path.resolve(__dirname, './src'),
        },
    },
    server: {
        proxy: {
            '/api/v1/auth': {
                target: 'http://localhost:8081',
                changeOrigin: true
            },
            '/api/v1/stock': {
                target: 'http://localhost:8082',
                changeOrigin: true
            },
            '/api/v1/dispatch': {
                target: 'http://localhost:8083',
                changeOrigin: true
            },
            '/api/v1/logistics': {
                target: 'http://localhost:8085',
                changeOrigin: true
            },
            '/api/v1/statistics': {
                target: 'http://localhost:8084',
                changeOrigin: true
            }
        }
    }
})
