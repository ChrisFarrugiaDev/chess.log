import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { useLocalStorage } from '@vueuse/core';

export const useDashboardStore = defineStore('dashboardStore', () => {


	// ---- State ------------------------------------------------------

	const isMenuOpen = ref(false);
	const isLoading = ref(false);

	type Theme = 'light' | 'dark' ;
	const themeLocal = useLocalStorage<string>('iotrack.theme', 'light');

	const theme = computed({
		get: () => themeLocal.value,
		set: (theme: Theme) => themeLocal.value = theme,
	})


	// ---- Getters ----------------------------------------------------

	const getIsMenuOpen = computed(()=>{
		return isMenuOpen.value;
	});

	const getIsLoading = computed(() => isLoading.value )

	const getTheme = computed(() => theme.value);

	// ---- Actions ----------------------------------------------------

	function toggleTheme() {
		switch (theme.value) {
			case 'light':
				theme.value = 'dark';
				break;
			default:
				theme.value = 'light';
				break;		
		}

	}

	function toggleMenuState() {

		isMenuOpen.value = !isMenuOpen.value;
	}

	function updateUserMenuState(val: boolean) {
		isMenuOpen.value = val;
	}

	function setIsLoading(val: boolean) {
		isLoading.value = val;
	}

	// - Expose --------------------------------------------------------
	return {
		getIsMenuOpen,
		toggleMenuState,
		updateUserMenuState,
		getIsLoading,
		setIsLoading,
		getTheme,
		toggleTheme,
	};
})
