<script setup lang="ts">
import { computed, ref, onMounted } from 'vue'
import { useInventoryStore, type Material } from '@/stores/inventory'
import { stockApi } from '@/api/stock'
import { ElMessage } from 'element-plus'
import { Plus, Search, Filter } from '@element-plus/icons-vue'

const inventoryStore = useInventoryStore()
const searchQuery = ref('')
const categoryFilter = ref('')
const dialogVisible = ref(false)
const isEditing = ref(false)
const currentItem = ref<Partial<Material>>({})

onMounted(() => {
    inventoryStore.fetchMaterials()
})

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

const getStatusLabel = (status: string) => {
    switch(status) {
        case 'High': return '充足'
        case 'Low': return '待补货'
        case 'Critical': return '严重短缺'
        default: return status
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
    const idx = inventoryStore.materials.findIndex(m => m.id === id)
    if (idx !== -1) {
        inventoryStore.materials.splice(idx, 1)
        ElMessage.success('删除成功')
    }
}

const saveItem = async () => {
    try {
        if (isEditing.value && currentItem.value.id) {
            // Update via API
            await stockApi.updateInventory({
                material_id: currentItem.value.id,
                quantity: currentItem.value.quantity || 0,
                operation: 'set'
            })
            
            const idx = inventoryStore.materials.findIndex(m => m.id === currentItem.value.id)
            if (idx !== -1) {
                const qty = currentItem.value.quantity || 0
                inventoryStore.materials[idx] = {
                    ...currentItem.value as Material,
                    status: qty < 10 ? 'Critical' : (qty < 50 ? 'Low' : 'High')
                }
            }
            ElMessage.success('更新成功')
        } else {
            // Add new - would need POST /stock/materials endpoint
            const newId = Math.max(0, ...inventoryStore.materials.map(m => m.id)) + 1
            const qty = currentItem.value.quantity || 0
            const newItem: Material = {
                ...currentItem.value as Material,
                id: newId,
                status: qty < 10 ? 'Critical' : (qty < 50 ? 'Low' : 'High')
            }
            inventoryStore.materials.push(newItem)
            ElMessage.success('添加成功')
        }
        dialogVisible.value = false
    } catch (error) {
        console.error('Save failed', error)
        ElMessage.error('操作失败')
    }
}
</script>

<template>
    <div class="space-y-6">
        <div class="flex flex-col md:flex-row md:items-center justify-between gap-4">
            <div>
                 <h1 class="text-2xl font-bold text-gray-800">物资管理</h1>
                 <p class="text-gray-500 text-sm">监控库存水平，管理物资调配</p>
            </div>
           
            <div class="flex flex-col sm:flex-row gap-2">
                 <el-input 
                    v-model="searchQuery" 
                    placeholder="搜索物资名称或ID..." 
                    class="w-full sm:w-64" 
                    prefix-icon="Search" 
                    clearable
                 />
                 
                 <el-select v-model="categoryFilter" placeholder="全部分类" clearable class="w-full sm:w-48">
                    <el-option v-for="cat in categories" :key="cat" :label="cat" :value="cat" />
                 </el-select>

                 <el-button type="primary" :icon="Plus" @click="handleAdd">添加物资</el-button>
            </div>
        </div>

        <el-card shadow="hover" class="border-gray-200 rounded-xl overflow-hidden">
            <el-table :data="filteredData" style="width: 100%" highlight-current-row v-loading="inventoryStore.isLoading">
                <el-table-column prop="id" label="ID" width="80" sortable />
                <el-table-column prop="name" label="物资名称" min-width="180">
                     <template #default="scope">
                        <span class="font-medium text-gray-700">{{ scope.row.name }}</span>
                    </template>
                </el-table-column>
                <el-table-column prop="category" label="分类" width="120">
                    <template #default="scope">
                        <el-tag size="small" type="info" effect="plain" round>{{ scope.row.category }}</el-tag>
                    </template>
                </el-table-column>
                <el-table-column prop="quantity" label="库存数量" sortable width="120">
                     <template #default="scope">
                        <span :class="{'text-red-600 font-bold': scope.row.quantity < 10}">{{ scope.row.quantity }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="状态" width="120">
                    <template #default="scope">
                        <div class="flex items-center gap-2">
                            <span class="h-2 w-2 rounded-full" :class="{
                                'bg-green-500': scope.row.status === 'High',
                                'bg-yellow-500': scope.row.status === 'Low',
                                'bg-red-500 animate-pulse': scope.row.status === 'Critical'
                            }"></span>
                            <span class="text-sm">{{ getStatusLabel(scope.row.status) }}</span>
                        </div>
                    </template>
                </el-table-column>
                <el-table-column label="操作" align="right" min-width="150">
                    <template #default="scope">
                        <el-button size="small" @click="handleEdit(scope.row)">调整</el-button>
                        <el-button size="small" type="danger" plain @click="handleDelete(scope.row.id)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-card>

        <!-- Manage Dialog -->
        <el-dialog v-model="dialogVisible" :title="isEditing ? '调整库存' : '添加新物资'" width="500px">
            <el-form label-position="top">
                <el-form-item label="物资名称">
                    <el-input v-model="currentItem.name" placeholder="例如: 医用口罩 (N95)" />
                </el-form-item>
                <el-form-item label="分类">
                    <el-select v-model="currentItem.category" placeholder="选择分类" class="w-full" allow-create filterable default-first-option>
                         <el-option v-for="cat in categories" :key="cat" :label="cat" :value="cat" />
                    </el-select>
                </el-form-item>
                 <el-form-item label="数量">
                    <el-input-number v-model="currentItem.quantity" :min="0" class="w-full" />
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="dialogVisible = false">取消</el-button>
                    <el-button type="primary" @click="saveItem">确认</el-button>
                </span>
            </template>
         </el-dialog>
    </div>
</template>
