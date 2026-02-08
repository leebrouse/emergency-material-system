import { Shield, Check, X } from 'lucide-react';

interface PermissionTableProps {
  userRole: string;
}

export function RolePermissions({ userRole }: PermissionTableProps) {
  const permissions = [
    { module: '系统概览', admin: true, warehouse: true, rescue: false },
    { module: '物资管理', admin: true, warehouse: true, rescue: false },
    { module: '需求管理', admin: true, warehouse: true, rescue: true },
    { module: '物流追踪', admin: true, warehouse: true, rescue: true },
    { module: '预警中心', admin: true, warehouse: true, rescue: false },
    { module: '数据分析', admin: true, warehouse: true, rescue: false },
    { module: '系统设置', admin: true, warehouse: false, rescue: false }
  ];

  const roles = [
    { key: 'admin', label: '系统管理员', color: 'purple' },
    { key: 'warehouse', label: '仓库管理员', color: 'blue' },
    { key: 'rescue', label: '救援人员', color: 'green' }
  ];

  return (
    <div className="bg-white rounded-xl shadow-sm border border-slate-200 overflow-hidden">
      <div className="p-6 border-b border-slate-200 bg-gradient-to-r from-blue-50 to-white">
        <div className="flex items-center gap-3">
          <Shield className="w-6 h-6 text-blue-600" />
          <div>
            <h3 className="font-semibold text-slate-800">角色权限说明</h3>
            <p className="text-sm text-slate-600 mt-1">不同角色可访问的功能模块</p>
          </div>
        </div>
      </div>

      <div className="overflow-x-auto">
        <table className="w-full">
          <thead className="bg-slate-50 border-b border-slate-200">
            <tr>
              <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">功能模块</th>
              {roles.map((role) => (
                <th key={role.key} className="px-6 py-4 text-center text-sm font-semibold text-slate-700">
                  <div className="flex flex-col items-center gap-1">
                    <span>{role.label}</span>
                    {userRole === role.key && (
                      <span className="px-2 py-0.5 bg-blue-100 text-blue-700 rounded text-xs">
                        当前角色
                      </span>
                    )}
                  </div>
                </th>
              ))}
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-200">
            {permissions.map((perm, idx) => (
              <tr key={idx} className="hover:bg-slate-50 transition-colors">
                <td className="px-6 py-4 font-medium text-slate-800">{perm.module}</td>
                <td className="px-6 py-4 text-center">
                  {perm.admin ? (
                    <div className="inline-flex items-center justify-center w-6 h-6 bg-green-100 rounded-full">
                      <Check className="w-4 h-4 text-green-600" />
                    </div>
                  ) : (
                    <div className="inline-flex items-center justify-center w-6 h-6 bg-red-100 rounded-full">
                      <X className="w-4 h-4 text-red-600" />
                    </div>
                  )}
                </td>
                <td className="px-6 py-4 text-center">
                  {perm.warehouse ? (
                    <div className="inline-flex items-center justify-center w-6 h-6 bg-green-100 rounded-full">
                      <Check className="w-4 h-4 text-green-600" />
                    </div>
                  ) : (
                    <div className="inline-flex items-center justify-center w-6 h-6 bg-red-100 rounded-full">
                      <X className="w-4 h-4 text-red-600" />
                    </div>
                  )}
                </td>
                <td className="px-6 py-4 text-center">
                  {perm.rescue ? (
                    <div className="inline-flex items-center justify-center w-6 h-6 bg-green-100 rounded-full">
                      <Check className="w-4 h-4 text-green-600" />
                    </div>
                  ) : (
                    <div className="inline-flex items-center justify-center w-6 h-6 bg-red-100 rounded-full">
                      <X className="w-4 h-4 text-red-600" />
                    </div>
                  )}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      <div className="p-6 bg-slate-50 border-t border-slate-200">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div className="flex items-start gap-3">
            <div className="w-10 h-10 bg-purple-100 rounded-lg flex items-center justify-center flex-shrink-0">
              <Shield className="w-5 h-5 text-purple-600" />
            </div>
            <div>
              <p className="font-medium text-slate-800 mb-1">系统管理员</p>
              <p className="text-sm text-slate-600">拥有所有功能的完整访问权限，可管理用户和系统配置</p>
            </div>
          </div>

          <div className="flex items-start gap-3">
            <div className="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center flex-shrink-0">
              <Shield className="w-5 h-5 text-blue-600" />
            </div>
            <div>
              <p className="font-medium text-slate-800 mb-1">仓库管理员</p>
              <p className="text-sm text-slate-600">可管理库存、处理需求、查看预警和分析数据</p>
            </div>
          </div>

          <div className="flex items-start gap-3">
            <div className="w-10 h-10 bg-green-100 rounded-lg flex items-center justify-center flex-shrink-0">
              <Shield className="w-5 h-5 text-green-600" />
            </div>
            <div>
              <p className="font-medium text-slate-800 mb-1">救援人员</p>
              <p className="text-sm text-slate-600">可申报物资需求和追踪物流运输状态</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
