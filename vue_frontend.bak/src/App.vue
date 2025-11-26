<template>
	<div class="dashboard v-ui" :data-theme="'light'">

		<aside class="aside">
			<TheSidebar></TheSidebar>
		</aside>
		<div class="collactions"></div>
		<div class="games"></div>
		<main class="main"></main>
	</div>
</template>

<!-- --------------------------------------------------------------- -->

<script setup lang="ts">
import { storeToRefs } from 'pinia';
import { useDashboardStore } from './stores/dashboardStore';
import TheSidebar from './components/TheSidebar.vue';


// - Store -------------------------------------------------------------

const dashboardStore = useDashboardStore();
const { getIsLoading, getTheme } = storeToRefs(dashboardStore);

</script>

<!-- --------------------------------------------------------------- -->

<style lang="scss" scoped>
// Placeholder comment to ensure global styles are imported correctly

.dashboard {

	width: 100vw;
	min-height: 100vh;

	display: grid;
	grid-template-columns: 3rem 11rem 11rem 1fr;
}



.aside {
	border-right: 1px solid var(--color-slate-300);
	background-color: #0f172a;
}

.collactions {
	border-right: 1px solid var(--color-slate-300);
	background-color: var(--color-slate-200);
}
.games {
	border-right: 1px solid var(--color-slate-300);
	background-color: var(--color-slate-200);
}

.main {
	background-color: var(--color-slate-100);
}
</style>