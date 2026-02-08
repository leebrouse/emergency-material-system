import { useState } from 'react';
import { Truck, MapPin, Package, Clock, CheckCircle, Navigation } from 'lucide-react';

interface LogisticsItem {
  id: string;
  trackingNumber: string;
  destination: string;
  items: string[];
  status: 'preparing' | 'transit' | 'arrived' | 'delivered';
  currentLocation: string;
  progress: number;
  estimatedTime: string;
  driver: string;
  vehicle: string;
  timeline: {
    time: string;
    location: string;
    status: string;
    description: string;
  }[];
}

export function Logistics() {
  const [selectedTracking, setSelectedTracking] = useState<string | null>(null);

  const logistics: LogisticsItem[] = [
    {
      id: 'L001',
      trackingNumber: 'TN20260208001',
      destination: '四川省雅安市应急指挥中心',
      items: ['帐篷 x50', '应急食品 x200', '饮用水 x100'],
      status: 'transit',
      currentLocation: '成都市武侯区中转站',
      progress: 65,
      estimatedTime: '2026-02-08 18:00',
      driver: '张师傅',
      vehicle: '川A·12345',
      timeline: [
        {
          time: '2026-02-08 09:00',
          location: '成都市物资仓库',
          status: 'completed',
          description: '物资装车完成，准备发车'
        },
        {
          time: '2026-02-08 11:30',
          location: '成都市武侯区中转站',
          status: 'current',
          description: '车辆已到达中转站，正在进行检查'
        },
        {
          time: '2026-02-08 14:00',
          location: '雅安市名山区',
          status: 'pending',
          description: '预计通过名山区检查站'
        },
        {
          time: '2026-02-08 18:00',
          location: '雅安市应急指挥中心',
          status: 'pending',
          description: '预计送达目的地'
        }
      ]
    },
    {
      id: 'L002',
      trackingNumber: 'TN20260208002',
      destination: '湖南省长沙市防汛指挥部',
      items: ['冲锋舟 x5', '救生衣 x100', '应急灯 x50'],
      status: 'preparing',
      currentLocation: '武汉市物资仓库',
      progress: 15,
      estimatedTime: '2026-02-09 10:00',
      driver: '李师傅',
      vehicle: '鄂A·67890',
      timeline: [
        {
          time: '2026-02-08 13:00',
          location: '武汉市物资仓库',
          status: 'current',
          description: '正在进行物资清点和装车'
        },
        {
          time: '2026-02-08 16:00',
          location: '武汉市物资仓库',
          status: 'pending',
          description: '预计完成装车并发车'
        },
        {
          time: '2026-02-09 10:00',
          location: '长沙市防汛指挥部',
          status: 'pending',
          description: '预计送达目的地'
        }
      ]
    },
    {
      id: 'L003',
      trackingNumber: 'TN20260207001',
      destination: '北京市朝阳区社区服务中心',
      items: ['医用口罩 x5000', '消毒液 x30', '体温计 x20'],
      status: 'delivered',
      currentLocation: '北京市朝阳区社区服务中心',
      progress: 100,
      estimatedTime: '2026-02-07 16:00',
      driver: '王师傅',
      vehicle: '京A·11111',
      timeline: [
        {
          time: '2026-02-07 09:00',
          location: '北京市物资仓库',
          status: 'completed',
          description: '物资装车完成，准备发车'
        },
        {
          time: '2026-02-07 12:00',
          location: '北京市朝阳区',
          status: 'completed',
          description: '车辆进入朝阳区'
        },
        {
          time: '2026-02-07 14:30',
          location: '北京市朝阳区社区服务中心',
          status: 'completed',
          description: '物资送达并签收'
        }
      ]
    }
  ];

  const getStatusInfo = (status: string) => {
    switch (status) {
      case 'preparing':
        return { label: '准备中', color: 'bg-orange-100 text-orange-700', icon: Package };
      case 'transit':
        return { label: '运输中', color: 'bg-blue-100 text-blue-700', icon: Truck };
      case 'arrived':
        return { label: '已到达', color: 'bg-green-100 text-green-700', icon: MapPin };
      case 'delivered':
        return { label: '已送达', color: 'bg-slate-100 text-slate-700', icon: CheckCircle };
      default:
        return { label: status, color: 'bg-slate-100 text-slate-700', icon: Package };
    }
  };

  return (
    <div className="space-y-6">
      {/* Header */}
      <div>
        <h1 className="text-2xl font-bold text-slate-800">物流追踪</h1>
        <p className="text-slate-500 mt-1">实时监控物资运输状态</p>
      </div>

      {/* Map Placeholder */}
      <div className="bg-white rounded-xl shadow-sm border border-slate-200 overflow-hidden">
        <div className="h-96 bg-gradient-to-br from-blue-50 to-slate-100 flex items-center justify-center relative">
          <div className="absolute inset-0 opacity-10">
            <svg className="w-full h-full" viewBox="0 0 800 400">
              <path d="M100,200 Q250,100 400,200 T700,200" stroke="#3b82f6" strokeWidth="3" fill="none" strokeDasharray="5,5" />
            </svg>
          </div>
          <div className="text-center z-10">
            <MapPin className="w-16 h-16 text-blue-600 mx-auto mb-4" />
            <h3 className="text-xl font-semibold text-slate-800 mb-2">地图可视化</h3>
            <p className="text-slate-600">集成高德/百度地图 API 展示运输路径</p>
            <p className="text-sm text-slate-500 mt-2">实际部署时接入真实地图服务</p>
          </div>
        </div>
      </div>

      {/* Logistics List */}
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        {/* List */}
        <div className="lg:col-span-1 space-y-4">
          <h3 className="font-semibold text-slate-800">运输任务</h3>
          {logistics.map((item) => {
            const statusInfo = getStatusInfo(item.status);
            const StatusIcon = statusInfo.icon;
            
            return (
              <button
                key={item.id}
                onClick={() => setSelectedTracking(item.trackingNumber)}
                className={`w-full text-left bg-white rounded-xl p-4 shadow-sm border transition-all ${
                  selectedTracking === item.trackingNumber
                    ? 'border-blue-500 ring-2 ring-blue-100'
                    : 'border-slate-200 hover:border-blue-300'
                }`}
              >
                <div className="flex items-start gap-3">
                  <div className={`w-10 h-10 rounded-lg flex items-center justify-center ${statusInfo.color}`}>
                    <StatusIcon className="w-5 h-5" />
                  </div>
                  <div className="flex-1 min-w-0">
                    <div className="flex items-center gap-2 mb-1">
                      <span className="font-medium text-slate-800 truncate">{item.trackingNumber}</span>
                      <span className={`px-2 py-0.5 rounded text-xs font-medium ${statusInfo.color}`}>
                        {statusInfo.label}
                      </span>
                    </div>
                    <p className="text-sm text-slate-600 truncate">{item.destination}</p>
                    <div className="flex items-center gap-1 text-xs text-slate-500 mt-2">
                      <Navigation className="w-3 h-3" />
                      {item.currentLocation}
                    </div>
                  </div>
                </div>
                
                {/* Progress Bar */}
                <div className="mt-3">
                  <div className="flex justify-between text-xs text-slate-600 mb-1">
                    <span>运输进度</span>
                    <span>{item.progress}%</span>
                  </div>
                  <div className="h-2 bg-slate-100 rounded-full overflow-hidden">
                    <div 
                      className="h-full bg-blue-600 rounded-full transition-all duration-500"
                      style={{ width: `${item.progress}%` }}
                    />
                  </div>
                </div>
              </button>
            );
          })}
        </div>

        {/* Details */}
        <div className="lg:col-span-2">
          {selectedTracking ? (
            (() => {
              const item = logistics.find(l => l.trackingNumber === selectedTracking);
              if (!item) return null;
              
              const statusInfo = getStatusInfo(item.status);
              
              return (
                <div className="bg-white rounded-xl shadow-sm border border-slate-200 overflow-hidden">
                  {/* Header */}
                  <div className="p-6 border-b border-slate-200 bg-gradient-to-r from-blue-50 to-white">
                    <div className="flex items-start justify-between mb-4">
                      <div>
                        <h3 className="text-xl font-semibold text-slate-800 mb-1">
                          {item.trackingNumber}
                        </h3>
                        <p className="text-slate-600">{item.destination}</p>
                      </div>
                      <span className={`px-3 py-1 rounded-full text-sm font-medium ${statusInfo.color}`}>
                        {statusInfo.label}
                      </span>
                    </div>
                    
                    {/* Info Grid */}
                    <div className="grid grid-cols-2 gap-4">
                      <div>
                        <p className="text-sm text-slate-600 mb-1">当前位置</p>
                        <p className="font-medium text-slate-800">{item.currentLocation}</p>
                      </div>
                      <div>
                        <p className="text-sm text-slate-600 mb-1">预计送达</p>
                        <p className="font-medium text-slate-800">{item.estimatedTime}</p>
                      </div>
                      <div>
                        <p className="text-sm text-slate-600 mb-1">驾驶员</p>
                        <p className="font-medium text-slate-800">{item.driver}</p>
                      </div>
                      <div>
                        <p className="text-sm text-slate-600 mb-1">车辆号牌</p>
                        <p className="font-medium text-slate-800">{item.vehicle}</p>
                      </div>
                    </div>
                  </div>

                  {/* Items */}
                  <div className="p-6 border-b border-slate-200">
                    <h4 className="font-semibold text-slate-800 mb-3">运输物资</h4>
                    <div className="flex flex-wrap gap-2">
                      {item.items.map((materialItem, idx) => (
                        <span key={idx} className="px-3 py-1 bg-blue-50 text-blue-700 rounded-lg text-sm">
                          {materialItem}
                        </span>
                      ))}
                    </div>
                  </div>

                  {/* Timeline */}
                  <div className="p-6">
                    <h4 className="font-semibold text-slate-800 mb-4">物流轨迹</h4>
                    <div className="space-y-4">
                      {item.timeline.map((event, idx) => (
                        <div key={idx} className="flex gap-4">
                          {/* Timeline Line */}
                          <div className="flex flex-col items-center">
                            <div className={`w-3 h-3 rounded-full ${
                              event.status === 'completed' ? 'bg-green-500' :
                              event.status === 'current' ? 'bg-blue-500 ring-4 ring-blue-100' :
                              'bg-slate-300'
                            }`} />
                            {idx < item.timeline.length - 1 && (
                              <div className={`w-0.5 h-12 ${
                                event.status === 'completed' ? 'bg-green-200' : 'bg-slate-200'
                              }`} />
                            )}
                          </div>

                          {/* Content */}
                          <div className="flex-1 pb-4">
                            <div className="flex items-center gap-2 mb-1">
                              <Clock className="w-4 h-4 text-slate-400" />
                              <span className="text-sm font-medium text-slate-800">{event.time}</span>
                            </div>
                            <p className="font-medium text-slate-800 mb-1">{event.location}</p>
                            <p className="text-sm text-slate-600">{event.description}</p>
                          </div>
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
              );
            })()
          ) : (
            <div className="bg-white rounded-xl shadow-sm border border-slate-200 h-full flex items-center justify-center">
              <div className="text-center">
                <Truck className="w-16 h-16 text-slate-300 mx-auto mb-4" />
                <p className="text-slate-500">请从左侧选择运输任务查看详情</p>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
