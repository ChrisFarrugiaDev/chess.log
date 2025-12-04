<template>
    <div class="games">
        <div class="games__item" v-for="game in getGames(getActiveCollection ?? 1)" :key="game.id"
            :class="{ 'games__item--selected': game.id == getActiveGame }"
            @click="selectGame($event, Number(game.id))">
            <!-- Name / edit field -->
            <div class="games__main">
                <template v-if="editingGameId === Number(game.id)">
                    <input v-model="editingName" class="games__input" type="text" @keyup.enter="saveEdit(game)"
                        @keyup.esc="cancelEditDelete" @blur="saveEdit(game)" autofocus />
                </template>

                <template v-else>
                    <span class="games__name">
                        {{ game.name }}
                    </span>
                </template>
            </div>


            <!-- Actions -->
            <div class="games__actions" v-if="getGameId == game.id">
                <!-- Edit / Save / Cancel -->
                <template v-if="editingGameId === Number(game.id) || deletingGameId === Number(game.id)">
                    <svg class="games__btn games__btn--ok" @click.stop="editOrDelete(game)">
                        <use xlink:href="@/ui/svg/sprite.svg#icon-very-good"></use>
                    </svg>
                    <svg class="games__btn games__btn--cancel" @click.stop="cancelEditDelete">
                        <use xlink:href="@/ui/svg/sprite.svg#icon-close-bold"></use>
                    </svg>
                </template>
                <template v-else>
                    <svg class="games__btn games__btn--edit" @click.stop="startEdit(game)">
                        <use xlink:href="@/ui/svg/sprite.svg#icon-edit"></use>
                    </svg>

                    <svg class="games__btn games__btn--delete" @click.stop="startDelete(game)">
                        <use xlink:href="@/ui/svg/sprite.svg#icon-delete"></use>
                    </svg>

                </template>


            </div>
        </div>
    </div>
</template>

<!-- ----------------------------------------------------------------------- -->

<script setup lang="ts">
import { computed, ref, watch } from "vue";
import { useChessLogStore } from "@/stores/chessLogStore";
import { storeToRefs } from "pinia";
import { useRoute, useRouter } from "vue-router";

const router = useRouter();
const route = useRoute();




// - Store -------------------------------------------------------------

const chessLogStore = useChessLogStore();
const { getGames, getActiveGame, getActiveCollection } = storeToRefs(chessLogStore);

// - State -------------------------------------------------------------
const deletingGameId = ref<number | null>(null);
const editingGameId = ref<number | null>(null);
const editingName = ref("");

function selectGame(e: Event, id: number) {
    // console.log(e.target)
    chessLogStore.setActiveGame(id);
    router.push(`/collection/${getActiveCollection.value}/game/${id}`);
}

// - Computed ----------------------------------------------------------

const getGameId = computed(() => {
    return route.params.gameId;
});


// - Wachers -----------------------------------------------------------

watch(
    () => route.params.gameId,
    (newGameId, oldGameId) => {
        cancelEditDelete()
    }
);

// - Methods -----------------------------------------------------------

function startEdit(game: Record<string, string | number>) {

    // if (props.gameId != game.id) { return }

    deletingGameId.value = null;
    editingGameId.value = Number(game.id);
    editingName.value = String(game.name ?? "");
}

function startDelete(game: Record<string, string | number>) {
    // if (props.gameId != game.id) { return }

    deletingGameId.value = Number(game.id);
    editingGameId.value = null;
    editingName.value = "";
}


function saveEdit(game: Record<string, string | number>) {
    if (editingGameId.value !== Number(game.id)) return;

    const newName = editingName.value.trim();
    if (!newName) {
        // empty → just cancel
        cancelEditDelete();
        return;
    }

    console.log("SAVE GAME NAME (not sent to backend yet):", {
        id: Number(game.id),
        newName,
    });

    // update locally so UI reflects the change
    (game as any).name = newName;

    editingGameId.value = null;
    editingName.value = "";
}


function deleteGame(game: Record<string, string | number>) {
    console.log("DELETE GAME (not sent to backend yet):", {
        id: Number(game.id),
        name: game.name,
    });

    const collectionId = route.params.collectionId
    
    if (!collectionId) return

    const games = chessLogStore.getGames(collectionId as string)

    const index = games.findIndex(gg => gg.id == getGameId.value)

    games.splice(index, 1);

    if (getGameId.value == game.id) {
        router.push(`/`);
    }

    deletingGameId.value = null;
}

function editOrDelete (game: Record<string, string | number>) {
    if ( deletingGameId.value) {
        deleteGame(game);
    }

    if (editingGameId.value) {
        saveEdit(game);
    }
}

function cancelEditDelete() {
    deletingGameId.value = null;
    editingGameId.value = null;
    editingName.value = "";
}
</script>

<!-- ----------------------------------------------------------------------- -->

<style lang="scss" scoped>
.games {
    font-family: var(--font-primary);
    color: var(--color-text-1);
    font-size: 0.9rem;


    &__item {
        padding: 0.5rem;
        cursor: pointer;
        display: grid;
        grid-template-columns: 1fr 3rem;
        align-items: center;
        justify-content: space-between;
        transition: background-color 0.2s;
        height: 3rem;
        gap: 5px;

        &:hover {
            background-color: var(--color-slate-50);

            .games__actions {
                opacity: 1;
                pointer-events: auto;
            }
        }

        &--selected {
            background-color: var(--color-slate-100);
            font-weight: 700;
        }
    }

    &__main {
        flex: 1;
        max-width: 10rem;
    }

    &__name {
        text-wrap: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        padding: 2px 2px 2px 5px;
    }

    &__input {
        width: 100%;
        font: inherit;
        padding: 2px 2px 2px 4px;
        border-radius: 4px;
        border: 1px solid var(--color-slate-300);
        background-color: var(--color-slate-50);
        color: var(--color-slate-700);

        &:focus {
            outline: none;
            border-color: var(--color-slate-500);
        }
    }

    &__actions {
        display: flex;
        gap: .3rem;
        opacity: 0;
        pointer-events: none;
        transition: .2s;
        


        z-index: 100;
        height: 1.2rem;
        display: flex;
        align-items: center;
        // padding: 0px 5px;
        gap: 5px;
    }

    &__btn {        
        cursor: pointer;
        width: 1.4rem;
        height: 1.4rem;
        padding: 3px;
        background-color:  var(--color-slate-100);
      
        color: var(--color-slate-700);
        fill: currentColor;

        border-radius: 5px;
        border: 1px solid currentColor;

        &:hover {
            color: var(--color-rose-500);
        }
    }
}
</style>
