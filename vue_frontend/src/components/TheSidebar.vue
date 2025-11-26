<template>
  <div class="sidebar">

        <div class="sidebar__item" @click="toggleMenu">
            <svg class="sidebar__svg ">
                <use xlink:href="@/ui/svg/sprite.svg#icon-menu"></use>
            </svg>
        </div>

        <div class="sidebar__item" @click="toggleTheme">
            <svg class="sidebar__svg ">
                <use xlink:href="@/ui/svg/sprite.svg#icon-dark-mode"></use>
            </svg>
        </div>

        <div class="sidebar__item" @click="toggleTheme">
            <svg class="sidebar__svg ">
                <use xlink:href="@/ui/svg/sprite.svg#icon-chess"></use>
            </svg>
        </div>
        <div class="sidebar__item" @click="goToView('gameCreateView')">
            <svg class="sidebar__svg ">
                <use xlink:href="@/ui/svg/sprite.svg#icon-rock"></use>
            </svg>
        </div>

  </div>
</template>

<!-- --------------------------------------------------------------- -->

<script setup lang="ts">
import { useDashboardStore } from '@/stores/dashboardStore';
import { useRouter } from 'vue-router';


// - Router ------------------------------------------------------------

const router = useRouter();


// - Store -------------------------------------------------------------

const dashboardStore = useDashboardStore();


// - Methods -----------------------------------------------------------
function toggleMenu() {
    dashboardStore.toggleMenuState();
}

function toggleTheme() {
    dashboardStore.toggleTheme();
}

function goToView(view: string) {
    router.push({ name: view })
}


</script>

<!-- --------------------------------------------------------------- -->

<style lang="scss" scoped>
// Placeholder comment to ensure global styles are imported correctly


.sidebar {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding-top: .5rem;

    &__item {
        padding: 5px;
        margin-top: 0;
        cursor: pointer;
    }

    &__svg {
        width: 1.5rem;
        height: 1.5rem;
        fill: var(--color-slate-200);
        color:  var(--color-slate-200);
    }

}
</style>