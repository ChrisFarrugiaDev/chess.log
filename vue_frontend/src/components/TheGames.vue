<template>
    <div class="games">
        <div class="games__item"  :class="{ 'games__item--selected': game.id == getActiveGame }"
            v-for="game in getGames(getActiveCollection ?? 1)" @click="selectGame(Number(game.id))">           
            <span class="games__name">{{ game.name }}</span>
        </div>
    </div>
</template>

<!-- --------------------------------------------------------------- -->

<script setup lang="ts">
import { useChessLogStore } from '@/stores/chessLogStore';
import { storeToRefs } from 'pinia';
import { useRouter } from 'vue-router';


const router = useRouter();



// - Store -------------------------------------------------------------

const chessLogStore = useChessLogStore();
const { getGames, getActiveGame, getActiveCollection } = storeToRefs(chessLogStore);

function selectGame(id:number) {
    chessLogStore.setActiveGame(id);

    // navigate using active collection + selected game
    router.push(`/collection/${getActiveCollection.value}/game/${id}`);

}




</script>

<!-- --------------------------------------------------------------- -->

<style lang="scss" scoped>
.games {
    font-family: var(--font-primary);
    color: var(--color-text-1);
    font-size: .9rem;
    border: none;

    &__item {
        padding: .5rem;
        cursor: pointer;
        text-wrap: nowrap;

        &:hover {
            background-color: var(--color-slate-100);
        }

        &--selected {
            background-color: var(--color-slate-100);
            font-weight: 700;
        }

        display: flex;
    }
}
</style>