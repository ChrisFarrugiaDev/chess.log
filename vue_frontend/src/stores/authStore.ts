import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { useLocalStorage } from '@vueuse/core';

export const useAuthStore = defineStore('authStore', () => {

    // ---- State ------------------------------------------------------
    const jwtLocal = useLocalStorage<string | null>("jwt", null);


    const jwt = computed({

        get: () => jwtLocal.value,
        set: (val: string | null) => { jwtLocal.value = val; }
    });

    const redirectTo = ref("dashboadView");


    // ---- Getters ----------------------------------------------------
    const isAuthenticated = computed(() => {
        return Boolean(jwt.value);
    });

    const getJwt = computed(() => {

        return jwt.value;
    });

    const getRedirectTo = computed(() => {
        return redirectTo.value;
    });


    // ---- Actions ----------------------------------------------------
    function setJwt(token: string | null) {
        jwt.value = token;
    }

    function clearJwt() {
        jwt.value = null;
    }

    function resetAuthStore() {
        jwt.value = null;
    }

    function setRedirectTo(payload: any) {
        redirectTo.value = payload.name;
    }
    // - Log counter ---------------------------------------------------

    const logCounter = ref(0);

    const getLogCounter = computed(() => {
        return logCounter.value;
    })
    function updateLogCounter() {
        logCounter.value++;
        console.log(logCounter.value)
    }
    // - Expose --------------------------------------------------------

    return {
        // --- State ---
        jwt,
        redirectTo,
        logCounter,

        // --- Getters ---
        isAuthenticated,
        getJwt,
        getRedirectTo,
        getLogCounter,

        // --- Actions ---
        setJwt,
        clearJwt,
        resetAuthStore,
        setRedirectTo,
        updateLogCounter,
    };
});
