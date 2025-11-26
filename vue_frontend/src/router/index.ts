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
				{path: 'collection/:collectionId/game/:gameId', name: 'gameView', component: GameView, props: true},
				{
				   path: '/game/create',
				   name: 'gameCreateView',
				   component: GameCreateView
				}
			]
		 },
	],
})

export default router
