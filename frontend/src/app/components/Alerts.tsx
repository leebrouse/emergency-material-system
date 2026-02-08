import { AlertTriangle, Calendar, Bell, TrendingDown, Package } from 'lucide-react';

interface Alert {
  id: string;
  type: 'low_stock' | 'expiring' | 'expired';
  material: string;
  category: string;
  currentQuantity: number;
  threshold?: number;
  expiryDate?: string;
  location: string;
  severity: 'high' | 'medium' | 'low';
  createdAt: string;
}

export function Alerts() {
  const alerts: Alert[] = [
    {
      id: 'A001',
      type: 'low_stock',
      material: '应急食品包',
      category: '食品',
      currentQuantity: 45,
      threshold: 100,
      location: 'C1-05',
      severity: 'high',
      createdAt: '2026-02-08 10:30'
    },
    {
      id: 'A002',
      type: 'expired',
      material: '消毒液',
      category: '医疗用品',
      currentQuantity: 15,
      expiryDate: '2026-01-20',
      location: 'A1-08',
      severity: 'high',
      createdAt: '2026-02-08 09:00'
    },
    {
      id: 'A003',
      type: 'expiring',
      material: '应急食品包',
      category: '食品',
      currentQuantity: 45,
      expiryDate: '2026-03-10',
      location: 'C1-05',
      severity: 'medium',
      createdAt: '2026-02-08 08:15'
    },
    {
      id: 'A004',
      type: 'low_stock',
      material: '医用口罩',
      category: '医疗用品',
      currentQuantity: 850,
      threshold: 2000,
      location: 'A1-03',
      severity: 'medium',
      createdAt: '2026-02-07 16:45'
    },
    {
      id: 'A005',
      type: 'low_stock',
      material: '帐篷（4人）',
      category: '应急设施',
      currentQuantity: 120,
      threshold: 200,
      location: 'B2-03',
      severity: 'low',
      createdAt: '2026-02-07 14:20'
    },
    {
      id: 'A006',
      type: 'expiring',
      material: '医疗急救包',
      category: '医疗用品',
      currentQuantity: 580,
      expiryDate: '2026-04-15',
      location: 'A1-01',
      severity: 'low',
      createdAt: '2026-02-07 11:00'
    }
  ];

  const getAlertIcon = (type: string) => {
    switch (type) {
      case 'low_stock':
        return TrendingDown;
      case 'expiring':
      case 'expired':
        return Calendar;
      default:
        return AlertTriangle;
    }
  };

  const getAlertColor = (severity: string) => {
    switch (severity) {
      case 'high':
        return 'bg-red-50 border-red-200';
      case 'medium':
        return 'bg-orange-50 border-orange-200';
      case 'low':
        return 'bg-yellow-50 border-yellow-200';
      default:
        return 'bg-slate-50 border-slate-200';
    }
  };

  const getSeverityBadge = (severity: string) => {
    switch (severity) {
      case 'high':
        return <span className="px-2 py-1 bg-red-100 text-red-700 rounded text-xs font-medium">高</span>;
      case 'medium':
        return <span className="px-2 py-1 bg-orange-100 text-orange-700 rounded text-xs font-medium">中</span>;
      case 'low':
        return <span className="px-2 py-1 bg-yellow-100 text-yellow-700 rounded text-xs font-medium">低</span>;
      default:
        return null;
    }
  };

  const getAlertTitle = (alert: Alert) => {
    switch (alert.type) {
      case 'low_stock':
        return '库存低于安全阈值';
      case 'expiring':
        return '物资即将过期';
      case 'expired':
        return '物资已过期';
      default:
        return '预警提示';
    }
  };

  const getAlertDescription = (alert: Alert) => {
    switch (alert.type) {
      case 'low_stock':
        return `当前库存 ${alert.currentQuantity} 件，低于安全阈值 ${alert.threshold} 件，建议及时补货`;
      case 'expiring':
        return `有效期至 ${alert.expiryDate}，即将到期，请尽快使用或处理`;
      case 'expired':
        return `有效期 ${alert.expiryDate} 已过期，请立即处理并补充新物资`;
      default:
        return '';
    }
  };

  const stats = [
    {
      label: '高级预警',
      value: alerts.filter(a => a.severity === 'high').length,
      color: 'text-red-600',
      bg: 'bg-red-100'
    },
    {
      label: '中级预警',
      value: alerts.filter(a => a.severity === 'medium').length,
      color: 'text-orange-600',
      bg: 'bg-orange-100'
    },
    {
      label: '低级预警',
      value: alerts.filter(a => a.severity === 'low').length,
      color: 'text-yellow-600',
      bg: 'bg-yellow-100'
    }
  ];

  return (
    <div className="space-y-6">
      {/* Header */}
      <div>
        <h1 className="text-2xl font-bold text-slate-800">预警中心</h1>
        <p className="text-slate-500 mt-1">库存预警与智能补货建议</p>
      </div>

      {/* Stats */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        {stats.map((stat, idx) => (
          <div key={idx} className="bg-white rounded-xl p-6 shadow-sm border border-slate-200">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-slate-600 text-sm mb-1">{stat.label}</p>
                <p className={`text-3xl font-bold ${stat.color}`}>{stat.value}</p>
              </div>
              <div className={`w-12 h-12 ${stat.bg} rounded-lg flex items-center justify-center`}>
                <Bell className={`w-6 h-6 ${stat.color}`} />
              </div>
            </div>
          </div>
        ))}
      </div>

      {/* Alerts List */}
      <div className="space-y-4">
        {alerts.map((alert) => {
          const Icon = getAlertIcon(alert.type);
          
          return (
            <div
              key={alert.id}
              className={`rounded-xl p-6 shadow-sm border-2 ${getAlertColor(alert.severity)} transition-all hover:shadow-md`}
            >
              <div className="flex items-start gap-4">
                {/* Icon */}
                <div className={`w-12 h-12 rounded-lg flex items-center justify-center flex-shrink-0 ${
                  alert.severity === 'high' ? 'bg-red-100 text-red-600' :
                  alert.severity === 'medium' ? 'bg-orange-100 text-orange-600' :
                  'bg-yellow-100 text-yellow-600'
                }`}>
                  <Icon className="w-6 h-6" />
                </div>

                {/* Content */}
                <div className="flex-1 min-w-0">
                  <div className="flex items-start justify-between mb-2">
                    <div className="flex items-center gap-2">
                      <h3 className="font-semibold text-slate-800">{getAlertTitle(alert)}</h3>
                      {getSeverityBadge(alert.severity)}
                    </div>
                    <span className="text-sm text-slate-500">{alert.createdAt}</span>
                  </div>

                  {/* Material Info */}
                  <div className="mb-3">
                    <div className="flex items-center gap-2 mb-1">
                      <Package className="w-4 h-4 text-slate-400" />
                      <span className="font-medium text-slate-800">{alert.material}</span>
                      <span className="text-slate-500">·</span>
                      <span className="text-slate-600">{alert.category}</span>
                      <span className="text-slate-500">·</span>
                      <span className="text-slate-600">位置: {alert.location}</span>
                    </div>
                  </div>

                  {/* Description */}
                  <p className="text-slate-700 mb-4">{getAlertDescription(alert)}</p>

                  {/* Actions */}
                  <div className="flex gap-3">
                    <button className={`px-4 py-2 rounded-lg text-white text-sm font-medium transition-colors ${
                      alert.severity === 'high' ? 'bg-red-600 hover:bg-red-700' :
                      alert.severity === 'medium' ? 'bg-orange-600 hover:bg-orange-700' :
                      'bg-yellow-600 hover:bg-yellow-700'
                    }`}>
                      {alert.type === 'low_stock' ? '立即补货' : '处理物资'}
                    </button>
                    <button className="px-4 py-2 border border-slate-300 text-slate-700 rounded-lg hover:bg-white transition-colors text-sm font-medium">
                      查看详情
                    </button>
                    <button className="px-4 py-2 border border-slate-300 text-slate-700 rounded-lg hover:bg-white transition-colors text-sm font-medium">
                      忽略提醒
                    </button>
                  </div>
                </div>
              </div>
            </div>
          );
        })}
      </div>

      {/* Recommendations */}
      <div className="bg-white rounded-xl shadow-sm border border-slate-200">
        <div className="p-6 border-b border-slate-200">
          <h3 className="font-semibold text-slate-800">智能补货建议</h3>
          <p className="text-sm text-slate-600 mt-1">基于历史消耗数据分析</p>
        </div>
        <div className="p-6">
          <div className="space-y-4">
            {[
              { name: '应急食品包', category: '食品', suggested: 200, reason: '近期消耗量增加，建议补货至200件' },
              { name: '医用口罩', category: '医疗用品', suggested: 5000, reason: '季节性需求上升，建议提前储备' },
              { name: '帐篷（4人）', category: '应急设施', suggested: 100, reason: '库存偏低，建议补充至300顶' }
            ].map((item, idx) => (
              <div key={idx} className="flex items-center justify-between p-4 bg-slate-50 rounded-lg">
                <div className="flex-1">
                  <div className="flex items-center gap-2 mb-1">
                    <span className="font-medium text-slate-800">{item.name}</span>
                    <span className="text-sm text-slate-500">· {item.category}</span>
                  </div>
                  <p className="text-sm text-slate-600">{item.reason}</p>
                </div>
                <div className="flex items-center gap-3">
                  <div className="text-right">
                    <p className="text-sm text-slate-600">建议补货</p>
                    <p className="font-semibold text-blue-600">{item.suggested} 件</p>
                  </div>
                  <button className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors text-sm font-medium">
                    生成订单
                  </button>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}
