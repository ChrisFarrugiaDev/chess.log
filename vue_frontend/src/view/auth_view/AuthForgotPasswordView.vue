<template>
    <div class="auth v-ui" :data-theme="getTheme">
        <div class="auth__card">

            <!-- Title -->
            <h1 class="auth__title">
                <span>Reset Password</span>
                <svg class="auth__title-icon" aria-hidden="true">
                    <use xlink:href="@/ui/svg/sprite.svg#icon-pawn" />
                </svg>
            </h1>

            <p class="auth__subtitle">Enter your email to receive a reset link</p>

                        <!-- Loading -->
            <div v-if="loading" class="auth__loading">
                <div class="auth__spinner"></div>
                <p class="auth__loading-text">Please wait…</p>
            </div>

            <!-- Error -->
            <p v-if="errorMessage" class="auth__error">{{ errorMessage }}</p>

            <!-- Success -->
            <p v-if="successMessage" class="auth__success">{{ successMessage }}</p>

            <!-- Form -->
            <form class="auth__form" @submit.prevent="submitForgotPassword">

                <div class="auth__field">
                    <label class="auth__label">Email</label>
                    <input v-model="email" type="email" class="auth__input" placeholder="you@example.com" required />
                </div>

                <button type="submit" class="auth__btn vbtn--indigo" :disabled="loading">
                    <span v-if="!loading">Send Reset Link</span>
                    <span v-else>Sending…</span>
                </button>

            </form>

            <!-- Links -->
            <div class="auth__links">
                <router-link class="auth__link" to="/login">Back to login</router-link>
            </div>

            <p class="auth__footer">v1.0.0 — © 2025 ChessLog</p>

        </div>
    </div>
</template>

<!-- --------------------------------------------------------------- -->

<script setup lang="ts">
import { useApiError } from "@/composables/useApiError";
import { useAppStore } from "@/stores/appStore";
import axios from "@/axios";
import { ref } from "vue";
import { useDashboardStore } from "@/stores/dashboardStore";
import { storeToRefs } from "pinia";

// - composable --------------------------------------------------------
const { errorMessage, handleApiError } = useApiError();

// - store -------------------------------------------------------------

const appStore = useAppStore();
const dashboardStore = useDashboardStore();
const { getTheme } = storeToRefs(dashboardStore);

// - state -------------------------------------------------------------

const email = ref("");
const loading = ref(false);

const successMessage = ref("");

// - methods -----------------------------------------------------------

async function submitForgotPassword() {
    loading.value = true;
    errorMessage.value = "";
    successMessage.value = "";

    // Basic validation
    if (!email.value.includes("@")) {
        loading.value = false;
        errorMessage.value = "Please enter a valid email.";
        return;
    }

    try {
        const url = `${appStore.getAppUrl}/api/auth/forgot-password`;

        const r = await axios.post(url, {email: email.value});
        // --- Success message ---
        successMessage.value =
            r.data?.message ||
            "If an account exists, a reset link has been sent to your email.";

        // Clear inputs
        email.value = "";


    } catch (err) {
        handleApiError(err);
    } finally {
        loading.value = false;
    }

    // // Simulate request
    // setTimeout(() => {
    //     loading.value = false;
    //     successMessage.value =
    //         "If an account exists, a reset link has been sent to your email.";
    // }, 1000);
}
</script>

<!-- --------------------------------------------------------------- -->

<style scoped lang="scss">
/* ------------------------------------------------------
   Auth Forgot Password — matches Login + Register
------------------------------------------------------ */

.auth {
    width: 100vw;
    height: 100vh;
    background: linear-gradient(145deg, var(--color-slate-100), var(--color-slate-200));
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 1rem;
    font-family: var(--font-primary);
    padding-bottom: 5%;

    &__card {
        width: 360px;
        padding: 2rem;
        border-radius: 14px;
        background: var(--color-slate-200);
        border: 1px solid var(--color-slate-400);
        display: flex;
        flex-direction: column;
        gap: 1.6rem;
        box-shadow: 0 8px 28px rgba(0, 0, 0, 0.22);
        animation: fadeIn .4s ease-out;
  
    }

    &__title {
        color: var(--color-slate-900);
        font-size: 1.8rem;
        text-align: center;
        font-weight: 700;
        display: flex;
        justify-content: center;
        align-items: center;
        gap: .3rem;

        &-icon {
            width: 2.4rem;
            height: 2.4rem;
            fill: var(--color-slate-700);
            transition: .25s ease;
        }

        &:hover &-icon {
            transform: translateY(-4px) scale(1.05);
        }
    }

    &__subtitle {
        text-align: center;
        font-size: .9rem;
        color: var(--color-text-3);
        margin-top: -0.8rem;
    }

    &__error {
        color: var(--color-red-500);
        text-align: center;
        font-size: .85rem;
        margin-top: -0.5rem;
    }

    &__success {
        color: var(--color-green-500);
        text-align: center;
        font-size: .85rem;
        margin-top: -0.5rem;
    }

    &__form {
        display: flex;
        flex-direction: column;
        gap: 1.3rem;
    }

    &__field {
        display: flex;
        flex-direction: column;
        gap: .3rem;
    }

    &__label {
        font-size: .8rem;
        color: var(--color-text-4);
    }

    &__input {
        width: 100%;
        padding: .75rem;
        border-radius: 6px;
        background: var(--color-slate-300);
        border: 1px solid var(--color-slate-500);
        color: var(--color-slate-950);
        font-size: .95rem;
        outline: none;
        transition: all .2s ease;
        box-shadow: inset 0 1px 3px rgba(0, 0, 0, 0.35);

        &:focus {
            border-color: var(--color-indigo-500);
            box-shadow: 0 0 6px rgba(99, 102, 241, 0.25);
        }
    }

    &__btn {
        width: 100%;
        padding: .75rem;
        border-radius: 6px;
        font-weight: 600;
        border: none;
        cursor: pointer;
        color: var(--color-slate-200);
        transition: .2s;
        box-shadow: 0 3px 12px rgba(0, 0, 0, 0.4);

        &:disabled {
            opacity: .6;
            cursor: not-allowed;
        }
    }

    &__links {
        text-align: center;
        margin-top: -0.5rem;
    }

    &__link {
        font-size: .85rem;
        color: var(--color-indigo-500);
        text-decoration: none;
        transition: color .2s ease;

        &:hover {
            color: var(--color-indigo-400);
        }
    }

    &__footer {
        text-align: center;
        font-size: .75rem;
        color: var(--color-text-4);
        opacity: .7;
    }

        /* Loading state */
    &__loading {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: .8rem;
    }

    &__spinner {
        width: 28px;
        height: 28px;
        border: 3px solid var(--color-slate-400);
        border-top-color: var(--color-indigo-500);
        border-radius: 50%;
        animation: spin 0.8s linear infinite;
    }

    &__loading-text {
        color: var(--color-text-4);
        font-size: .85rem;
    }
}

@keyframes fadeIn {
    from {
        opacity: 0;
        transform: translateY(10px);
    }

    to {
        opacity: 1;
        transform: translateY(0);
    }
}

@keyframes spin {
    to {
        transform: rotate(360deg);
    }
}
</style>
