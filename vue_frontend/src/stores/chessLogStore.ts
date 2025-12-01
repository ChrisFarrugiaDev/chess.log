import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import axios from '@/axios';
import type { Game } from '@/types/game.type';
import type { Move } from '@/types/move.type';
import { useAppStore } from './appStore';





export const useChessLogStore = defineStore('chessLogStore', () => {

    const appStore = useAppStore();

    // ---- State ------------------------------------------------------
    // Collections list
    const collections = ref<Record<string, string | number>[]>([
        { id: 1, name: "Réti Opening",         sort_order: 1, color: "#ef4444" }, // red-500
        { id: 2, name: "Queen's Gambit",       sort_order: 2, color: "#f97316" }, // orange-500
        { id: 3, name: "Sicilian Defence",     sort_order: 3, color: "#f59e0b" }, // amber-500
        { id: 4, name: "Nimzo-Indian Defence", sort_order: 4, color: "#22c55e" }, // green-500
        { id: 5, name: "My Games",             sort_order: 5, color: "#06b6d4" }, // cyan-500
        { id: 6, name: "Tactical Patterns",    sort_order: 6, color: "#3b82f6" }, // blue-500
        { id: 7, name: "Endgames",             sort_order: 7, color: "#8b5cf6" }, // violet-500
    ]);
    // example: [{ id: 1, name: "Reti", sort_order: 1 }, ...]

    // Games list (for the currently loaded collections)
    const games = ref<Record<string, Record<string, string | number>[]>>({
        1: [{ id: 1, collection_id: 1, name: "Reti Study 1" }, { id: 102, collection_id: 1, name: "Reti vs 1700 Rapid" }, { id: 103, collection_id: 1, name: "Reti Idea: Qa4+" },],
        2: [{ id: 2, collection_id: 2, name: "QG Classical Line" }, { id: 202, collection_id: 2, name: "QG Accepted Trap" },],
        3: [{ id: 3, collection_id: 3, name: "Sicilian Najdorf Sample" }, { id: 302, collection_id: 3, name: "Sicilian Anti-Sicilian Line" },],
        4: [{ id: 4, collection_id: 4, name: "Nimzo: Rubinstein Setup" }, { id: 402, collection_id: 4, name: "Nimzo: My 2025 Study" },],
        5: [{ id: 5, collection_id: 5, name: "My Blitz Game #1" }, { id: 502, collection_id: 5, name: "My Rapid Game #2" }, { id: 503, collection_id: 5, name: "My Classical Game #3" },],
        6: [{ id: 6, collection_id: 6, name: "Tactic: Greek Gift" }, { id: 602, collection_id: 6, name: "Tactic: Smothered Mate" },],
        7: [{ id: 7, collection_id: 7, name: "Basic Pawn Ending" }, { id: 702, collection_id: 7, name: "Rook Endgame Lücena" },],
    });
    // example: [{ id: 10, collection_id: 1, name: "Reti Study 1" }, ...]

    // Moves, lazy-loaded per game → stored as { [gameId]: Move[] }
    const moves = ref({});
    // example: moves.value[10] = [{ move_number, color, san, fen }]


    const activeCollection = ref<null | number>(2);
    const activeGame = ref<null | number>(null);


    // ---- Getters ----------------------------------------------------
    const getCollections = computed(() => {
        return collections.value
    });

    const getGames = computed(() => {
        return (collectionId: number | string) => {
            return games.value[collectionId] ?? [];
        };
    });


    const getActiveCollection = computed(() => activeCollection.value);
    const getActiveGame = computed(() => activeGame.value);


    // ---- Actions ----------------------------------------------------

    function setActiveCollection (id: number) {
        activeCollection.value = id;
    }

    function setActiveGame (id: number) {
        activeGame.value = id;
    }

    async function saveGame(game: Game, moves: Move[]) {
        try {           
            
            const url = `${appStore.getAppUrl}/api/games`;

            const r = await axios.post(url, {game, moves});

            console.log(r);

        } catch (err) {
            console.log("! chessLogStore.ts saveGame ! \n", err);
        }
    }

    // - Expose --------------------------------------------------------
    return {
        getCollections,
        getGames,

        getActiveCollection,
        getActiveGame,

        setActiveCollection,
        setActiveGame,
        saveGame,
    };
});
