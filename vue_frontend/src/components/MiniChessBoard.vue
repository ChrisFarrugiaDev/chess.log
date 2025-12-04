<template>
    <div ref="boardEl" class="mini-board"></div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from "vue";
import { Chessground } from "chessground";
import type { Api } from "chessground/api";
// import "chessground/dist/chessground.css";

// Props ---------------------------------------------------------------

const props = defineProps<{
    after: string;        // FEN after the move
    before: string;       // optional, unused for now
    from_square: string;  // "e2"
    to_square: string;    // "e4"
}>();

// State ---------------------------------------------------------------

const boardEl = ref<HTMLDivElement | null>(null);
let cg: Api | null = null;

// Helpers -------------------------------------------------------------

function drawArrow() {
    if (!cg) return;
    if (!props.from_square || !props.to_square) return;

    cg.setAutoShapes([
        {
            orig: props.from_square as any,
            dest: props.to_square as any,
            brush: "green"
        }
    ]);
}

// Lifecycle ------------------------------------------------------------

onMounted(() => {
    if (!boardEl.value) return;

    cg = Chessground(boardEl.value, {
        fen: props.after,
        coordinates: false,

        movable: {
            free: false,
            color: undefined   // disable moves properly
        },

        draggable: {
            enabled: false
        },

        highlight: {
            lastMove: false,
            check: false
        },

        animation: {
            enabled: true
        }
    });


    drawArrow();
});

// If props change reactively → update board
watch(
    () => props.after,
    fen => {
        if (cg) cg.set({ fen });
        drawArrow();
    }
);

watch([() => props.from_square, () => props.to_square], () => {
    drawArrow();
});
</script>

<style scoped>
.mini-board {
    width: 16rem;
    height: 16rem;
    border-radius: 6px;
    overflow: hidden;
}
</style>
