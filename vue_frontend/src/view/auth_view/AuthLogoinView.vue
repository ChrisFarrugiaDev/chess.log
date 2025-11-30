<template>
    <div class="auth v-ui" data-theme="light">
        <div class="auth__card">

            <!-- Title -->
            <h1 class="auth__title">
                <span>ChessLog</span>
                <svg class="auth__title-icon" aria-hidden="true">
                    <use xlink:href="@/ui/svg/sprite.svg#icon-pawn" />
                </svg>
            </h1>

            <p class="auth__subtitle">Sign in to continue</p>

            <!-- Error -->
            <p v-if="errorMessage" class="auth__error">{{ errorMessage }}</p>

            <!-- Form -->
            <form class="auth__form" @submit.prevent="handleSubmit">

                <div class="auth__field">
                    <label class="auth__label">Email</label>
                    <input v-model="email" type="email" placeholder="you@example.com" class="auth__input" required
                        autofocus />
                </div>

                <div class="auth__field auth__field--password">
                    <label class="auth__label">Password</label>

                    <div class="auth__password-wrapper">
                        <input v-model="password" :type="showPassword ? 'text' : 'password'" placeholder="••••••••"
                            class="auth__input" required />
                        <button type="button" class="auth__toggle-password" @click="showPassword = !showPassword">
                            {{ showPassword ? 'Hide' : 'Show' }}
                        </button>
                    </div>
                </div>

                <button type="submit" class="auth__btn vbtn--indigo" :disabled="loading">
                    <span v-if="!loading">Login</span>
                    <span v-else>Loading…</span>
                </button>
            </form>

            <!-- Links -->
            <div class="auth__links">
                <router-link class="auth__link auth__link--primary" to="/register">
                    Create account
                </router-link>
                <router-link class="auth__link" to="/forgot-password">
                    Forgot password?
                </router-link>
            </div>

            <p class="auth__footer">v1.0.0 — © 2025 ChessLog</p>

        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

const email = ref("");
const password = ref("");

const showPassword = ref(false);
const loading = ref(false);
const errorMessage = ref("");

async function handleSubmit() {
    loading.value = true;
    errorMessage.value = "";

    setTimeout(() => {
        loading.value = false;

        if (email.value.length < 4 || password.value.length < 4) {
            errorMessage.value = "Invalid email or password.";
            return;
        }

        console.log("Login successful");
    }, 1000);
}
</script>

<style scoped lang="scss">
/* ===============================================
   Block: auth
================================================ */
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
        width: 340px;
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

    /* ------------------------------------------- */
    /* Title */
    &__title {
        color: var(--color-slate-900);
        font-size: 1.9rem;
        text-align: center;
        font-weight: 700;
        display: flex;
        justify-content: center;
        align-items: center;
        gap: .3rem;

        &-icon {
            width: 2.5rem;
            height: 2.5rem;
            fill: var(--color-slate-700);
            transition: transform .25s ease;
        }

        &:hover &-icon {
            transform: translateY(-5px) scale(1.05);
        }
    }

    &__subtitle {
        color: var(--color-text-3);
        text-align: center;
        font-size: .9rem;
        margin-top: -0.8rem;
    }

    /* ------------------------------------------- */
    /* Messages */
    &__error {
        color: var(--color-red-500);
        text-align: center;
        font-size: .85rem;
        margin-top: -0.5rem;
    }

    /* ------------------------------------------- */
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
        transition: all .2s ease;
        box-shadow: inset 0 1px 3px rgba(0, 0, 0, 0.35);

        &:focus {
            border-color: var(--color-indigo-500);
            box-shadow: 0 0 6px rgba(99, 102, 241, 0.25);
        }
    }

    /* Password wrapper */
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
        padding: 0;
        font-size: .75rem;
        color: var(--color-indigo-500);
        cursor: pointer;
    }

    /* ------------------------------------------- */
    /* Submit button */
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

    /* ------------------------------------------- */
    /* Links */
    &__links {
        margin-top: -.5rem;
        display: flex;
        justify-content: space-between;
    }

    &__link {
        font-size: .85rem;
        color: var(--color-indigo-500);
        text-decoration: none;
        transition: .2s;

        &:hover {
            color: var(--color-indigo-400);
        }

        &--primary {
            color: var(--color-slate-600);

            &:hover {
                color: var(--color-indigo-500);
            }
        }
    }

    &__footer {
        text-align: center;
        font-size: .75rem;
        color: var(--color-text-4);
        opacity: .7;
    }
}

/* Animation */
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
</style>
