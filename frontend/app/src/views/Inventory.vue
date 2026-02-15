<script setup lang="ts">
import { computed, ref, onMounted } from 'vue'
import { useInventoryStore, type Material } from '@/stores/inventory'
import { stockApi } from '@/api/stock'
import { ElMessage } from 'element-plus'
import { Plus, Search } from '@element-plus/icons-vue'

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
    currentItem.value = { 
        name: '', 
        category: '', 
        specs: '', 
        unit: '个', 
        quantity: 0, 
        min_stock: 10,
        description: '' 
    }
    dialogVisible.value = true
}

const handleEdit = (item: Material) => {
    isEditing.value = true
    currentItem.value = { ...item }
    dialogVisible.value = true
}

const handleDelete = async (id: number) => {
    try {
        await stockApi.deleteMaterial(id)
        await inventoryStore.fetchMaterials()
        ElMessage.success('删除成功')
    } catch (error) {
        console.error('Delete failed', error)
        ElMessage.error('删除失败')
    }
}

const saveItem = async () => {
    try {
        if (isEditing.value && currentItem.value.id) {
            // 1. Update material metadata (name, specs, min_stock, etc)
            await stockApi.updateMaterial(currentItem.value.id, currentItem.value)

            // 2. Handle inventory quantity changes
            const oldItem = inventoryStore.materials.find(m => m.id === currentItem.value.id)
            const oldQty = oldItem?.quantity || 0
            const newQty = currentItem.value.quantity || 0
            const delta = newQty - oldQty

            if (delta !== 0) {
                await stockApi.updateInventory({
                    material_id: currentItem.value.id,
                    quantity: Math.abs(delta),
                    operation: delta > 0 ? 'inbound' : 'outbound'
                })
            }
            
            await inventoryStore.fetchMaterials()
            ElMessage.success('更新成功')
        } else {
            // Add new
            const res = await stockApi.createMaterial(currentItem.value)
            const newMaterial = res.data
            
            // If initial quantity is set, do an inbound operation
            if (currentItem.value.quantity && currentItem.value.quantity > 0) {
                await stockApi.updateInventory({
                    material_id: newMaterial.id,
                    quantity: currentItem.value.quantity,
                    operation: 'inbound'
                })
            }
            
            await inventoryStore.fetchMaterials()
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
                    :prefix-icon="Search" 
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
                        <span :class="{'text-red-600 font-bold': scope.row.quantity < (scope.row.min_stock || 10)}">{{ scope.row.quantity }}</span>
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
                <el-form-item label="规格">
                    <el-input v-model="currentItem.specs" placeholder="例如: 50只/盒" />
                </el-form-item>
                <div class="grid grid-cols-3 gap-4">
                    <el-form-item label="单位">
                        <el-input v-model="currentItem.unit" placeholder="例如: 盒" />
                    </el-form-item>
                    <el-form-item label="数量">
                        <el-input-number v-model="currentItem.quantity" :min="0" class="w-full" :disabled="isEditing" />
                    </el-form-item>
                    <el-form-item label="安全库存">
                        <el-input-number v-model="currentItem.min_stock" :min="0" class="w-full" />
                    </el-form-item>
                </div>
                <el-form-item label="备注说明">
                    <el-input v-model="currentItem.description" type="textarea" placeholder="物资详细描述..." />
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
