import { useState } from 'react';
import {
  LayoutDashboard,
  Package,
  FileText,
  Truck,
  Bell,
  BarChart3,
  Settings as SettingsIcon,
  LogOut,
  Menu,
  X
} from 'lucide-react';
import { Login } from './components/Login';
import { Dashboard } from './components/Dashboard';
import { Inventory } from './components/Inventory';
import { Demands } from './components/Demands';
import { Logistics } from './components/Logistics';
import { Alerts } from './components/Alerts';
import { Analytics } from './components/Analytics';
import { Settings } from './components/Settings';

type Page = 'dashboard' | 'inventory' | 'demands' | 'logistics' | 'alerts' | 'analytics' | 'settings';

// 定义角色权限
type Role = 'admin' | 'warehouse' | 'rescue';

interface MenuItem {
  id: Page;
  label: string;
  icon: any;
  roles: Role[]; // 允许访问的角色
}

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [currentPage, setCurrentPage] = useState<Page>('dashboard');
  const [userRole, setUserRole] = useState<Role>('rescue');
  const [username, setUsername] = useState('');
  const [sidebarOpen, setSidebarOpen] = useState(true);

  const handleLogin = (role: string, name: string) => {
    setUserRole(role as Role);
    setUsername(name);
    setIsLoggedIn(true);
    // 根据角色设置默认页面
    if (role === 'rescue') {
      setCurrentPage('demands');
    } else {
      setCurrentPage('dashboard');
    }
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
    setUserRole('rescue');
    setUsername('');
    setCurrentPage('dashboard');
  };

  // 定义所有菜单项及其访问权限
  const allMenuItems: MenuItem[] = [
    { 
      id: 'dashboard', 
      label: '系统概览', 
      icon: LayoutDashboard,
      roles: ['admin', 'warehouse'] // 仅管理员和仓库管理员可见
    },
    { 
      id: 'inventory', 
      label: '物资管理', 
      icon: Package,
      roles: ['admin', 'warehouse'] // 仅管理员和仓库管理员可见
    },
    { 
      id: 'demands', 
      label: '需求管理', 
      icon: FileText,
      roles: ['admin', 'warehouse', 'rescue'] // 所有角色可见
    },
    { 
      id: 'logistics', 
      label: '物流追踪', 
      icon: Truck,
      roles: ['admin', 'warehouse', 'rescue'] // 所有角色可见
    },
    { 
      id: 'alerts', 
      label: '预警中心', 
      icon: Bell,
      roles: ['admin', 'warehouse'] // 仅管理员和仓库管理员可见
    },
    { 
      id: 'analytics', 
      label: '数据分析', 
      icon: BarChart3,
      roles: ['admin', 'warehouse'] // 仅管理员和仓库管理员可见
    },
    { 
      id: 'settings', 
      label: '系统设置', 
      icon: SettingsIcon,
      roles: ['admin'] // 仅管理员可见
    }
  ];

  // 根据用户角色过滤菜单
  const menuItems = allMenuItems.filter(item => item.roles.includes(userRole));

  const getRoleName = (role: string) => {
    switch (role) {
      case 'admin':
        return '系统管理员';
      case 'warehouse':
        return '仓库管理员';
      case 'rescue':
        return '救援人员';
      default:
        return role;
    }
  };

  // 检查用户是否有权限访问当前页面
  const hasPageAccess = (page: Page): boolean => {
    const menuItem = allMenuItems.find(item => item.id === page);
    return menuItem ? menuItem.roles.includes(userRole) : false;
  };

  const renderPage = () => {
    // 权限检查
    if (!hasPageAccess(currentPage)) {
      return (
        <div className="flex items-center justify-center h-full">
          <div className="text-center">
            <div className="w-16 h-16 bg-red-100 rounded-full flex items-center justify-center mx-auto mb-4">
              <Bell className="w-8 h-8 text-red-600" />
            </div>
            <h2 className="text-2xl font-bold text-slate-800 mb-2">访问受限</h2>
            <p className="text-slate-600 mb-4">您没有权限访问此页面</p>
            <button 
              onClick={() => setCurrentPage(menuItems[0]?.id || 'demands')}
              className="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
            >
              返回首页
            </button>
          </div>
        </div>
      );
    }

    switch (currentPage) {
      case 'dashboard':
        return <Dashboard />;
      case 'inventory':
        return <Inventory />;
      case 'demands':
        return <Demands />;
      case 'logistics':
        return <Logistics />;
      case 'alerts':
        return <Alerts />;
      case 'analytics':
        return <Analytics />;
      case 'settings':
        return <Settings userRole={userRole} />;
      default:
        return <Dashboard />;
    }
  };

  if (!isLoggedIn) {
    return <Login onLogin={handleLogin} />;
  }

  return (
    <div className="min-h-screen bg-slate-50 flex">
      {/* Sidebar */}
      <aside
        className={`fixed lg:static inset-y-0 left-0 z-50 bg-white border-r border-slate-200 transition-all duration-300 ${
          sidebarOpen ? 'w-64' : 'w-0 lg:w-20'
        }`}
      >
        {/* Logo */}
        <div className="h-16 flex items-center justify-between px-6 border-b border-slate-200">
          {sidebarOpen && (
            <div className="flex items-center gap-3">
              <div className="w-8 h-8 bg-blue-600 rounded-lg flex items-center justify-center">
                <Package className="w-5 h-5 text-white" />
              </div>
              <span className="font-bold text-slate-800">救援物资管理</span>
            </div>
          )}
          <button
            onClick={() => setSidebarOpen(!sidebarOpen)}
            className="lg:hidden p-2 hover:bg-slate-100 rounded-lg transition-colors"
          >
            <X className="w-5 h-5 text-slate-600" />
          </button>
        </div>

        {/* Navigation */}
        <nav className="p-4 space-y-2">
          {menuItems.map((item) => {
            const Icon = item.icon;
            const isActive = currentPage === item.id;

            return (
              <button
                key={item.id}
                onClick={() => {
                  setCurrentPage(item.id as Page);
                  if (window.innerWidth < 1024) {
                    setSidebarOpen(false);
                  }
                }}
                className={`w-full flex items-center gap-3 px-4 py-3 rounded-lg transition-all ${
                  isActive
                    ? 'bg-blue-600 text-white shadow-lg shadow-blue-600/20'
                    : 'text-slate-700 hover:bg-slate-100'
                } ${!sidebarOpen && 'lg:justify-center'}`}
              >
                <Icon className="w-5 h-5 flex-shrink-0" />
                {sidebarOpen && <span className="font-medium">{item.label}</span>}
              </button>
            );
          })}
        </nav>

        {/* User Info */}
        {sidebarOpen && (
          <div className="absolute bottom-0 left-0 right-0 p-4 border-t border-slate-200 bg-white">
            <div className="flex items-center gap-3 mb-3">
              <div className="w-10 h-10 bg-blue-100 rounded-full flex items-center justify-center">
                <span className="font-semibold text-blue-600">
                  {getRoleName(userRole)[0]}
                </span>
              </div>
              <div className="flex-1 min-w-0">
                <p className="font-medium text-slate-800 truncate">当前用户</p>
                <p className="text-sm text-slate-500 truncate">{getRoleName(userRole)}</p>
              </div>
            </div>
            <button
              onClick={handleLogout}
              className="w-full flex items-center justify-center gap-2 px-4 py-2 border border-slate-300 text-slate-700 rounded-lg hover:bg-slate-50 transition-colors"
            >
              <LogOut className="w-4 h-4" />
              <span className="font-medium">退出登录</span>
            </button>
          </div>
        )}
      </aside>

      {/* Main Content */}
      <div className="flex-1 flex flex-col min-w-0">
        {/* Header */}
        <header className="h-16 bg-white border-b border-slate-200 flex items-center justify-between px-6">
          <button
            onClick={() => setSidebarOpen(!sidebarOpen)}
            className="p-2 hover:bg-slate-100 rounded-lg transition-colors lg:hidden"
          >
            <Menu className="w-6 h-6 text-slate-600" />
          </button>

          <div className="flex items-center gap-4">
            {/* Notifications */}
            <button className="relative p-2 hover:bg-slate-100 rounded-lg transition-colors">
              <Bell className="w-5 h-5 text-slate-600" />
              <span className="absolute top-1 right-1 w-2 h-2 bg-red-500 rounded-full"></span>
            </button>

            {/* User Menu */}
            <div className="hidden lg:flex items-center gap-3 pl-4 border-l border-slate-200">
              <div className="text-right">
                <p className="text-sm font-medium text-slate-800">当前用户</p>
                <p className="text-xs text-slate-500">{getRoleName(userRole)}</p>
              </div>
              <div className="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                <span className="text-sm font-semibold text-blue-600">
                  {getRoleName(userRole)[0]}
                </span>
              </div>
            </div>
          </div>
        </header>

        {/* Page Content */}
        <main className="flex-1 overflow-auto p-6">
          {renderPage()}
        </main>

        {/* Footer */}
        <footer className="bg-white border-t border-slate-200 px-6 py-4">
          <div className="flex flex-col md:flex-row items-center justify-between gap-4">
            <p className="text-sm text-slate-600">
              © 2026 救援物资管理系统 · 基于 Golang 构建
            </p>
            <div className="flex items-center gap-6 text-sm text-slate-600">
              <span>JWT 安全认证</span>
              <span>·</span>
              <span>接口级权限控制</span>
              <span>·</span>
              <span>数据加密存储</span>
            </div>
          </div>
        </footer>
      </div>

      {/* Mobile Overlay */}
      {sidebarOpen && (
        <div
          className="fixed inset-0 bg-black/50 z-40 lg:hidden"
          onClick={() => setSidebarOpen(false)}
        />
      )}
    </div>
  );
}

export default App;