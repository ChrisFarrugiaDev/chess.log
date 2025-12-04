import { useAuthStore } from '@/stores/authStore'
import AuthForgotPasswordView from '@/view/auth_view/AuthForgotPasswordView.vue'
import AuthLogoinView from '@/view/auth_view/AuthLogoinView.vue'
import AuthRegisterView from '@/view/auth_view/AuthRegisterView.vue'
import AuthVerifyEmailView from '@/view/auth_view/AuthVerifyEmailView.vue'
import DashboardView from '@/view/DashboardView.vue'
import GameCreateView from '@/view/GameCreateView.vue'
import GameView from '@/view/GameView.vue'
import { storeToRefs } from 'pinia'
import { createRouter, createWebHistory, type NavigationGuardNext, type RouteLocationNormalized } from 'vue-router'

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: '/',
			name: 'dashboadView',
			component: DashboardView,
			props: true,
			meta: { requiresAuth: true },
			children: [
				{ path: 'collection/:collectionId/game/:gameId', name: 'gameView', component: GameView, props: true },
				{
					path: '/game/create',
					name: 'gameCreateView',
					component: GameCreateView,
					
				}
			]
		},
		{ path: '/login', name: 'loginView', component: AuthLogoinView, props: true,  meta: { requiresAuth: false, requiresGuest: true } },
		{ path: '/register', name: 'registerView', component: AuthRegisterView, props: true, meta: { requiresAuth: false, requiresGuest: true } },
		{ path: '/forgot-password', name: 'forgotPasswordView', component: AuthForgotPasswordView, props: true, meta: { requiresAuth: false, requiresGuest: true } },
		{ path: "/verify-email", name: "verifyEmail", component: AuthVerifyEmailView, meta: { requiresAuth: false, requiresGuest: true } }
	],
});


router.beforeEach((to, from, next) => {
    const authStore = useAuthStore();
    const { isAuthenticated } = storeToRefs(authStore);

    const requiresAuth = to.meta.requiresAuth !== false;
    const requiresGuest = to.meta.requiresGuest === true;

    // If user is already logged in → block access to login/register/etc
    if (requiresGuest && isAuthenticated.value) {
        return next({ name: 'dashboadView' });
    }

    // If route requires auth but not authenticated → redirect to login
    if (requiresAuth && !isAuthenticated.value) {
        authStore.setRedirectTo(to);
        return next({ name: 'loginView' });
    }

    return next();
});

export default router
