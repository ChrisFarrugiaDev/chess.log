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

            <p class="auth__subtitle">Enter your new password</p>

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
            <form v-if="!successMessage" class="auth__form" @submit.prevent="submitResetPassword">

                <div class="auth__field auth__field--password">
                    <label class="auth__label">New Password</label>

                    <div class="auth__password-wrapper">
                        <input 
                            v-model="password1"
                            :type="showPassword ? 'text' : 'password'"
                            class="auth__input"
                            required 
                            placeholder="••••••••"
                        />
                        <button type="button" class="auth__toggle-password" @click="showPassword = !showPassword">
                            {{ showPassword ? 'Hide' : 'Show' }}
                        </button>
                    </div>
                </div>

                <div class="auth__field auth__field--password">
                    <label class="auth__label">Confirm Password</label>

                    <div class="auth__password-wrapper">
                        <input 
                            v-model="password2"
                            :type="showPassword ? 'text' : 'password'"
                            class="auth__input"
                            required 
                            placeholder="••••••••"
                        />
                    </div>
                </div>

                <button type="submit" class="auth__btn vbtn--indigo" :disabled="loading">
                    <span v-if="!loading">Reset Password</span>
                    <span v-else>Saving…</span>
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

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import axios from "axios";
import { useApiError } from "@/composables/useApiError";
import { useAppStore } from "@/stores/appStore";
import { useDashboardStore } from "@/stores/dashboardStore";
import { storeToRefs } from "pinia";

// - router ------------------------------------------------------------
const router = useRouter();
const route = useRoute();

// - store -------------------------------------------------------------

const appStore = useAppStore();
const dashboardStore = useDashboardStore();
const { getTheme } = storeToRefs(dashboardStore);

// - composable --------------------------------------------------------
const { errorMessage, handleApiError } = useApiError();

// - state -------------------------------------------------------------
const loading = ref(false);
const successMessage = ref("");

const password1 = ref("");
const password2 = ref("");
const showPassword = ref(false);

const token = ref<string | null>(null);

// - hooks -------------------------------------------------------------
onMounted(() => {
    token.value = route.query.token as string | null;

    if (!token.value) {
        errorMessage.value = "Invalid or missing reset token.";
    }
});

// - methods -----------------------------------------------------------

async function submitResetPassword() {

    if (!token.value) {
        errorMessage.value = "Missing token.";
        return;
    }

    if (password1.value.length < 6) {
        errorMessage.value = "Password must be at least 6 characters.";
        return;
    }

    if (password1.value !== password2.value) {
        errorMessage.value = "Passwords do not match.";
        return;
    }

    loading.value = true;
    errorMessage.value = "";
    successMessage.value = "";

    const url = `${appStore.getAppUrl}/api/auth/reset-password`;
    const payload = {
        token: token.value,
        password1: password1.value,
        password2: password2.value
    };

    try {
        const r = await axios.post(url, payload);
        successMessage.value = r.data.message;

        // Redirect after 2 seconds
        setTimeout(() => {
            router.push({ name: "loginView" });
        }, 2000);

    } catch (err) {
        handleApiError(err);
    }
    finally {
        loading.value = false;
    }
}
</script>

<style scoped lang="scss">
.auth {
    width: 100vw;
    height: 100vh;
    background: linear-gradient(145deg, var(--color-slate-100), var(--color-slate-200));
    display: flex;
    justify-content: center;
    align-items: center;
    font-family: var(--font-primary);
    padding: 1rem;
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
            transition: .25s;
        }

        &:hover &-icon {
            transform: translateY(-4px) scale(1.05);
        }
    }

    &__subtitle {
        color: var(--color-text-3);
        text-align: center;
        font-size: .9rem;
        margin-top: -0.8rem;
    }

    &__error {
        color: var(--color-red-500);
        text-align: center;
        font-size: .85rem;
        margin-top: -0.5rem;
    }

    &__message {
        color: var(--color-green-500);
        text-align: center;
        font-size: .85rem;
        margin-top: -0.5rem;
    }

    /* Form */
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
        transition: .2s;
        box-shadow: inset 0 1px 3px rgba(0, 0, 0, 0.35);

        &:focus {
            border-color: var(--color-indigo-500);
            box-shadow: 0 0 6px rgba(99, 102, 241, 0.25);
        }
    }

    &__password-wrapper {
        position: relative;
    }

    &__toggle-password {
        position: absolute;
        top: 50%;
        right: 8px;
        transform: translateY(-50%);
        background: none;
        border: none;
        font-size: .75rem;
        color: var(--color-indigo-500);
        cursor: pointer;
    }

    &__btn {
        width: 100%;
        padding: .75rem;
        border-radius: 6px;
        font-size: 1rem;
        font-weight: 600;
        border: none;
        cursor: pointer;
        transition: .2s;
        color: var(--color-slate-200);
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
        transition: .2s;

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
