<template>
    <div class="collections">
        <div class="collections__item" :class="{ 'collections__item--selected': collection.id == getActiveCollection }"
            v-for="collection in getCollections" @click="selectCollection(Number(collection.id))">
            <span class="collections__color" :style="{ borderLeftColor: collection.color }">&nbsp;</span>
            <span class="collections__name">{{ collection.name }}</span>
        </div>
    </div>
</template>

<!-- --------------------------------------------------------------- -->

<script setup lang="ts">
import { useChessLogStore } from '@/stores/chessLogStore';
import { storeToRefs } from 'pinia';

// - Store -------------------------------------------------------------

const chessLogStore = useChessLogStore();
const { getCollections, getActiveCollection } = storeToRefs(chessLogStore);

// - Computed ----------------------------------------------------------

// - Methods -----------------------------------------------------------

function selectCollection(id: number) {
    chessLogStore.setActiveCollection(id);
}

</script>

<!-- --------------------------------------------------------------- -->

<style lang="scss" scoped>
// Placeholder comment to ensure global styles are imported correctly

.collections {
    font-family: var(--font-primary);
    color: var(--color-text-1);
    font-size: .9rem;

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

    &__color {
        border-left: 5px solid black;
        border-radius: 10px;
        display: inline-block;
        height: 100%;
    }
}
</style>