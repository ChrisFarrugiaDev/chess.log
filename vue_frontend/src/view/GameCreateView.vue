<template>
    <div class="display">


        <section class="chess">

            <div class="buttons">
                <button class="buttons__btn  vbtn--sky" @click="toggleOrientation"> Toggle </button>
                <button class="buttons__btn vbtn--sky" @click="resetBoard">Reset</button>
                <button class="buttons__btn vbtn--sky" @click="undoMove">Undo</button>
            </div>

            <TheChessboard class="the-chess-board" @board-created="onBoardCreated" @move="handleMove" />

        </section>

        <section class="game-form">

            <!-- Error -->
            <p v-if="errorMessage" class="game-form__error">{{ errorMessage }}</p>


            <div class="game-form__field" v-if="!gameState.create_collection">
                <label class="game-form__label" for="select-collection">Collection</label>
                <VueSelect v-model="gameState.collection_id" id="select-collection" :style="[vueSelectStyles]"
                    :options="collectionOptions" label-by="label" value-by="value" placeholder="Select collection"
                    class="game-form__select" />
            </div>

            <div class="game-form__field" v-if="gameState.create_collection">
                <label class="game-form__label" for="collection_name">Collection Name</label>
                <input v-model="gameState.collection_name" id="collection_name" type="text" class="game-form__input"
                    placeholder="Enter collection name…" />
            </div>





            <div class="game-form__field">
                <label class="game-form__label cursor-pointer">
                    <input type="checkbox" v-model="gameState.create_collection" />
                    Create new collection
                </label>
            </div>

            <div class="game-form__field">
                <label class="game-form__label" for="name">Game name</label>
                <input v-model="gameState.name" id="name" type="text" class="game-form__input"
                    placeholder="e.g. Reti - Bg4 Ne5 idea" />
            </div>


            <div class="game-form__field game-form__field--full">
                <label class="game-form__label" for="notes">Notes</label>
                <textarea v-model="gameState.notes" id="notes" rows="3" class="game-form__textarea"
                    placeholder="Ideas, plans, why this game is important..."></textarea>
            </div>

            <div class=" mt-9 flex">
                <button v-if="!confirmOn" class="game-form__btn vbtn--sky mt-3" @click.prevent="initSaveGame">Save
                    Game</button>
                <button v-if="confirmOn" class="game-form__btn vbtn--slate mt-3"
                    @click.prevent="confirmOn = false">Cancel</button>
                <button v-if="confirmOn" class="game-form__btn vbtn--sky mt-3"
                    @click.prevent="saveGame">Confirm</button>
            </div>
        </section>



    </div>
</template>

<!-- --------------------------------------------------------------- -->

<script setup lang="ts">
import { useVueSelectStyles } from '@/composables/useVueSelectStyles';
import type { Option } from "vue3-select-component";
import { useChessLogStore } from '@/stores/chessLogStore';
import type { Game } from '@/types/game.type';
import type { Move } from '@/types/move.type';
import { onMounted, ref, watch } from 'vue';
import { TheChessboard, type BoardApi, type MoveEvent } from 'vue3-chessboard';
import 'vue3-chessboard/style.css';

import VueSelect from "vue3-select-component";
import { useApiError } from '@/composables/useApiError';
import { useDashboardStore } from '@/stores/dashboardStore';

// - Composable --------------------------------------------------------
const { errorMessage, handleApiError } = useApiError();
const vueSelectStyles = useVueSelectStyles();

// - Store -------------------------------------------------------------

const chessLogStore = useChessLogStore();
const dashboardStore = useDashboardStore();

// - State -------------------------------------------------------------


// Internal board API reference
const boardApi = ref<BoardApi | null>(null);

// Local game state used when creating a new game
const gameState = ref({
    id: null,
    collection_id: null as null | number,
    collection_name: "",

    create_collection: false,
    name: "",
    orientation: "white",
    notes: null,
});



const movesState = ref<MoveEvent[]>([]);

const collectionOptions = ref<Option<number>[]>([]);

const confirmOn = ref(false);

// - Watch -------------------------------------------------------------
watch(
    () => chessLogStore.getCollections,
    (val: any[]) => {
        if (val && val.length > 0) {
            collectionOptions.value = val.map<Option<number>>((c) => ({
                value: c.id,
                label: c.name,
            }));
        } else {
            collectionOptions.value = [];
        }
    },
    { immediate: true }
);


// - Methods -----------------------------------------------------------


function initSaveGame() {
    if (
        (!gameState.value.create_collection && !gameState.value.collection_id) ||
        (gameState.value.create_collection && !gameState.value.name) ||
        movesState.value.length < 3
    ) {
        return
    }
    confirmOn.value = true;
}

async function saveGame() {
    errorMessage.value = "";

    const moves: Move[] = [];

    const game: Game = {
        create_collection: gameState.value.create_collection,  
        collection_name: gameState.value.collection_name ?? "",
        collection_id: gameState.value.collection_id ?? 0,
        name: gameState.value.name!,
        orientation: gameState.value.orientation,
        notes: gameState.value.notes ?? "",
    }



    movesState.value.forEach((m, i) => {
        const movie: Move = {
            move_number: i,
            color: m.color,
            before: m.before,
            after: m.after,
            san: m.san,
            lan: m.san,
            piece: m.piece,
            from: m.from,
            to: m.to,
            flags: m.flags,
        }

        moves.push(movie);
    });


    try {
        dashboardStore.setIsLoading(true);
        await chessLogStore.saveGame(game, moves);

        await chessLogStore.fetchProfile();


        gameState.value.id = null;
        gameState.value.collection_id = null;
        gameState.value.collection_name = ""; 
        gameState.value.create_collection = false;
        gameState.value.name = "";
        gameState.value.orientation = "white";
        gameState.value.notes = null;

        resetBoard();

    } catch (err) {
        handleApiError(err);
    } finally {
        setTimeout(() => { dashboardStore.setIsLoading(false) }, 1000);
    }
}

// Called by vue3-chessboard when ready
function onBoardCreated(api: BoardApi) {
    boardApi.value = api;
    api.setPosition("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1");
}

function handleMove(move: MoveEvent) {
    movesState.value.push(move); // tracked for saving
    console.log(move);
}

function resetBoard() {
    movesState.value = [];

    if (!boardApi.value) return;
    boardApi.value.setPosition("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1");
}

function undoMove() {
    if (!boardApi.value) return;

    // if only 1 move → reset everything
    if (movesState.value.length <= 1) {
        resetBoard();
        return;
    }

    // undo last move (pop it)
    const lastMove = movesState.value.pop();

    // set board to previous FEN
    boardApi.value.setPosition(lastMove!.before);
}

function toggleOrientation() {
    if (!boardApi.value) return;
    boardApi.value.toggleOrientation();

    gameState.value.orientation =
        gameState.value.orientation === "white" ? "black" : "white";
}

</script>

<!-- --------------------------------------------------------------- -->

<style lang="scss" scoped>
.display {
    display: flex;
    padding: 2rem;
    max-height: 100vh;
    overflow: scroll;

    @media (max-width: 1100px) {
        flex-direction: column;
    }

}

.chess {
    flex: 1;
}

.the-chess-board {
    width: 100%;
    padding: 1rem;
    opacity: 0.9;

}

.buttons {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 1rem;
    margin: 1rem;


    &__btn {
        cursor: pointer;
        font-family: var(--font-primary);
        font-size: 1rem;
        color: var(--color-slate-100);
        border: 1px solid var(--color-slate-100);
        border-radius: 5px;
        padding: .7rem 1rem;
        display: flex;
        justify-content: center;
        align-items: center;
        min-width: 7rem;
    }
}



.game-form {
    flex: 0 0 40%;
    display: flex;
    flex-direction: column;

    gap: 1rem;
    padding: 1rem 1rem 0 1rem;
    font-family: var(--font-primary);
    font-size: 0.9rem;
    color: var(--color-slate-800);

    &__field {
        display: flex;
        flex-direction: column;
        gap: 0.3rem;
    }


    &__label {
        font-family: var(--font-primary);
        font-size: 1rem;
        color: var(--color-slate-700);
        font-display: flex;
        align-items: center;

        input {
            transform: translateY(1px);
        }
    }

    &__input,
    &__textarea {
        padding: 0.4rem 0.6rem;
    }

    &__input,
    &__select {
        height: 3rem;
    }

    &__input,
    &__textarea,
    &__select {
        border-radius: 4px;
        border: 1px solid var(--color-slate-300);
        color: var(--color-slate-800);

        font-family: inherit;
        font-size: 0.9rem;
        background-color: var(--color-slate-50);
        outline: none;

        &:focus {
            border-color: var(--color-slate-500);
        }
    }

    &__textarea {
        resize: vertical;
        min-height: 3rem;
    }

    &__btn {
        cursor: pointer;
        font-family: var(--font-primary);
        font-size: 1rem;
        color: var(--color-slate-100);
        border: 1px solid var(--color-slate-100);
        border-radius: 5px;
        padding: .7rem 1rem;
        display: flex;
        justify-content: center;
        align-items: center;
        flex: 1;
        color: var(--color-slate-100);
        border: 1px solid var(--color-slate-100);
        padding: .7rem 1rem;
        font-family: var(--font-primary);
        font-size: 1rem;
        font-weight: 400;
    }

    &__error {
        color: var(--color-red-500);
        text-align: center;
        font-size: .85rem;
        margin-top: -0.5rem;
    }

    &__message {
        color: var(--color-green-500);
        text-align: center;
        font-size: .85rem;
        margin-top: -0.5rem;
    }
}

.flex {
    display: flex;
    gap: 1rem;
}
</style>
