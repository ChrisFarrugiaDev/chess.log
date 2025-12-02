<template>
    <div class="auth v-ui" data-theme="light">
        <div class="auth__card">

            <!-- Title -->
            <h1 class="auth__title">
                <span>Create Account</span>
                <svg class="auth__title-icon" aria-hidden="true">
                    <use xlink:href="@/ui/svg/sprite.svg#icon-pawn" />
                </svg>
            </h1>

            <p class="auth__subtitle">Start your ChessLog journey</p>

            <!-- Error -->
            <p v-if="errorMessage" class="auth__error">{{ errorMessage }}</p>

            <!-- Register Form -->
            <form class="auth__form" @submit.prevent="submitRegister">

                <!-- Name -->
                <div class="auth__field">
                    <label class="auth__label">Name</label>
                    <input v-model="name" type="text" class="auth__input" placeholder="Your name" required />
                </div>

                <!-- Email -->
                <div class="auth__field">
                    <label class="auth__label">Email</label>
                    <input v-model="email" type="email" class="auth__input" placeholder="you@example.com" required />
                </div>

                <!-- Password -->
                <div class="auth__field auth__field--password">
                    <label class="auth__label">Password</label>

                    <div class="auth__password-wrapper">
                        <input v-model="password" :type="showPassword ? 'text' : 'password'" class="auth__input"
                            placeholder="••••••••" required />
                        <button type="button" class="auth__toggle-password" @click="showPassword = !showPassword">
                            {{ showPassword ? 'Hide' : 'Show' }}
                        </button>
                    </div>
                </div>

                <!-- Confirm Password -->
                <div class="auth__field auth__field--password">
                    <label class="auth__label">Confirm Password</label>

                    <div class="auth__password-wrapper">
                        <input v-model="confirmPassword" :type="showPassword ? 'text' : 'password'" class="auth__input"
                            placeholder="••••••••" required />
                    </div>
                </div>

                <!-- Register Button -->
                <button type="submit" class="auth__btn vbtn--indigo" :disabled="loading">
                    <span v-if="!loading">Create Account</span>
                    <span v-else>Creating…</span>
                </button>

            </form>

            <!-- Links -->
            <div class="auth__links">
                <router-link class="auth__link" to="/login">Already have an account?</router-link>
            </div>

            <p class="auth__footer">v1.0.0 — © 2025 ChessLog</p>

        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

const name = ref("");
const email = ref("");
const password = ref("");
const confirmPassword = ref("");

const showPassword = ref(false);
const loading = ref(false);
const errorMessage = ref("");

async function submitRegister() {
    loading.value = true;
    errorMessage.value = "";

    // Basic validation
    if (password.value !== confirmPassword.value) {
        errorMessage.value = "Passwords do not match.";
        loading.value = false;
        return;
    }

    if (password.value.length < 6) {
        errorMessage.value = "Password must be at least 6 characters.";
        loading.value = false;
        return;
    }

    // Simulate request for now
    setTimeout(() => {
        loading.value = false;
        console.log("REGISTER OK:", { name: name.value, email: email.value });
    }, 1000);
}
</script>

<style scoped lang="scss">
/* ------------------------------------------
   Same BEM + SCSS structure as Login view
------------------------------------------ */

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
</style>
