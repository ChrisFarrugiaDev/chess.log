<template>
	<div class="game-view">

		<!-- Game Header -->
		<div class="game-view__header" v-if="game">
			<h1 class="game-view__title">{{ game.name }}</h1>
			<p class="game-view__subtitle">
				{{ moves.length }} moves — {{ game.orientation === 'white' ? 'White' : 'Black' }} to start
			</p>

			<button class="game-view__toggle-btn vbtn vbtn--sky" @click="isSingleView = !isSingleView">
				{{ isSingleView ? "Show All Positions" : "Show Step-by-Step" }}
			</button>
		</div>

		<template v-if="game && moves">


			<div class="game-view__one_item" v-show="isSingleView">
				<MyChessBoard :moves="moves" :orientation="game?.orientation!" :reset="Number(props.gameId)"
					@updatePostion="currentPostion = $event" />

				<p class="game-view__move-label">
					Move {{ moves[currentPostion]!.move_number! + 1 }} — <strong>{{ moves[currentPostion]!.san
						}}</strong>
				</p>
			</div>


			<!-- Grid of positions -->
			<div class="game-view__grid" v-show="!isSingleView">

				<div v-for="(m, index) in moves" :key="index" class="game-view__item">

					<MyChessBoardPosition :after="m.after!" :before="m.before!" :from_square="m.from!"
						:to_square="m.to!" :orientation="game?.orientation!" class="game-view__board" />

					<p class="game-view__move-label">
						Move {{ m.move_number! + 1 }} — <strong>{{ m.san }}</strong>
					</p>
				</div>
			</div>
		</template>


	</div>
</template>

<!-- --------------------------------------------------------------- -->

<script setup lang="ts">
import { useChessLogStore } from '@/stores/chessLogStore';
import { onMounted, ref, watch } from 'vue';
import type { Move } from '@/types/move.type';
import MyChessBoardPosition from '@/components/MyChessBoardPosition .vue';
import MyChessBoard from '@/components/MyChessBoard.vue';
import type { Game } from '@/types/game.type';


const chessLogStore = useChessLogStore();
const isSingleView = ref(true);

const props = defineProps<{
	gameId: string,
	collectionId: string
}>();

const moves = ref<Move[]>([]);
const game = ref<Game | null>(null);


const currentPostion = ref(0);

async function loadGame() {
	const gameID = props.gameId; 
	const mm = await chessLogStore.getGameMoves(gameID);
	moves.value = mm;

	const gg = await chessLogStore.getGame(props.collectionId, gameID);
	if (gg) game.value = gg;
}

watch(
	() => props.gameId,
	async () => {
		loadGame();	
	},
	{ immediate: true }
);

onMounted(async () => {
	if (props.collectionId) {
		chessLogStore.setActiveCollection(Number(props.collectionId));
	}
	if (props.gameId) {
		chessLogStore.setActiveGame(Number(props.gameId));
	}
	loadGame();

})


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

	&__toggle-btn {
		width: 12rem;
		margin: 1rem auto;
	}

	/* Grid of boards */
	&__grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
		gap: 1.5rem;
	}

	&__item,
	&__one_item {
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

	&__one_item {

		width: fit-content;
		margin: 0 auto;
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
