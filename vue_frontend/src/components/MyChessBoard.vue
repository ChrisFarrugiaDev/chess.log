<template>
    <div>
        <TheChessboard class="the-chess-board" :drag-disabled="true" @board-created="onBoardCreated" />

        <div class="btns">
            <button class="vbtn vbtn--sky" @click="prev">prev</button>
            <button class="vbtn vbtn--sky" @click="next">next</button>
        </div>
    </div>
</template>

<!-- --------------------------------------------------------------- -->

<script setup lang="ts">
import type { Move } from '@/types/move.type';
import { ref, watch } from 'vue';
import { TheChessboard, type BoardApi, type MoveEvent } from 'vue3-chessboard';
import 'vue3-chessboard/style.css';

// - Props -------------------------------------------------------------

const props = defineProps<{
    moves: Move[],
    orientation: string,
    reset: number,

    // default: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
}>()


const emits = defineEmits<{
    (e: 'updatePostion', tab: number): void;
}>()

// - State -------------------------------------------------------------



// Internal board API reference
const boardApi = ref<BoardApi | null>(null);

const currentPostion = ref(0);


watch(() => props.reset, () => {
    currentPostion.value = 0;

    const move = props.moves[currentPostion.value];
    if (!move || !move.after) return;
    boardApi.value!.setPosition(move.after);
})

// - Methods -----------------------------------------------------------

function onBoardCreated(api: BoardApi) {
    boardApi.value = api;


    api.setPosition(props.moves[currentPostion.value]!.after!);

    if (props.orientation == "black") {
        boardApi.value!.toggleOrientation();
    }
}

function next() {
    if ((currentPostion.value + 1) < props.moves.length) {
        currentPostion.value++

        const move = props.moves[currentPostion.value];
        if (!move || !move.after) return;
        boardApi.value!.setPosition(move.after);

        emits("updatePostion", currentPostion.value);
    }


}

function prev() {
    if (currentPostion.value > 0) {
        currentPostion.value--

        const move = props.moves[currentPostion.value];
        if (!move || !move.after) return;
        boardApi.value!.setPosition(move.after);

        emits("updatePostion", currentPostion.value);
    }
}
</script>

<!-- --------------------------------------------------------------- -->

<style lang="scss" scoped>
// Placeholder comment to ensure global styles are imported correctly
.the-chess-board {
    width: 30rem;
    height: 30rem;
    padding: 1rem;
    opacity: 1;

}


.btns {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 1rem;

    &>* {
        width: 5rem;
    }
}
</style>