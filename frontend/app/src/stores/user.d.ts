import { type LoginRequest, type RegisterRequest } from '@/api/auth';
export declare const useUserStore: import("pinia").StoreDefinition<"user", Pick<{
    token: import("vue").Ref<string | null, string | null>;
    role: import("vue").Ref<string | null, string | null>;
    login: (data: LoginRequest, selectedRole: string) => Promise<boolean>;
    register: (data: RegisterRequest) => Promise<boolean>;
    setToken: (newToken: string) => void;
    removeToken: () => void;
    setRole: (newRole: string) => void;
}, "role" | "token">, Pick<{
    token: import("vue").Ref<string | null, string | null>;
    role: import("vue").Ref<string | null, string | null>;
    login: (data: LoginRequest, selectedRole: string) => Promise<boolean>;
    register: (data: RegisterRequest) => Promise<boolean>;
    setToken: (newToken: string) => void;
    removeToken: () => void;
    setRole: (newRole: string) => void;
}, never>, Pick<{
    token: import("vue").Ref<string | null, string | null>;
    role: import("vue").Ref<string | null, string | null>;
    login: (data: LoginRequest, selectedRole: string) => Promise<boolean>;
    register: (data: RegisterRequest) => Promise<boolean>;
    setToken: (newToken: string) => void;
    removeToken: () => void;
    setRole: (newRole: string) => void;
}, "login" | "register" | "setToken" | "removeToken" | "setRole">>;
