import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import { stockApi } from '@/api/stock';
export const useInventoryStore = defineStore('inventory', () => {
    const materials = ref([]);
    const isLoading = ref(false);
    const alertCount = computed(() => {
        return materials.value.filter(m => m.status === 'Critical' || m.status === 'Low').length;
    });
    const criticalItems = computed(() => {
        return materials.value.filter(m => m.status === 'Critical');
    });
    async function fetchMaterials() {
        isLoading.value = true;
        try {
            const res = await stockApi.getMaterials({ page: 1, page_size: 100 });
            const apiMaterials = res.data.data || []; // Backend returns { data: [...] }
            // Transform to internal model with status
            materials.value = apiMaterials.map(m => {
                const qty = m.quantity || 0;
                const min = m.min_stock || 10;
                let status = 'High';
                if (qty < min)
                    status = 'Critical';
                else if (qty < min * 2)
                    status = 'Low';
                return {
                    ...m,
                    status
                };
            });
        }
        catch (error) {
            console.error('Failed to fetch materials', error);
        }
        finally {
            isLoading.value = false;
        }
    }
    return { materials, alertCount, criticalItems, fetchMaterials, isLoading };
});
