import { useState } from 'react';
import { Search, Plus, Download, Package, Calendar, AlertCircle, ChevronRight } from 'lucide-react';

interface Material {
  id: string;
  name: string;
  category: string;
  specification: string;
  quantity: number;
  unit: string;
  location: string;
  batch: string;
  expiryDate: string;
  status: 'normal' | 'low' | 'expired';
}

export function Inventory() {
  const [searchTerm, setSearchTerm] = useState('');
  const [selectedCategory, setSelectedCategory] = useState('all');
  const [showAddDialog, setShowAddDialog] = useState(false);

  const materials: Material[] = [
    {
      id: 'M001',
      name: '医疗急救包',
      category: '医疗用品',
      specification: '标准型',
      quantity: 580,
      unit: '套',
      location: 'A1-01',
      batch: 'B2024011',
      expiryDate: '2027-01-15',
      status: 'normal'
    },
    {
      id: 'M002',
      name: '帐篷（4人）',
      category: '应急设施',
      specification: '防水防风',
      quantity: 120,
      unit: '顶',
      location: 'B2-03',
      batch: 'B2024025',
      expiryDate: '2030-06-20',
      status: 'normal'
    },
    {
      id: 'M003',
      name: '应急食品包',
      category: '食品',
      specification: '7天装',
      quantity: 45,
      unit: '箱',
      location: 'C1-05',
      batch: 'B2024008',
      expiryDate: '2026-03-10',
      status: 'low'
    },
    {
      id: 'M004',
      name: '饮用水',
      category: '食品',
      specification: '500ml*24瓶',
      quantity: 850,
      unit: '箱',
      location: 'C1-02',
      batch: 'B2024030',
      expiryDate: '2026-12-31',
      status: 'normal'
    },
    {
      id: 'M005',
      name: '发电机',
      category: '设备',
      specification: '5KW汽油',
      quantity: 25,
      unit: '台',
      location: 'D1-01',
      batch: 'B2023088',
      expiryDate: '2028-08-15',
      status: 'normal'
    },
    {
      id: 'M006',
      name: '消毒液',
      category: '医疗用品',
      specification: '5L装',
      quantity: 15,
      unit: '桶',
      location: 'A1-08',
      batch: 'B2023125',
      expiryDate: '2026-01-20',
      status: 'expired'
    }
  ];

  const categories = ['全部', '医疗用品', '食品', '应急设施', '设备'];

  const getStatusBadge = (status: string) => {
    switch (status) {
      case 'normal':
        return <span className="px-2 py-1 bg-green-100 text-green-700 rounded text-xs font-medium">正常</span>;
      case 'low':
        return <span className="px-2 py-1 bg-orange-100 text-orange-700 rounded text-xs font-medium">库存低</span>;
      case 'expired':
        return <span className="px-2 py-1 bg-red-100 text-red-700 rounded text-xs font-medium">已过期</span>;
      default:
        return null;
    }
  };

  const filteredMaterials = materials.filter(material => {
    const matchesSearch = material.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
                         material.id.toLowerCase().includes(searchTerm.toLowerCase());
    const matchesCategory = selectedCategory === 'all' || material.category === selectedCategory;
    return matchesSearch && matchesCategory;
  });

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-bold text-slate-800">物资库存管理</h1>
          <p className="text-slate-500 mt-1">实时库存监控与管理</p>
        </div>
        <div className="flex gap-3">
          <button className="px-4 py-2 border border-slate-300 text-slate-700 rounded-lg hover:bg-slate-50 transition-colors flex items-center gap-2">
            <Download className="w-4 h-4" />
            导出数据
          </button>
          <button 
            onClick={() => setShowAddDialog(true)}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2"
          >
            <Plus className="w-4 h-4" />
            添加物资
          </button>
        </div>
      </div>

      {/* Filters */}
      <div className="bg-white rounded-xl p-6 shadow-sm border border-slate-200">
        <div className="flex flex-col lg:flex-row gap-4">
          {/* Search */}
          <div className="flex-1 relative">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400" />
            <input
              type="text"
              placeholder="搜索物资名称或编号..."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              className="w-full pl-10 pr-4 py-2 border border-slate-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>

          {/* Category Filter */}
          <div className="flex gap-2 flex-wrap">
            {categories.map((category) => (
              <button
                key={category}
                onClick={() => setSelectedCategory(category === '全部' ? 'all' : category)}
                className={`px-4 py-2 rounded-lg font-medium transition-colors ${
                  (category === '全部' && selectedCategory === 'all') || selectedCategory === category
                    ? 'bg-blue-600 text-white'
                    : 'bg-slate-100 text-slate-700 hover:bg-slate-200'
                }`}
              >
                {category}
              </button>
            ))}
          </div>
        </div>
      </div>

      {/* Inventory Table */}
      <div className="bg-white rounded-xl shadow-sm border border-slate-200 overflow-hidden">
        <div className="overflow-x-auto">
          <table className="w-full">
            <thead className="bg-slate-50 border-b border-slate-200">
              <tr>
                <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">物资编号</th>
                <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">物资名称</th>
                <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">类别</th>
                <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">规格</th>
                <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">库存数量</th>
                <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">存储位置</th>
                <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">批次号</th>
                <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">有效期</th>
                <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">状态</th>
                <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">操作</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-slate-200">
              {filteredMaterials.map((material) => (
                <tr key={material.id} className="hover:bg-slate-50 transition-colors">
                  <td className="px-6 py-4 text-sm font-medium text-slate-800">{material.id}</td>
                  <td className="px-6 py-4">
                    <div className="flex items-center gap-3">
                      <div className="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center">
                        <Package className="w-5 h-5 text-blue-600" />
                      </div>
                      <span className="font-medium text-slate-800">{material.name}</span>
                    </div>
                  </td>
                  <td className="px-6 py-4 text-sm text-slate-600">{material.category}</td>
                  <td className="px-6 py-4 text-sm text-slate-600">{material.specification}</td>
                  <td className="px-6 py-4">
                    <span className={`font-semibold ${
                      material.status === 'low' ? 'text-orange-600' : 'text-slate-800'
                    }`}>
                      {material.quantity} {material.unit}
                    </span>
                  </td>
                  <td className="px-6 py-4 text-sm text-slate-600">{material.location}</td>
                  <td className="px-6 py-4 text-sm text-slate-600">{material.batch}</td>
                  <td className="px-6 py-4">
                    <div className="flex items-center gap-2 text-sm text-slate-600">
                      <Calendar className="w-4 h-4" />
                      {material.expiryDate}
                    </div>
                  </td>
                  <td className="px-6 py-4">{getStatusBadge(material.status)}</td>
                  <td className="px-6 py-4">
                    <button className="text-blue-600 hover:text-blue-700 flex items-center gap-1 text-sm font-medium">
                      详情
                      <ChevronRight className="w-4 h-4" />
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>

        {/* Empty State */}
        {filteredMaterials.length === 0 && (
          <div className="py-12 text-center">
            <Package className="w-12 h-12 text-slate-300 mx-auto mb-3" />
            <p className="text-slate-500">未找到相关物资</p>
          </div>
        )}
      </div>

      {/* Summary Stats */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div className="bg-white rounded-xl p-6 shadow-sm border border-slate-200">
          <div className="flex items-center gap-3">
            <div className="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center">
              <Package className="w-6 h-6 text-blue-600" />
            </div>
            <div>
              <p className="text-sm text-slate-600">物资种类</p>
              <p className="text-2xl font-bold text-slate-800">{materials.length}</p>
            </div>
          </div>
        </div>

        <div className="bg-white rounded-xl p-6 shadow-sm border border-slate-200">
          <div className="flex items-center gap-3">
            <div className="w-12 h-12 bg-orange-100 rounded-lg flex items-center justify-center">
              <AlertCircle className="w-6 h-6 text-orange-600" />
            </div>
            <div>
              <p className="text-sm text-slate-600">库存预警</p>
              <p className="text-2xl font-bold text-slate-800">
                {materials.filter(m => m.status === 'low' || m.status === 'expired').length}
              </p>
            </div>
          </div>
        </div>

        <div className="bg-white rounded-xl p-6 shadow-sm border border-slate-200">
          <div className="flex items-center gap-3">
            <div className="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center">
              <Package className="w-6 h-6 text-green-600" />
            </div>
            <div>
              <p className="text-sm text-slate-600">总库存量</p>
              <p className="text-2xl font-bold text-slate-800">
                {materials.reduce((sum, m) => sum + m.quantity, 0)}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
