import { useState } from 'react';
import { AlertCircle, CheckCircle, Clock, XCircle, Plus, MapPin, User } from 'lucide-react';

interface Demand {
  id: string;
  title: string;
  location: string;
  applicant: string;
  urgency: 'high' | 'medium' | 'low';
  status: 'pending' | 'approved' | 'allocated' | 'completed' | 'rejected';
  items: { name: string; quantity: number; unit: string }[];
  createdAt: string;
  description: string;
}

export function Demands() {
  const [selectedStatus, setSelectedStatus] = useState<string>('all');

  const demands: Demand[] = [
    {
      id: 'D001',
      title: '地震灾区应急物资申请',
      location: '四川省雅安市',
      applicant: '张伟（救援队长）',
      urgency: 'high',
      status: 'pending',
      items: [
        { name: '帐篷', quantity: 50, unit: '顶' },
        { name: '应急食品', quantity: 200, unit: '份' },
        { name: '饮用水', quantity: 100, unit: '箱' }
      ],
      createdAt: '2026-02-08 09:30',
      description: '雅安地区发生6.5级地震，需要紧急救援物资支持'
    },
    {
      id: 'D002',
      title: '洪涝灾害救援物资需求',
      location: '湖南省长沙市',
      applicant: '李明（应急办）',
      urgency: 'high',
      status: 'approved',
      items: [
        { name: '冲锋舟', quantity: 5, unit: '艘' },
        { name: '救生衣', quantity: 100, unit: '件' },
        { name: '应急灯', quantity: 50, unit: '个' }
      ],
      createdAt: '2026-02-08 08:15',
      description: '持续强降雨导致多处内涝，需要水上救援设备'
    },
    {
      id: 'D003',
      title: '疫情防控物资补充',
      location: '北京市朝阳区',
      applicant: '王芳（社区负责人）',
      urgency: 'medium',
      status: 'allocated',
      items: [
        { name: '医用口罩', quantity: 5000, unit: '只' },
        { name: '消毒液', quantity: 30, unit: '桶' },
        { name: '体温计', quantity: 20, unit: '支' }
      ],
      createdAt: '2026-02-07 16:20',
      description: '社区防疫物资库存不足，需要及时补充'
    },
    {
      id: 'D004',
      title: '山火救援设备申请',
      location: '云南省丽江市',
      applicant: '赵强（消防支队）',
      urgency: 'high',
      status: 'completed',
      items: [
        { name: '灭火器', quantity: 100, unit: '个' },
        { name: '防护服', quantity: 50, unit: '套' },
        { name: '应急水箱', quantity: 10, unit: '个' }
      ],
      createdAt: '2026-02-06 14:00',
      description: '山火蔓延，需要大量灭火和防护设备'
    },
    {
      id: 'D005',
      title: '冬季取暖物资需求',
      location: '黑龙江省哈尔滨市',
      applicant: '刘洋（民政局）',
      urgency: 'low',
      status: 'rejected',
      items: [
        { name: '电暖器', quantity: 30, unit: '台' },
        { name: '棉被', quantity: 100, unit: '床' }
      ],
      createdAt: '2026-02-05 10:30',
      description: '部分困难群众冬季取暖物资不足'
    }
  ];

  const getUrgencyBadge = (urgency: string) => {
    switch (urgency) {
      case 'high':
        return <span className="px-2 py-1 bg-red-100 text-red-700 rounded text-xs font-medium">紧急</span>;
      case 'medium':
        return <span className="px-2 py-1 bg-orange-100 text-orange-700 rounded text-xs font-medium">一般</span>;
      case 'low':
        return <span className="px-2 py-1 bg-blue-100 text-blue-700 rounded text-xs font-medium">低</span>;
      default:
        return null;
    }
  };

  const getStatusInfo = (status: string) => {
    switch (status) {
      case 'pending':
        return { label: '待审核', icon: Clock, color: 'text-orange-600 bg-orange-100' };
      case 'approved':
        return { label: '已批准', icon: CheckCircle, color: 'text-green-600 bg-green-100' };
      case 'allocated':
        return { label: '已分配', icon: CheckCircle, color: 'text-blue-600 bg-blue-100' };
      case 'completed':
        return { label: '已完成', icon: CheckCircle, color: 'text-slate-600 bg-slate-100' };
      case 'rejected':
        return { label: '已拒绝', icon: XCircle, color: 'text-red-600 bg-red-100' };
      default:
        return { label: status, icon: Clock, color: 'text-slate-600 bg-slate-100' };
    }
  };

  const statusFilters = [
    { value: 'all', label: '全部', count: demands.length },
    { value: 'pending', label: '待审核', count: demands.filter(d => d.status === 'pending').length },
    { value: 'approved', label: '已批准', count: demands.filter(d => d.status === 'approved').length },
    { value: 'allocated', label: '已分配', count: demands.filter(d => d.status === 'allocated').length },
    { value: 'completed', label: '已完成', count: demands.filter(d => d.status === 'completed').length }
  ];

  const filteredDemands = selectedStatus === 'all' 
    ? demands 
    : demands.filter(d => d.status === selectedStatus);

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-bold text-slate-800">需求管理</h1>
          <p className="text-slate-500 mt-1">物资需求申报与分配调度</p>
        </div>
        <button className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2">
          <Plus className="w-4 h-4" />
          新建需求
        </button>
      </div>

      {/* Status Filters */}
      <div className="bg-white rounded-xl p-6 shadow-sm border border-slate-200">
        <div className="flex gap-3 flex-wrap">
          {statusFilters.map((filter) => (
            <button
              key={filter.value}
              onClick={() => setSelectedStatus(filter.value)}
              className={`px-4 py-2 rounded-lg font-medium transition-all ${
                selectedStatus === filter.value
                  ? 'bg-blue-600 text-white shadow-sm'
                  : 'bg-slate-100 text-slate-700 hover:bg-slate-200'
              }`}
            >
              {filter.label}
              <span className={`ml-2 px-2 py-0.5 rounded text-xs ${
                selectedStatus === filter.value
                  ? 'bg-white/20'
                  : 'bg-slate-200'
              }`}>
                {filter.count}
              </span>
            </button>
          ))}
        </div>
      </div>

      {/* Demands List */}
      <div className="space-y-4">
        {filteredDemands.map((demand) => {
          const statusInfo = getStatusInfo(demand.status);
          const StatusIcon = statusInfo.icon;

          return (
            <div key={demand.id} className="bg-white rounded-xl shadow-sm border border-slate-200 overflow-hidden hover:shadow-md transition-shadow">
              {/* Header */}
              <div className="p-6 border-b border-slate-200">
                <div className="flex items-start justify-between">
                  <div className="flex-1">
                    <div className="flex items-center gap-3 mb-2">
                      <h3 className="text-lg font-semibold text-slate-800">{demand.title}</h3>
                      {getUrgencyBadge(demand.urgency)}
                      <span className={`px-3 py-1 rounded-full text-xs font-medium flex items-center gap-1 ${statusInfo.color}`}>
                        <StatusIcon className="w-3 h-3" />
                        {statusInfo.label}
                      </span>
                    </div>
                    <div className="flex items-center gap-4 text-sm text-slate-600">
                      <div className="flex items-center gap-1">
                        <MapPin className="w-4 h-4" />
                        {demand.location}
                      </div>
                      <div className="flex items-center gap-1">
                        <User className="w-4 h-4" />
                        {demand.applicant}
                      </div>
                      <div className="flex items-center gap-1">
                        <Clock className="w-4 h-4" />
                        {demand.createdAt}
                      </div>
                    </div>
                  </div>
                  <span className="text-sm font-medium text-slate-500">#{demand.id}</span>
                </div>
              </div>

              {/* Content */}
              <div className="p-6">
                <p className="text-slate-600 mb-4">{demand.description}</p>
                
                {/* Items List */}
                <div className="bg-slate-50 rounded-lg p-4">
                  <h4 className="text-sm font-semibold text-slate-700 mb-3">申请物资清单</h4>
                  <div className="grid grid-cols-1 md:grid-cols-3 gap-3">
                    {demand.items.map((item, idx) => (
                      <div key={idx} className="bg-white rounded-lg p-3 flex items-center justify-between">
                        <span className="text-slate-700">{item.name}</span>
                        <span className="font-semibold text-blue-600">
                          {item.quantity} {item.unit}
                        </span>
                      </div>
                    ))}
                  </div>
                </div>

                {/* Actions */}
                <div className="flex gap-3 mt-4">
                  {demand.status === 'pending' && (
                    <>
                      <button className="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors text-sm font-medium">
                        批准申请
                      </button>
                      <button className="px-4 py-2 border border-slate-300 text-slate-700 rounded-lg hover:bg-slate-50 transition-colors text-sm font-medium">
                        拒绝申请
                      </button>
                    </>
                  )}
                  {demand.status === 'approved' && (
                    <button className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors text-sm font-medium">
                      分配物资
                    </button>
                  )}
                  {demand.status === 'allocated' && (
                    <button className="px-4 py-2 bg-purple-600 text-white rounded-lg hover:bg-purple-700 transition-colors text-sm font-medium">
                      查看配送
                    </button>
                  )}
                  <button className="px-4 py-2 border border-slate-300 text-slate-700 rounded-lg hover:bg-slate-50 transition-colors text-sm font-medium">
                    查看详情
                  </button>
                </div>
              </div>
            </div>
          );
        })}
      </div>

      {/* Empty State */}
      {filteredDemands.length === 0 && (
        <div className="bg-white rounded-xl shadow-sm border border-slate-200 py-12">
          <div className="text-center">
            <AlertCircle className="w-12 h-12 text-slate-300 mx-auto mb-3" />
            <p className="text-slate-500">暂无相关需求</p>
          </div>
        </div>
      )}
    </div>
  );
}
