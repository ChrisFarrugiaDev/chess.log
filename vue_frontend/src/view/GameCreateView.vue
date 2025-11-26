<template>
    <div class="display">


        <section class="chess">

            <div class="buttons">
                <button class="buttons__btn" @click="toggleOrientation"> Toggle </button>
                <button  class="buttons__btn" @click="resetBoard">Reset</button>
                <button  class="buttons__btn" @click="undoMove">Undo</button>
            </div>
            <TheChessboard class="the-chess-board"  @board-created="onBoardCreated" @move="handleMove"/>

        </section>

            <section class="game-form">
            <div class="game-form__field">
                <label class="game-form__label">Collection</label>
                <VueSelect
                    v-model="gameState.collection_id"
                    :style="[vueSelectStyles]"
                    :options="collectionOptions"
                    label-by="label"
                    value-by="value"
                    placeholder="Select collection"
                    class="game-form__select"
                />
            </div>

            <div class="game-form__field">
                <label class="game-form__label">Game name</label>
                <input
                    v-model="gameState.name"
                    type="text"
                    class="game-form__input"
                    placeholder="e.g. Reti – Bg4 Ne5 idea"
                />
            </div>

            <div class="game-form__field">
                <label class="game-form__label">Date</label>
                <input
                    v-model="gameState.date"
                    type="date"
                    class="game-form__input"
                />
            </div>

            <div class="game-form__field game-form__field--full">
                <label class="game-form__label">Notes</label>
                <textarea
                    v-model="gameState.notes"
                    rows="3"
                    class="game-form__textarea"
                    placeholder="Ideas, plans, why this game is important..."
                ></textarea>
            </div>
        </section>

    </div>
</template>

<script setup lang="ts">
import { useVueSelectStyles } from '@/composables/useVueSelectStyles';
import { onMounted, ref } from 'vue';
import { TheChessboard, type BoardApi, type MoveEvent } from 'vue3-chessboard';
import 'vue3-chessboard/style.css';

import VueSelect from "vue3-select-component";


// - State -------------------------------------------------------------

const vueSelectStyles = useVueSelectStyles();

// Internal board API reference
const boardApi = ref<BoardApi | null>(null);

// Local game state used when creating a new game
const gameState = ref({
    id: 0,
    collection_id: 0,
    name: "",    
    orientation: "white",
    notes: "",
    date: "",
    moves: [] as MoveEvent[],
});

const collectionOptions = [
    { value: 1, label: 'Réti Opening' },
    { value: 2, label: "Queen's Gambit" },
    { value: 3, label: 'Sicilian Defence' },
    { value: 4, label: 'Nimzo-Indian Defence' },
    { value: 5, label: 'My Games' },
];

// -Methods ------------------------------------------------------------

// Called by vue3-chessboard when ready
function onBoardCreated(api: BoardApi) {
    boardApi.value = api; 
    api.setPosition("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1");
}

function handleMove(move: MoveEvent) {
    gameState.value.moves.push(move); // tracked for saving
}

function resetBoard() {
    gameState.value.moves = [];

    if (!boardApi.value) return;
    boardApi.value.setPosition("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1");
}

function undoMove() {
    if (!boardApi.value) return;

    // if only 1 move → reset everything
    if (gameState.value.moves.length <= 1) {
        resetBoard();
        return;
    }

    // undo last move (pop it)
    const lastMove = gameState.value.moves.pop();

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

<style lang="scss" scoped>
.display {
    display:  flex;
    padding: 2rem;
    height: 100vh;
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
        color: var(--color-slate-700);
        background-color: var(--color-slate-100);
        border: 1px solid var(--color-slate-800);
        border-radius: 3px;
        padding: .7rem 1rem;
        display: flex;
        justify-content: center;
        align-items: center;
        min-width: 5rem;
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
}
</style>
