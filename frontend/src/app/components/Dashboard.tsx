import { BarChart, Bar, LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer, PieChart, Pie, Cell } from 'recharts';
import { Package, TrendingUp, AlertTriangle, Truck, Users, Calendar } from 'lucide-react';

const inventoryData = [
  { name: '食品', value: 2580 },
  { name: '医疗', value: 1850 },
  { name: '帐篷', value: 520 },
  { name: '衣物', value: 3200 },
  { name: '工具', value: 980 }
];

const trendData = [
  { month: '1月', 入库: 4000, 出库: 2400 },
  { month: '2月', 入库: 3000, 出库: 1398 },
  { month: '3月', 入库: 2000, 出库: 3800 },
  { month: '4月', 入库: 2780, 出库: 3908 },
  { month: '5月', 入库: 1890, 出库: 4800 },
  { month: '6月', 入库: 2390, 出库: 3800 }
];

const COLORS = ['#3b82f6', '#10b981', '#f59e0b', '#8b5cf6', '#ef4444'];

export function Dashboard() {
  const stats = [
    {
      title: '物资总量',
      value: '12,580',
      unit: '件',
      icon: Package,
      color: 'bg-blue-500',
      change: '+12.5%',
      trend: 'up'
    },
    {
      title: '待处理需求',
      value: '28',
      unit: '项',
      icon: AlertTriangle,
      color: 'bg-orange-500',
      change: '-3.2%',
      trend: 'down'
    },
    {
      title: '运输中',
      value: '15',
      unit: '批次',
      icon: Truck,
      color: 'bg-green-500',
      change: '+8.1%',
      trend: 'up'
    },
    {
      title: '活跃用户',
      value: '186',
      unit: '人',
      icon: Users,
      color: 'bg-purple-500',
      change: '+5.3%',
      trend: 'up'
    }
  ];

  const recentActivities = [
    { id: 1, type: '出库', item: '医疗急救包', quantity: 50, user: '张三', time: '10分钟前' },
    { id: 2, type: '入库', item: '帐篷（4人）', quantity: 30, user: '李四', time: '25分钟前' },
    { id: 3, type: '调拨', item: '饮用水', quantity: 200, user: '王五', time: '1小时前' },
    { id: 4, type: '出库', item: '应急食品包', quantity: 100, user: '赵六', time: '2小时前' },
    { id: 5, type: '入库', item: '发电机', quantity: 5, user: '钱七', time: '3小时前' }
  ];

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-bold text-slate-800">系统概览</h1>
          <p className="text-slate-500 mt-1">实时监控救援物资状态</p>
        </div>
        <div className="flex items-center gap-2 text-slate-600">
          <Calendar className="w-5 h-5" />
          <span>2026年2月8日</span>
        </div>
      </div>

      {/* Stats Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        {stats.map((stat, index) => (
          <div key={index} className="bg-white rounded-xl p-6 shadow-sm border border-slate-200">
            <div className="flex items-start justify-between">
              <div className="flex-1">
                <p className="text-slate-600 text-sm">{stat.title}</p>
                <div className="flex items-baseline gap-2 mt-2">
                  <h3 className="text-3xl font-bold text-slate-800">{stat.value}</h3>
                  <span className="text-slate-500 text-sm">{stat.unit}</span>
                </div>
                <div className={`flex items-center gap-1 mt-2 text-sm ${
                  stat.trend === 'up' ? 'text-green-600' : 'text-red-600'
                }`}>
                  <TrendingUp className={`w-4 h-4 ${stat.trend === 'down' ? 'rotate-180' : ''}`} />
                  <span>{stat.change}</span>
                </div>
              </div>
              <div className={`${stat.color} w-12 h-12 rounded-lg flex items-center justify-center`}>
                <stat.icon className="w-6 h-6 text-white" />
              </div>
            </div>
          </div>
        ))}
      </div>

      {/* Charts Grid */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {/* Inventory Distribution */}
        <div className="bg-white rounded-xl p-6 shadow-sm border border-slate-200">
          <h3 className="font-semibold text-slate-800 mb-4">物资分类分布</h3>
          <ResponsiveContainer width="100%" height={300}>
            <PieChart>
              <Pie
                data={inventoryData}
                cx="50%"
                cy="50%"
                labelLine={false}
                label={({ name, percent }) => `${name} ${(percent * 100).toFixed(0)}%`}
                outerRadius={100}
                fill="#8884d8"
                dataKey="value"
              >
                {inventoryData.map((entry, index) => (
                  <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
                ))}
              </Pie>
              <Tooltip />
            </PieChart>
          </ResponsiveContainer>
        </div>

        {/* Trend Chart */}
        <div className="bg-white rounded-xl p-6 shadow-sm border border-slate-200">
          <h3 className="font-semibold text-slate-800 mb-4">出入库趋势</h3>
          <ResponsiveContainer width="100%" height={300}>
            <LineChart data={trendData}>
              <CartesianGrid strokeDasharray="3 3" stroke="#e2e8f0" />
              <XAxis dataKey="month" stroke="#64748b" />
              <YAxis stroke="#64748b" />
              <Tooltip />
              <Line type="monotone" dataKey="入库" stroke="#3b82f6" strokeWidth={2} />
              <Line type="monotone" dataKey="出库" stroke="#10b981" strokeWidth={2} />
            </LineChart>
          </ResponsiveContainer>
        </div>
      </div>

      {/* Recent Activities */}
      <div className="bg-white rounded-xl shadow-sm border border-slate-200">
        <div className="p-6 border-b border-slate-200">
          <h3 className="font-semibold text-slate-800">最近活动</h3>
        </div>
        <div className="divide-y divide-slate-200">
          {recentActivities.map((activity) => (
            <div key={activity.id} className="p-6 flex items-center justify-between hover:bg-slate-50 transition-colors">
              <div className="flex items-center gap-4">
                <div className={`w-10 h-10 rounded-lg flex items-center justify-center ${
                  activity.type === '入库' ? 'bg-green-100 text-green-600' :
                  activity.type === '出库' ? 'bg-blue-100 text-blue-600' :
                  'bg-orange-100 text-orange-600'
                }`}>
                  <Package className="w-5 h-5" />
                </div>
                <div>
                  <div className="flex items-center gap-2">
                    <span className={`px-2 py-0.5 rounded text-xs font-medium ${
                      activity.type === '入库' ? 'bg-green-100 text-green-700' :
                      activity.type === '出库' ? 'bg-blue-100 text-blue-700' :
                      'bg-orange-100 text-orange-700'
                    }`}>
                      {activity.type}
                    </span>
                    <span className="font-medium text-slate-800">{activity.item}</span>
                  </div>
                  <p className="text-sm text-slate-500 mt-1">
                    数量: {activity.quantity} · 操作人: {activity.user}
                  </p>
                </div>
              </div>
              <span className="text-sm text-slate-400">{activity.time}</span>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
