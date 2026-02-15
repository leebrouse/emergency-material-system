export interface DispatchRequest {
    id: number;
    region: string;
    urgency: 'High' | 'Medium' | 'Critical';
    items: string;
    status: 'Pending' | 'Approved' | 'In Transit' | 'Delivered';
    coordinates?: [number, number];
    warehouse?: string;
    estimatedTime?: string;
    driver?: string;
}
export declare const useDispatchStore: import("pinia").StoreDefinition<"dispatch", Pick<{
    requests: import("vue").Ref<{
        id: number;
        region: string;
        urgency: "High" | "Medium" | "Critical";
        items: string;
        status: "Pending" | "Approved" | "In Transit" | "Delivered";
        coordinates?: [number, number] | undefined;
        warehouse?: string | undefined;
        estimatedTime?: string | undefined;
        driver?: string | undefined;
    }[], DispatchRequest[] | {
        id: number;
        region: string;
        urgency: "High" | "Medium" | "Critical";
        items: string;
        status: "Pending" | "Approved" | "In Transit" | "Delivered";
        coordinates?: [number, number] | undefined;
        warehouse?: string | undefined;
        estimatedTime?: string | undefined;
        driver?: string | undefined;
    }[]>;
    activeDispatches: import("vue").ComputedRef<{
        id: number;
        region: string;
        urgency: "High" | "Medium" | "Critical";
        items: string;
        status: "Pending" | "Approved" | "In Transit" | "Delivered";
        coordinates?: [number, number] | undefined;
        warehouse?: string | undefined;
        estimatedTime?: string | undefined;
        driver?: string | undefined;
    }[]>;
    pendingRequests: import("vue").ComputedRef<{
        id: number;
        region: string;
        urgency: "High" | "Medium" | "Critical";
        items: string;
        status: "Pending" | "Approved" | "In Transit" | "Delivered";
        coordinates?: [number, number] | undefined;
        warehouse?: string | undefined;
        estimatedTime?: string | undefined;
        driver?: string | undefined;
    }[]>;
    fetchDemands: () => Promise<void>;
    approveRequest: (id: number, warehouse: string, eta: string) => Promise<void>;
    startDispatch: (id: number, driver: string) => void;
    createRequest: (data: {
        material_id: number;
        quantity: number;
        urgency: "L1" | "L2" | "L3";
        region: string;
    }) => Promise<void>;
    getSuggestion: (id: number) => Promise<any>;
    isLoading: import("vue").Ref<boolean, boolean>;
}, "requests" | "isLoading">, Pick<{
    requests: import("vue").Ref<{
        id: number;
        region: string;
        urgency: "High" | "Medium" | "Critical";
        items: string;
        status: "Pending" | "Approved" | "In Transit" | "Delivered";
        coordinates?: [number, number] | undefined;
        warehouse?: string | undefined;
        estimatedTime?: string | undefined;
        driver?: string | undefined;
    }[], DispatchRequest[] | {
        id: number;
        region: string;
        urgency: "High" | "Medium" | "Critical";
        items: string;
        status: "Pending" | "Approved" | "In Transit" | "Delivered";
        coordinates?: [number, number] | undefined;
        warehouse?: string | undefined;
        estimatedTime?: string | undefined;
        driver?: string | undefined;
    }[]>;
    activeDispatches: import("vue").ComputedRef<{
        id: number;
        region: string;
        urgency: "High" | "Medium" | "Critical";
        items: string;
        status: "Pending" | "Approved" | "In Transit" | "Delivered";
        coordinates?: [number, number] | undefined;
        warehouse?: string | undefined;
        estimatedTime?: string | undefined;
        driver?: string | undefined;
    }[]>;
    pendingRequests: import("vue").ComputedRef<{
        id: number;
        region: string;
        urgency: "High" | "Medium" | "Critical";
        items: string;
        status: "Pending" | "Approved" | "In Transit" | "Delivered";
        coordinates?: [number, number] | undefined;
        warehouse?: string | undefined;
        estimatedTime?: string | undefined;
        driver?: string | undefined;
    }[]>;
    fetchDemands: () => Promise<void>;
    approveRequest: (id: number, warehouse: string, eta: string) => Promise<void>;
    startDispatch: (id: number, driver: string) => void;
    createRequest: (data: {
        material_id: number;
        quantity: number;
        urgency: "L1" | "L2" | "L3";
        region: string;
    }) => Promise<void>;
    getSuggestion: (id: number) => Promise<any>;
    isLoading: import("vue").Ref<boolean, boolean>;
}, "activeDispatches" | "pendingRequests">, Pick<{
    requests: import("vue").Ref<{
        id: number;
        region: string;
        urgency: "High" | "Medium" | "Critical";
        items: string;
        status: "Pending" | "Approved" | "In Transit" | "Delivered";
        coordinates?: [number, number] | undefined;
        warehouse?: string | undefined;
        estimatedTime?: string | undefined;
        driver?: string | undefined;
    }[], DispatchRequest[] | {
        id: number;
        region: string;
        urgency: "High" | "Medium" | "Critical";
        items: string;
        status: "Pending" | "Approved" | "In Transit" | "Delivered";
        coordinates?: [number, number] | undefined;
        warehouse?: string | undefined;
        estimatedTime?: string | undefined;
        driver?: string | undefined;
    }[]>;
    activeDispatches: import("vue").ComputedRef<{
        id: number;
        region: string;
        urgency: "High" | "Medium" | "Critical";
        items: string;
        status: "Pending" | "Approved" | "In Transit" | "Delivered";
        coordinates?: [number, number] | undefined;
        warehouse?: string | undefined;
        estimatedTime?: string | undefined;
        driver?: string | undefined;
    }[]>;
    pendingRequests: import("vue").ComputedRef<{
        id: number;
        region: string;
        urgency: "High" | "Medium" | "Critical";
        items: string;
        status: "Pending" | "Approved" | "In Transit" | "Delivered";
        coordinates?: [number, number] | undefined;
        warehouse?: string | undefined;
        estimatedTime?: string | undefined;
        driver?: string | undefined;
    }[]>;
    fetchDemands: () => Promise<void>;
    approveRequest: (id: number, warehouse: string, eta: string) => Promise<void>;
    startDispatch: (id: number, driver: string) => void;
    createRequest: (data: {
        material_id: number;
        quantity: number;
        urgency: "L1" | "L2" | "L3";
        region: string;
    }) => Promise<void>;
    getSuggestion: (id: number) => Promise<any>;
    isLoading: import("vue").Ref<boolean, boolean>;
}, "fetchDemands" | "approveRequest" | "startDispatch" | "createRequest" | "getSuggestion">>;
