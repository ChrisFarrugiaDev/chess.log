<template>
    <div class="auth v-ui" data-theme="light">
        <div class="auth__card">

            <!-- Title -->
            <h1 class="auth__title">
                <span>Email Verification</span>
                <svg class="auth__title-icon" aria-hidden="true">
                    <use xlink:href="@/ui/svg/sprite.svg#icon-pawn" />
                </svg>
            </h1>

            <p class="auth__subtitle">We are verifying your email address…</p>

            <!-- Loading -->
            <div v-if="loading" class="auth__loading">
                <div class="auth__spinner"></div>
                <p class="auth__loading-text">Please wait…</p>
            </div>

            <!-- Success -->
            <p v-if="successMessage" class="auth__success">{{ successMessage }}</p>

            <!-- Error -->
            <p v-if="errorMessage" class="auth__error">{{ errorMessage }}</p>

            <!-- Actions -->
            <div v-if="!loading" class="auth__links">

                <router-link v-if="successMessage" class="auth__link" to="/login">
                    Go to login
                </router-link>

                <button v-if="errorMessage" class="auth__link auth__link--button" @click="resendEmail">
                    Resend verification email
                </button>

                <router-link class="auth__link" to="/login">
                    Back to login
                </router-link>
            </div>

            <!-- Footer -->
            <p class="auth__footer">v1.0.0 — © 2025 ChessLog</p>

        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();

const loading = ref(true);
const successMessage = ref("");
const errorMessage = ref("");

async function verifyEmail(token: string | null) {
    if (!token) {
        loading.value = false;
        errorMessage.value = "Invalid or missing verification token.";
        return;
    }

    // Simulate backend check
    setTimeout(() => {
        loading.value = false;

        if (token === "fail") {
            errorMessage.value = "Verification failed. Token may be expired.";
        } else {
            successMessage.value = "Your email has been successfully verified!";
        }
    }, 1200);
}

function resendEmail() {
    loading.value = true;
    errorMessage.value = "";
    successMessage.value = "";

    setTimeout(() => {
        loading.value = false;
        successMessage.value = "A new verification email has been sent.";
    }, 1200);
}

onMounted(() => {
    const token = route.query.token as string | null;
    verifyEmail(token);
});
</script>

<style scoped lang="scss">
/* ------------------------------------------------------
   Auth Verify Email — Matches all auth pages
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
        font-size: 1.8rem;
        font-weight: 700;
        color: var(--color-slate-900);
        display: flex;
        justify-content: center;
        align-items: center;
        gap: .3rem;
        text-align: center;

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

    &__success {
        color: var(--color-green-500);
        text-align: center;
        font-size: .85rem;
    }

    &__error {
        color: var(--color-red-500);
        text-align: center;
        font-size: .85rem;
    }

    &__links {
        display: flex;
        flex-direction: column;
        gap: .6rem;
        text-align: center;
    }

    &__link {
        font-size: .85rem;
        text-decoration: none;
        color: var(--color-indigo-500);
        cursor: pointer;
        transition: color .2s ease;

        &:hover {
            color: var(--color-indigo-400);
        }

        &--button {
            background: transparent;
            border: none;
            padding: 0;
            font-size: .85rem;
            cursor: pointer;
        }
    }

    &__footer {
        text-align: center;
        font-size: .75rem;
        color: var(--color-text-4);
        opacity: .7;
    }
}

/* Animations */
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
