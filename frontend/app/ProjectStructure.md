# Emergency Material Dispatch System - Frontend Architecture

## Tech Stack
- **Framework**: Vue 3 (Composition API)
- **Build Tool**: Vite
- **UI Library**: Element Plus
- **Styling**: Tailwind CSS + SCSS
- **Charts**: ECharts
- **State Management**: Pinia
- **Routing**: Vue Router

## Directory Structure
```
frontend-app/
├── public/                 # Static assets
├── src/
│   ├── api/                # Axios client and API modules
│   ├── assets/             # Images, fonts, global styles
│   ├── components/         # Reusable UI components
│   ├── layouts/            # Page layouts (MainLayout.vue)
│   ├── router/             # Route definitions and guards
│   ├── stores/             # Pinia stores (user, inventory, etc.)
│   ├── views/              # Page components
│   │   ├── Dashboard.vue   # Anayltics dashboard
│   │   ├── Inventory.vue   # Stock management
│   │   ├── Dispatch.vue    # Dispatch workflow
│   │   ├── Logistics.vue   # Map tracking
│   │   └── Login.vue       # Auth page
│   ├── App.vue             # Root component
│   └── main.ts             # App entry point
├── index.html
├── package.json
├── tailwind.config.js
├── tsconfig.json
└── vite.config.ts
```

## Key Features Implemented
1. **RBAC & Auth**: 
   - Login page with glassmorphism design.
   - Route guards preventing unauthorized access.
   - Dynamic sidebar menu based on user role (Admin, Warehouse, Rescue).

2. **Responsive Layout**:
   - `MainLayout.vue` uses CSS Grid `grid-template-columns: 250px 1fr`.
   - Sidebar collapses on mobile (implementation pending, currently fixed width).
   - Top navigation with alert integration.

3. **Dashboard & Visualization**:
   - ECharts integration for real-time data visualization.
   - Logistics map placeholder ready for AMap integration.

4. **Inventory & Dispatch**:
   - Element Plus tables for inventory.
   - Workflow cards for dispatch requests.
