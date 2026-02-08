import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export interface Material {
    id: number
    name: string
    category: string
    quantity: number
    status: 'High' | 'Low' | 'Critical'
}

export const useInventoryStore = defineStore('inventory', () => {
    const materials = ref<Material[]>([
        { id: 1, name: 'Surgical Masks', category: 'Medical', quantity: 5000, status: 'High' },
        { id: 2, name: 'Bottled Water', category: 'Food', quantity: 1200, status: 'High' },
        { id: 3, name: 'Tents', category: 'Shelter', quantity: 45, status: 'Low' },
        { id: 4, name: 'Generators', category: 'Tools', quantity: 5, status: 'Critical' },
        { id: 5, name: 'First Aid Kits', category: 'Medical', quantity: 15, status: 'Critical' },
    ])

    const alertCount = computed(() => {
        return materials.value.filter(m => m.status === 'Critical' || m.status === 'Low').length
    })

    const criticalItems = computed(() => {
        return materials.value.filter(m => m.status === 'Critical')
    })

    function updateStatus(id: number, newStatus: 'High' | 'Low' | 'Critical') {
        const item = materials.value.find(m => m.id === id)
        if (item) {
            item.status = newStatus
        }
    }

    return { materials, alertCount, criticalItems, updateStatus }
})
