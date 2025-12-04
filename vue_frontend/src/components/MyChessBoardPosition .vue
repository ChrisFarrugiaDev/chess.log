<template>
    <div>
        <TheChessboard class="the-chess-board" :drag-disabled="true" @board-created="onBoardCreated" />
    </div>
</template>

<!-- --------------------------------------------------------------- -->

<script setup lang="ts">
import { ref } from 'vue';
import { TheChessboard, type BoardApi, type MoveEvent } from 'vue3-chessboard';
import 'vue3-chessboard/style.css';

// - Props -------------------------------------------------------------

const props = defineProps<{
    after: string,
    before: string
    from_square: string,
    to_square: string

    // default: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
}>()


// - State -------------------------------------------------------------

// Internal board API reference
const boardApi = ref<BoardApi | null>(null);

// - Methods -----------------------------------------------------------

function onBoardCreated(api: BoardApi) {
    boardApi.value = api;

    // Show the final position of the move
    api.setPosition(props.after);

    if (props.from_square && props.to_square) {
        // Draw a single green arrow from from_square → to_square


        setTimeout(() => {

            api.setShapes([
                {
                    orig: props.from_square as any,
                    dest: props.to_square as any,
                    brush: 'green',
                },
            ]);
        }, 100)
    }
}
</script>

<!-- --------------------------------------------------------------- -->

<style lang="scss" scoped>
// Placeholder comment to ensure global styles are imported correctly
.the-chess-board {
    width: 20rem;
    height: 20rem;
    padding: 1rem;
    opacity: 1;

}
</style>