<template>
	<div class="game-view">

		<!-- Game Header -->
		<div class="game-view__header" v-if="game">
			<h1 class="game-view__title">{{ game.name }}</h1>
			<p class="game-view__subtitle">
				{{ moves.length }} moves — {{ game.orientation === 'white' ? 'White' : 'Black' }} to start
			</p>
		</div>

		<!-- Grid of positions -->
		<div class="game-view__grid">
			<div v-for="(m, index) in moves" :key="index" class="game-view__item">


				<MyChessBoardPosition :after="m.after!" :before="m.before!" :from_square="m.from!" :to_square="m.to!"
					class="game-view__board" />

				<p class="game-view__move-label">
					Move {{ m.move_number! + 1 }} — <strong>{{ m.san }}</strong>
				</p>
			</div>
		</div>

	</div>
</template>

<!-- --------------------------------------------------------------- -->

<script setup lang="ts">
import { useChessLogStore } from '@/stores/chessLogStore';
import { ref, watch } from 'vue';
import type { Move } from '@/types/move.type';
import MyChessBoardPosition from '@/components/MyChessBoardPosition .vue';
import type { Game } from '@/types/game.type';
import MiniChessBoard from '@/components/MiniChessBoard.vue';

const chessLogStore = useChessLogStore();

const props = defineProps<{
	gameId: string,
	collectionId: string
}>();

const moves = ref<Move[]>([]);
const game = ref<Game | null>(null);

watch(
	() => props.gameId,
	async (gameID) => {
		const mm = await chessLogStore.getGameMoves(gameID);
		moves.value = mm;

		const gg = await chessLogStore.getGame(props.collectionId, gameID);
		if (gg) game.value = gg;
	},
	{ immediate: true }
);
</script>

<!-- --------------------------------------------------------------- -->

<style lang="scss" scoped>
.game-view {
	padding: 2rem;
	max-height: 100vh;
	overflow-y: scroll;
	opacity: .9;

	&__header {
		text-align: center;
		margin-bottom: 2rem;
	}

	&__title {
		font-size: 1.8rem;
		font-weight: 600;
		margin: 0;
		color: var(--color-slate-800);
		font-family: var(--font-primary);
	}

	&__subtitle {
		font-size: 1rem;
		color: var(--color-slate-600);
		margin-top: 0.3rem;
		font-family: var(--font-primary);
	}

	/* Grid of boards */
	&__grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
		gap: 1.5rem;
	}

	&__item {
		background: var(--color-slate-50);
		padding: 1rem;
		border-radius: 8px;
		border: 1px solid var(--color-slate-200);
		box-shadow: 0 2px 5px rgba(0, 0, 0, 0.06);
		display: flex;
		flex-direction: column;
		align-items: center;
		font-family: var(--font-primary);
	}

	&__board {
		width: 100%;
		max-width: 350px;
		border-radius: 6px;
		overflow: hidden;
	}

	&__move-label {
		margin-top: 0.6rem;
		font-size: 0.9rem;
		color: var(--color-slate-700);
	}
}
</style>
