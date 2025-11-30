import AuthForgotPasswordView from '@/view/auth_view/AuthForgotPasswordView.vue'
import AuthLogoinView from '@/view/auth_view/AuthLogoinView.vue'
import AuthRegisterView from '@/view/auth_view/AuthRegisterView.vue'
import AuthVerifyEmailView from '@/view/auth_view/AuthVerifyEmailView.vue'
import DashboardView from '@/view/DashboardView.vue'
import GameCreateView from '@/view/GameCreateView.vue'
import GameView from '@/view/GameView.vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: '/',
			name: 'dashboadView',
			component: DashboardView,
			children: [
				{ path: 'collection/:collectionId/game/:gameId', name: 'gameView', component: GameView, props: true },
				{
					path: '/game/create',
					name: 'gameCreateView',
					component: GameCreateView
				}
			]
		},
		{ path: '/login', name: 'loginView', component: AuthLogoinView, props: true },
		{ path: '/register', name: 'registerView', component: AuthRegisterView, props: true },
		{ path: '/forgot-password', name: 'forgotPasswordView', component: AuthForgotPasswordView, props: true },
		{ path: "/verify-email", name: "verifyEmail", component: AuthVerifyEmailView, }
	],
})

export default router
