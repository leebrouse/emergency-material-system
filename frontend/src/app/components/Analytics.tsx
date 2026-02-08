import { BarChart, Bar, LineChart, Line, AreaChart, Area, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from 'recharts';
import { TrendingUp, Download, Calendar } from 'lucide-react';

export function Analytics() {
  const monthlyData = [
    { month: '1月', 入库: 4200, 出库: 2800, 调拨: 1200 },
    { month: '2月', 入库: 3800, 出库: 3200, 调拨: 1500 },
    { month: '3月', 入库: 2500, 出库: 4100, 调拨: 1800 },
    { month: '4月', 入库: 3200, 出库: 3600, 调拨: 1400 },
    { month: '5月', 入库: 2800, 出库: 4800, 调拨: 2100 },
    { month: '6月', 入库: 3500, 出库: 3900, 调拨: 1600 }
  ];

  const categoryUsage = [
    { category: '医疗用品', 使用量: 3200, 增长率: 12 },
    { category: '食品', 使用量: 4500, 增长率: -5 },
    { category: '应急设施', 使用量: 1800, 增长率: 8 },
    { category: '设备', 使用量: 950, 增长率: 15 },
    { category: '衣物', 使用量: 2100, 增长率: 3 }
  ];

  const inventoryTrend = [
    { date: '2/3', 总量: 12000 },
    { date: '2/4', 总量: 11500 },
    { date: '2/5', 总量: 12300 },
    { date: '2/6', 总量: 11800 },
    { date: '2/7', 总量: 12100 },
    { date: '2/8', 总量: 12580 }
  ];

  const regionDistribution = [
    { region: '华北', 数量: 3200 },
    { region: '华东', 数量: 2800 },
    { region: '华南', 数量: 2100 },
    { region: '西南', 数量: 1900 },
    { region: '西北', 数量: 1580 },
    { region: '东北', 数量: 1000 }
  ];

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-bold text-slate-800">数据统计分析</h1>
          <p className="text-slate-500 mt-1">物资使用趋势与资源规划</p>
        </div>
        <div className="flex gap-3">
          <button className="px-4 py-2 border border-slate-300 text-slate-700 rounded-lg hover:bg-slate-50 transition-colors flex items-center gap-2">
            <Calendar className="w-4 h-4" />
            选择日期
          </button>
          <button className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2">
            <Download className="w-4 h-4" />
            导出报表
          </button>
        </div>
      </div>

      {/* Summary Cards */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-6">
        {[
          { label: '本月入库总量', value: '3,500', unit: '件', change: '+8.5%', trend: 'up' },
          { label: '本月出库总量', value: '3,900', unit: '件', change: '+12.3%', trend: 'up' },
          { label: '本月调拨次数', value: '42', unit: '次', change: '-3.2%', trend: 'down' },
          { label: '物资周转率', value: '85', unit: '%', change: '+5.1%', trend: 'up' }
        ].map((stat, idx) => (
          <div key={idx} className="bg-white rounded-xl p-6 shadow-sm border border-slate-200">
            <p className="text-slate-600 text-sm mb-2">{stat.label}</p>
            <div className="flex items-baseline gap-2 mb-2">
              <span className="text-3xl font-bold text-slate-800">{stat.value}</span>
              <span className="text-slate-500 text-sm">{stat.unit}</span>
            </div>
            <div className={`flex items-center gap-1 text-sm ${
              stat.trend === 'up' ? 'text-green-600' : 'text-red-600'
            }`}>
              <TrendingUp className={`w-4 h-4 ${stat.trend === 'down' ? 'rotate-180' : ''}`} />
              <span>{stat.change}</span>
            </div>
          </div>
        ))}
      </div>

      {/* Charts Grid */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {/* Monthly Trend */}
        <div className="bg-white rounded-xl p-6 shadow-sm border border-slate-200">
          <h3 className="font-semibold text-slate-800 mb-4">月度物资流动趋势</h3>
          <ResponsiveContainer width="100%" height={300}>
            <BarChart data={monthlyData}>
              <CartesianGrid strokeDasharray="3 3" stroke="#e2e8f0" />
              <XAxis dataKey="month" stroke="#64748b" />
              <YAxis stroke="#64748b" />
              <Tooltip />
              <Legend />
              <Bar dataKey="入库" fill="#3b82f6" />
              <Bar dataKey="出库" fill="#10b981" />
              <Bar dataKey="调拨" fill="#f59e0b" />
            </BarChart>
          </ResponsiveContainer>
        </div>

        {/* Inventory Trend */}
        <div className="bg-white rounded-xl p-6 shadow-sm border border-slate-200">
          <h3 className="font-semibold text-slate-800 mb-4">库存总量变化趋势</h3>
          <ResponsiveContainer width="100%" height={300}>
            <AreaChart data={inventoryTrend}>
              <CartesianGrid strokeDasharray="3 3" stroke="#e2e8f0" />
              <XAxis dataKey="date" stroke="#64748b" />
              <YAxis stroke="#64748b" />
              <Tooltip />
              <Area type="monotone" dataKey="总量" stroke="#3b82f6" fill="#3b82f6" fillOpacity={0.2} />
            </AreaChart>
          </ResponsiveContainer>
        </div>

        {/* Category Usage */}
        <div className="bg-white rounded-xl p-6 shadow-sm border border-slate-200">
          <h3 className="font-semibold text-slate-800 mb-4">分类消耗统计</h3>
          <ResponsiveContainer width="100%" height={300}>
            <BarChart data={categoryUsage} layout="vertical">
              <CartesianGrid strokeDasharray="3 3" stroke="#e2e8f0" />
              <XAxis type="number" stroke="#64748b" />
              <YAxis dataKey="category" type="category" stroke="#64748b" width={80} />
              <Tooltip />
              <Bar dataKey="使用量" fill="#8b5cf6" />
            </BarChart>
          </ResponsiveContainer>
        </div>

        {/* Region Distribution */}
        <div className="bg-white rounded-xl p-6 shadow-sm border border-slate-200">
          <h3 className="font-semibold text-slate-800 mb-4">区域分布统计</h3>
          <ResponsiveContainer width="100%" height={300}>
            <LineChart data={regionDistribution}>
              <CartesianGrid strokeDasharray="3 3" stroke="#e2e8f0" />
              <XAxis dataKey="region" stroke="#64748b" />
              <YAxis stroke="#64748b" />
              <Tooltip />
              <Line type="monotone" dataKey="数量" stroke="#ef4444" strokeWidth={2} />
            </LineChart>
          </ResponsiveContainer>
        </div>
      </div>

      {/* Detailed Table */}
      <div className="bg-white rounded-xl shadow-sm border border-slate-200">
        <div className="p-6 border-b border-slate-200">
          <h3 className="font-semibold text-slate-800">分类详细统计</h3>
        </div>
        <div className="overflow-x-auto">
          <table className="w-full">
            <thead className="bg-slate-50 border-b border-slate-200">
              <tr>
                <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">物资类别</th>
                <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">使用量</th>
                <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">库存量</th>
                <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">周转率</th>
                <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">增长率</th>
                <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">趋势</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-slate-200">
              {categoryUsage.map((item, idx) => (
                <tr key={idx} className="hover:bg-slate-50 transition-colors">
                  <td className="px-6 py-4 font-medium text-slate-800">{item.category}</td>
                  <td className="px-6 py-4 text-slate-700">{item.使用量} 件</td>
                  <td className="px-6 py-4 text-slate-700">
                    {Math.floor(item.使用量 * 0.4)} 件
                  </td>
                  <td className="px-6 py-4 text-slate-700">
                    {Math.floor(70 + Math.random() * 30)}%
                  </td>
                  <td className="px-6 py-4">
                    <span className={`font-semibold ${
                      item.增长率 > 0 ? 'text-green-600' : 'text-red-600'
                    }`}>
                      {item.增长率 > 0 ? '+' : ''}{item.增长率}%
                    </span>
                  </td>
                  <td className="px-6 py-4">
                    <TrendingUp className={`w-5 h-5 ${
                      item.增长率 > 0 ? 'text-green-600' : 'text-red-600 rotate-180'
                    }`} />
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>

      {/* Performance Metrics */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div className="bg-gradient-to-br from-blue-50 to-blue-100 rounded-xl p-6 border border-blue-200">
          <h4 className="font-semibold text-blue-900 mb-4">系统性能指标</h4>
          <div className="space-y-3">
            <div className="flex justify-between">
              <span className="text-blue-700">平均响应时间</span>
              <span className="font-semibold text-blue-900">85ms</span>
            </div>
            <div className="flex justify-between">
              <span className="text-blue-700">并发处理能力</span>
              <span className="font-semibold text-blue-900">1000+/s</span>
            </div>
            <div className="flex justify-between">
              <span className="text-blue-700">系统可用性</span>
              <span className="font-semibold text-blue-900">99.9%</span>
            </div>
          </div>
        </div>

        <div className="bg-gradient-to-br from-green-50 to-green-100 rounded-xl p-6 border border-green-200">
          <h4 className="font-semibold text-green-900 mb-4">业务效率指标</h4>
          <div className="space-y-3">
            <div className="flex justify-between">
              <span className="text-green-700">平均处理时长</span>
              <span className="font-semibold text-green-900">2.3小时</span>
            </div>
            <div className="flex justify-between">
              <span className="text-green-700">任务完成率</span>
              <span className="font-semibold text-green-900">96.5%</span>
            </div>
            <div className="flex justify-between">
              <span className="text-green-700">用户满意度</span>
              <span className="font-semibold text-green-900">4.8/5.0</span>
            </div>
          </div>
        </div>

        <div className="bg-gradient-to-br from-purple-50 to-purple-100 rounded-xl p-6 border border-purple-200">
          <h4 className="font-semibold text-purple-900 mb-4">资源利用率</h4>
          <div className="space-y-3">
            <div className="flex justify-between">
              <span className="text-purple-700">仓储空间利用</span>
              <span className="font-semibold text-purple-900">78%</span>
            </div>
            <div className="flex justify-between">
              <span className="text-purple-700">物资周转效率</span>
              <span className="font-semibold text-purple-900">85%</span>
            </div>
            <div className="flex justify-between">
              <span className="text-purple-700">配送准时率</span>
              <span className="font-semibold text-purple-900">92%</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
