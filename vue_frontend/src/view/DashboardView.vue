<template>
		<div class="dashboard v-ui" :data-theme="getTheme" ref="dashboardRef">

			<VueLoadingOverlay :active="getIsLoading" :is-full-page="true" :lock-scroll="true" 
			 	:width="256" :height="256" transition="fade" :opacity="0.5" />

			<aside class="aside">
				<TheSidebar></TheSidebar>
			</aside>
			<div class="collactions">
				<TheCollections></TheCollections>
			</div>
			<div class="games">
				<TheGames></TheGames>
			</div>
			<main class="main">
                <RouterView v-slot="{ Component, route }">
                    <KeepAlive>
                        <component :is="Component" :key="route.name" />
                    </KeepAlive>
                </RouterView>
            </main>
		</div>
</template>

<!-- --------------------------------------------------------------- -->

<script setup lang="ts">
import VueLoadingOverlay from 'vue-loading-overlay';
import 'vue-loading-overlay/dist/css/index.css';
import { storeToRefs } from 'pinia';
import { useDashboardStore } from '../stores/dashboardStore';
import TheSidebar from '../components/TheSidebar.vue';
import { ref, watch } from 'vue';
import TheCollections from '../components/TheCollections.vue';
import TheGames from '../components/TheGames.vue';



// - Store -------------------------------------------------------------

const dashboardStore = useDashboardStore();
const { getIsLoading, getTheme, getIsMenuOpen } = storeToRefs(dashboardStore);

// - State -------------------------------------------------------------
const dashboardRef = ref<null | HTMLElement>(null)

// - Satchers ----------------------------------------------------------

watch(getIsMenuOpen, (val) => {
	if (!dashboardRef.value) return;
	if (val) {
		dashboardRef.value.classList.remove('dashboard--menu-close');
	} else {
		dashboardRef.value.classList.add('dashboard--menu-close');
	}
})
</script>

<!-- --------------------------------------------------------------- -->

<style lang="scss" scoped>
// Placeholder comment to ensure global styles are imported correctly

.dashboard {

	width: 100vw;
	height: 100vh;

	display: grid;
	grid-template-columns: 3rem 14rem 14rem 1fr;

	// transition: all .2s ease;

	&--menu-close {
		grid-template-columns: 3rem 0rem 0rem 1fr !important;
		// width: 50vw !important;
	}
}

.aside {
	border-right: 1px solid var(--color-slate-300);
	background-color: var(--color-slate-700);
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