<template>
	<div class="app">
		<RouterView v-slot="{ Component, route }">
			<KeepAlive>
				<component :is="Component" :key="route.name" />
			</KeepAlive>
		</RouterView>
	</div>
</template>

<!-- --------------------------------------------------------------- -->

<script setup lang="ts">
import { watch } from 'vue';
import { useAuthStore } from './stores/authStore';
import { useChessLogStore } from './stores/chessLogStore';


const authStore = useAuthStore();
const chessLogStore = useChessLogStore();


watch(()=>authStore.isAuthenticated, async (val) => {
	if (val) {	
		await chessLogStore.fetchProfile();
	} 
}, {
	 immediate: true,
})

</script>

<!-- --------------------------------------------------------------- -->

<style lang="scss" scoped>
// Placeholder comment to ensure global styles are imported correctly


</style>