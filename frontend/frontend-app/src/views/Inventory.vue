<script setup lang="ts">
import { computed, ref } from 'vue'
import { useInventoryStore } from '@/stores/inventory'
import { Plus, Search, Filter } from '@element-plus/icons-vue'
import type { Material } from '@/stores/inventory'

const inventoryStore = useInventoryStore()
const searchQuery = ref('')
const categoryFilter = ref('')
const dialogVisible = ref(false)
const isEditing = ref(false)
const currentItem = ref<Partial<Material>>({})

const categories = computed(() => {
    const cats = new Set(inventoryStore.materials.map(m => m.category))
    return Array.from(cats)
})

const filteredData = computed(() => {
    return inventoryStore.materials.filter(item => {
        const matchesSearch = item.name.toLowerCase().includes(searchQuery.value.toLowerCase()) || 
                              item.id.toString().includes(searchQuery.value)
        const matchesCategory = categoryFilter.value ? item.category === categoryFilter.value : true
        return matchesSearch && matchesCategory
    })
})

const getStatusType = (status: string) => {
    switch(status) {
        case 'High': return 'success'
        case 'Low': return 'warning'
        case 'Critical': return 'danger'
        default: return 'info'
    }
}

const handleAdd = () => {
    isEditing.value = false
    currentItem.value = { quantity: 0, status: 'High' }
    dialogVisible.value = true
}

const handleEdit = (item: Material) => {
    isEditing.value = true
    currentItem.value = { ...item }
    dialogVisible.value = true
}

const handleDelete = (id: number) => {
    // In a real app, delete from store
    const idx = inventoryStore.materials.findIndex(m => m.id === id)
    if (idx !== -1) inventoryStore.materials.splice(idx, 1)
}

const saveItem = () => {
    if (isEditing.value && currentItem.value.id) {
        const idx = inventoryStore.materials.findIndex(m => m.id === currentItem.value.id)
        if (idx !== -1) {
            inventoryStore.materials[idx] = currentItem.value as Material
            // Recalculate status simply based on quantity for demo
            if (currentItem.value.quantity! < 10) inventoryStore.materials[idx].status = 'Critical'
            else if (currentItem.value.quantity! < 50) inventoryStore.materials[idx].status = 'Low'
            else inventoryStore.materials[idx].status = 'High'
        }
    } else {
        const newId = Math.max(...inventoryStore.materials.map(m => m.id)) + 1
        const newItem = {
            ...currentItem.value,
            id: newId,
            status: currentItem.value.quantity! < 10 ? 'Critical' : (currentItem.value.quantity! < 50 ? 'Low' : 'High')
        } as Material
        inventoryStore.materials.push(newItem)
    }
    dialogVisible.value = false
}
</script>

<template>
    <div class="space-y-6">
        <div class="flex flex-col md:flex-row md:items-center justify-between gap-4">
            <div>
                 <h1 class="text-2xl font-bold text-gray-800">Inventory Management</h1>
                 <p class="text-gray-500 text-sm">Monitor stock levels and manage resources</p>
            </div>
           
            <div class="flex flex-col sm:flex-row gap-2">
                 <el-input 
                    v-model="searchQuery" 
                    placeholder="Search by name or ID..." 
                    class="w-full sm:w-64" 
                    prefix-icon="Search" 
                    clearable
                 />
                 
                 <el-select v-model="categoryFilter" placeholder="All Categories" clearable class="w-full sm:w-48">
                    <el-option v-for="cat in categories" :key="cat" :label="cat" :value="cat" />
                 </el-select>

                 <el-button type="primary" :icon="Plus" @click="handleAdd">Add Material</el-button>
            </div>
        </div>

        <el-card shadow="hover" class="border-gray-200 rounded-xl overflow-hidden">
            <el-table :data="filteredData" style="width: 100%" highlight-current-row>
                <el-table-column prop="id" label="ID" width="80" sortable />
                <el-table-column prop="name" label="Material Name" min-width="180">
                     <template #default="scope">
                        <span class="font-medium text-gray-700">{{ scope.row.name }}</span>
                    </template>
                </el-table-column>
                <el-table-column prop="category" label="Category" width="120">
                    <template #default="scope">
                        <el-tag size="small" type="info" effect="plain" round>{{ scope.row.category }}</el-tag>
                    </template>
                </el-table-column>
                <el-table-column prop="quantity" label="Quantity" sortable width="120">
                     <template #default="scope">
                        <span :class="{'text-red-600 font-bold': scope.row.quantity < 10}">{{ scope.row.quantity }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="Status" width="120">
                    <template #default="scope">
                        <div class="flex items-center gap-2">
                            <span class="h-2 w-2 rounded-full" :class="{
                                'bg-green-500': scope.row.status === 'High',
                                'bg-yellow-500': scope.row.status === 'Low',
                                'bg-red-500 animate-pulse': scope.row.status === 'Critical'
                            }"></span>
                            <span class="text-sm">{{ scope.row.status }}</span>
                        </div>
                    </template>
                </el-table-column>
                <el-table-column label="Actions" align="right" min-width="150">
                    <template #default="scope">
                        <el-button size="small" @click="handleEdit(scope.row)">Adjust</el-button>
                        <el-button size="small" type="danger" plain @click="handleDelete(scope.row.id)">Remove</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-card>

        <!-- Manage Dialog -->
        <el-dialog v-model="dialogVisible" :title="isEditing ? 'Adjust Inventory' : 'Add New Material'" width="500px">
            <el-form label-position="top">
                <el-form-item label="Material Name">
                    <el-input v-model="currentItem.name" placeholder="e.g. Surgical Masks" />
                </el-form-item>
                <el-form-item label="Category">
                    <el-select v-model="currentItem.category" placeholder="Select category" class="w-full" allow-create filterable default-first-option>
                         <el-option v-for="cat in categories" :key="cat" :label="cat" :value="cat" />
                    </el-select>
                </el-form-item>
                 <el-form-item label="Quantity">
                    <el-input-number v-model="currentItem.quantity" :min="0" class="w-full" />
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="dialogVisible = false">Cancel</el-button>
                    <el-button type="primary" @click="saveItem">Confirm</el-button>
                </span>
            </template>
         </el-dialog>
    </div>
</template>
