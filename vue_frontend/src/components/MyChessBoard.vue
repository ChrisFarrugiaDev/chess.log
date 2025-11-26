<template>
    <div>
        <TheChessboard fen="start" @board-created="onBoardCreated" @move="handleMove"/>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { TheChessboard, type BoardApi, type MoveEvent } from 'vue3-chessboard';
import 'vue3-chessboard/style.css';

const boardApi = ref<BoardApi | null>(null);

// Called by vue3-chessboard when ready
function onBoardCreated(api: BoardApi) {
    boardApi.value = api; // ✔ store the API instance

    // Example position
    api.setPosition("rnbqkb1r/ppp1pppp/5n2/8/1PPp4/2N2N2/P2PPPPP/R1BQKB1R b KQkq - 2 4");
}
function handleMove(move: MoveEvent) {
  console.log(move);
}

</script>

<style lang="scss" scoped>
/* Placeholder */
</style>
