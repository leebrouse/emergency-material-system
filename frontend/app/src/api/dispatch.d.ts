export interface DemandItem {
    material_id: number;
    quantity: number;
}
export interface Demand {
    id: number;
    location?: string;
    priority?: string;
    status: string;
    description: string;
    items?: DemandItem[];
    created_at?: string;
    material_id?: number;
    quantity?: number;
    urgency?: string;
    target_area?: string;
}
export declare const dispatchApi: {
    getDemands: (params?: {
        page: number;
        page_size: number;
        status?: string;
    }) => Promise<import("axios").AxiosResponse<any, any, {}>>;
    updateDemandStatus: (id: number, action: "approve" | "reject", remark?: string) => Promise<import("axios").AxiosResponse<any, any, {}>>;
    createOrder: (data: {
        request_id: number;
        allocations: {
            inventory_id: number;
            quantity: number;
        }[];
    }) => Promise<import("axios").AxiosResponse<any, any, {}>>;
    createRequest: (data: {
        material_id: number;
        quantity: number;
        urgency_level: "L1" | "L2" | "L3";
        target_area: string;
        description?: string;
    }) => Promise<import("axios").AxiosResponse<any, any, {}>>;
    getSuggestion: (id: number) => Promise<import("axios").AxiosResponse<any, any, {}>>;
    getOrders: (params?: {
        page: number;
        page_size: number;
        status?: string;
    }) => Promise<import("axios").AxiosResponse<any, any, {}>>;
};
