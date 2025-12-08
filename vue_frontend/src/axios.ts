import axios, { type InternalAxiosRequestConfig } from "axios";
import { useAuthStore } from "./stores/authStore";

// ---------------------------------------------------------------------

const apiClient = axios.create({
    timeout: 15000,
});

// ---------------------------------------------------------------------
// REQUEST INTERCEPTOR — Attach JWT

apiClient.interceptors.request.use(
    (config: InternalAxiosRequestConfig<any>) => {
        const authStore = useAuthStore();

        if (authStore.isAuthenticated) {
            const token = authStore.getJwt;
            config.headers.Authorization = `Bearer ${token}`;
        }

        return config;
    },
    (error) => {
        console.error("! axios.request error !", error);
        return Promise.reject(error);
    }
);

// ---------------------------------------------------------------------
// RESPONSE INTERCEPTOR — Handle 401

apiClient.interceptors.response.use(
    // SUCCESS → return response normally
    (response) => {
        return response;
    },

    // ERROR → handle token expiration / unauthorized
    (error) => {
        const status = error?.response?.status;
        const authStore = useAuthStore();

        if (status === 401 && authStore.isAuthenticated) {
            console.warn("401 caught → clearing JWT & redirecting to login");     

            authStore.clearJwt();

            // redirect
            window.location.assign("/login");
        }

        return Promise.reject(error);
    }
);

// ---------------------------------------------------------------------

export default apiClient;
