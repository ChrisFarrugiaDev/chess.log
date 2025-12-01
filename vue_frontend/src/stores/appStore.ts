import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

declare global {
    interface Window {
        GO_DOCKERIZED: boolean | undefined;
        GO_APP_URL: string;
    }
}
export const useAppStore = defineStore('appStore', () => {

    
    // ---- State ------------------------------------------------------
    
    const appUrl = ref(window.GO_DOCKERIZED === true
        ? window.GO_APP_URL
        : import.meta.env.VITE_APP_URL
    );


    // ---- Getters ----------------------------------------------------

    const getAppUrl = computed(()=>{
        return appUrl.value;
    })

    // ---- Actions ----------------------------------------------------


    // - Expose --------------------------------------------------------
    return {
        getAppUrl
    };
});
