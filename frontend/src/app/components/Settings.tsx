import { Users, Shield, Bell, Database, Key } from 'lucide-react';
import { RolePermissions } from './RolePermissions';

interface SettingsProps {
  userRole?: string;
}

export function Settings({ userRole = 'admin' }: SettingsProps) {
  const users = [
    { id: 1, name: '张伟', role: '系统管理员', email: 'zhangwei@example.com', status: 'active' },
    { id: 2, name: '李明', role: '仓库管理员', email: 'liming@example.com', status: 'active' },
    { id: 3, name: '王芳', role: '救援人员', email: 'wangfang@example.com', status: 'active' },
    { id: 4, name: '赵强', role: '救援人员', email: 'zhaoqiang@example.com', status: 'inactive' }
  ];

  return (
    <div className="space-y-6">
      {/* Header */}
      <div>
        <h1 className="text-2xl font-bold text-slate-800">系统设置</h1>
        <p className="text-slate-500 mt-1">权限管理与系统配置</p>
      </div>

      {/* Settings Tabs */}
      <div className="bg-white rounded-xl shadow-sm border border-slate-200">
        <div className="border-b border-slate-200">
          <div className="flex gap-6 px-6">
            {[
              { label: '用户管理', icon: Users, active: true },
              { label: '权限配置', icon: Shield, active: false },
              { label: '通知设置', icon: Bell, active: false },
              { label: '数据备份', icon: Database, active: false }
            ].map((tab, idx) => (
              <button
                key={idx}
                className={`flex items-center gap-2 px-4 py-4 border-b-2 transition-colors ${
                  tab.active
                    ? 'border-blue-600 text-blue-600'
                    : 'border-transparent text-slate-600 hover:text-slate-800'
                }`}
              >
                <tab.icon className="w-5 h-5" />
                <span className="font-medium">{tab.label}</span>
              </button>
            ))}
          </div>
        </div>

        {/* User Management Content */}
        <div className="p-6">
          <div className="flex items-center justify-between mb-6">
            <div>
              <h3 className="font-semibold text-slate-800">用户列表</h3>
              <p className="text-sm text-slate-600 mt-1">管理系统用户和角色权限</p>
            </div>
            <button className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2">
              <Users className="w-4 h-4" />
              添加用户
            </button>
          </div>

          {/* Users Table */}
          <div className="overflow-x-auto">
            <table className="w-full">
              <thead className="bg-slate-50 border-b border-slate-200">
                <tr>
                  <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">用户名</th>
                  <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">角色</th>
                  <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">邮箱</th>
                  <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">状态</th>
                  <th className="px-6 py-4 text-left text-sm font-semibold text-slate-700">操作</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-slate-200">
                {users.map((user) => (
                  <tr key={user.id} className="hover:bg-slate-50 transition-colors">
                    <td className="px-6 py-4">
                      <div className="flex items-center gap-3">
                        <div className="w-10 h-10 bg-blue-100 rounded-full flex items-center justify-center">
                          <span className="font-semibold text-blue-600">{user.name[0]}</span>
                        </div>
                        <span className="font-medium text-slate-800">{user.name}</span>
                      </div>
                    </td>
                    <td className="px-6 py-4">
                      <span className={`px-3 py-1 rounded-full text-xs font-medium ${
                        user.role === '系统管理员' ? 'bg-purple-100 text-purple-700' :
                        user.role === '仓库管理员' ? 'bg-blue-100 text-blue-700' :
                        'bg-green-100 text-green-700'
                      }`}>
                        {user.role}
                      </span>
                    </td>
                    <td className="px-6 py-4 text-slate-600">{user.email}</td>
                    <td className="px-6 py-4">
                      <span className={`px-3 py-1 rounded-full text-xs font-medium ${
                        user.status === 'active' 
                          ? 'bg-green-100 text-green-700' 
                          : 'bg-slate-100 text-slate-700'
                      }`}>
                        {user.status === 'active' ? '活跃' : '停用'}
                      </span>
                    </td>
                    <td className="px-6 py-4">
                      <div className="flex gap-2">
                        <button className="text-blue-600 hover:text-blue-700 text-sm font-medium">
                          编辑
                        </button>
                        <button className="text-red-600 hover:text-red-700 text-sm font-medium">
                          删除
                        </button>
                      </div>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>

      {/* Role Permissions */}
      <div className="bg-white rounded-xl shadow-sm border border-slate-200">
        <div className="p-6 border-b border-slate-200">
          <div className="flex items-center gap-3">
            <Shield className="w-5 h-5 text-blue-600" />
            <h3 className="font-semibold text-slate-800">角色权限配置</h3>
          </div>
        </div>
        <div className="p-6">
          <RolePermissions userRole={userRole} />
        </div>
      </div>

      {/* Security Settings */}
      <div className="bg-white rounded-xl shadow-sm border border-slate-200">
        <div className="p-6 border-b border-slate-200">
          <div className="flex items-center gap-3">
            <Key className="w-5 h-5 text-blue-600" />
            <h3 className="font-semibold text-slate-800">安全认证设置</h3>
          </div>
        </div>
        <div className="p-6">
          <div className="space-y-4">
            <div className="flex items-center justify-between p-4 bg-slate-50 rounded-lg">
              <div>
                <p className="font-medium text-slate-800">JWT Token 过期时间</p>
                <p className="text-sm text-slate-600">设置用户登录状态有效期</p>
              </div>
              <select className="px-4 py-2 border border-slate-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
                <option>24 小时</option>
                <option>7 天</option>
                <option>30 天</option>
              </select>
            </div>

            <div className="flex items-center justify-between p-4 bg-slate-50 rounded-lg">
              <div>
                <p className="font-medium text-slate-800">密码强度要求</p>
                <p className="text-sm text-slate-600">设置用户密码复杂度规则</p>
              </div>
              <label className="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" className="sr-only peer" defaultChecked />
                <div className="w-11 h-6 bg-slate-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-slate-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>

            <div className="flex items-center justify-between p-4 bg-slate-50 rounded-lg">
              <div>
                <p className="font-medium text-slate-800">数据加密传输</p>
                <p className="text-sm text-slate-600">启用 HTTPS 加密通信</p>
              </div>
              <label className="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" className="sr-only peer" defaultChecked />
                <div className="w-11 h-6 bg-slate-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-slate-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>

            <div className="flex items-center justify-between p-4 bg-slate-50 rounded-lg">
              <div>
                <p className="font-medium text-slate-800">登录失败锁定</p>
                <p className="text-sm text-slate-600">连续失败 5 次后锁定账户 30 分钟</p>
              </div>
              <label className="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" className="sr-only peer" defaultChecked />
                <div className="w-11 h-6 bg-slate-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-slate-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}