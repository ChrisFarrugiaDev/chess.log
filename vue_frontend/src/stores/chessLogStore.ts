import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import axios from '@/axios';
import type { Game } from '@/types/game.type';
import type { Move } from '@/types/move.type';
import { useAppStore } from './appStore';
import { useDashboardStore } from './dashboardStore';
import { isBigIntLiteral } from 'typescript';



export const useChessLogStore = defineStore('chessLogStore', () => {

    const appStore = useAppStore();
    const dashboardStore = useDashboardStore();

    // ---- State ------------------------------------------------------
    // Collections list
    const collections = ref<Record<string, string | number>[]>([
        // { id: 1, name: "Réti Opening",         sort_order: 1,  }, 
        // { id: 2, name: "Queen's Gambit",       sort_order: 2,  }, 
        // { id: 3, name: "Sicilian Defence",     sort_order: 3,  }, 
        // { id: 4, name: "Nimzo-Indian Defence", sort_order: 4,  }, 
        // { id: 5, name: "My Games",             sort_order: 5,  }, 
        // { id: 6, name: "Tactical Patterns",    sort_order: 6,  }, 
        // { id: 7, name: "Endgames",             sort_order: 7,  }, 
    ]);
    // example: [{ id: 1, name: "Reti", sort_order: 1 }, ...]

    // Games list (for the currently loaded collections)
    const games = ref<Record<number | string, Game[]>>({
        // 1: [{ id: 1, collection_id: 1, name: "Reti Study 1" }, { id: 102, collection_id: 1, name: "Reti vs 1700 Rapid" }, { id: 103, collection_id: 1, name: "Reti Idea: Qa4+" },],
        // 2: [{ id: 2, collection_id: 2, name: "QG Classical Line" }, { id: 202, collection_id: 2, name: "QG Accepted Trap" },],
        // 3: [{ id: 3, collection_id: 3, name: "Sicilian Najdorf Sample" }, { id: 302, collection_id: 3, name: "Sicilian Anti-Sicilian Line" },],
        // 4: [{ id: 4, collection_id: 4, name: "Nimzo: Rubinstein Setup" }, { id: 402, collection_id: 4, name: "Nimzo: My 2025 Study" },],
        // 5: [{ id: 5, collection_id: 5, name: "My Blitz Game #1" }, { id: 502, collection_id: 5, name: "My Rapid Game #2" }, { id: 503, collection_id: 5, name: "My Classical Game #3" },],
        // 6: [{ id: 6, collection_id: 6, name: "Tactic: Greek Gift" }, { id: 602, collection_id: 6, name: "Tactic: Smothered Mate" },],
        // 7: [{ id: 7, collection_id: 7, name: "Basic Pawn Ending" }, { id: 702, collection_id: 7, name: "Rook Endgame Lücena" },],
    });
    // example: [{ id: 10, collection_id: 1, name: "Reti Study 1" }, ...]

    // Moves, lazy-loaded per game → stored as { [gameId]: Move[] }
    const moves = ref<Record<number | string, Move[]>>({});
    // example: moves.value[10] = [{ move_number, color, san, fen }]


    const activeCollection = ref<null | number>(0);
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


    const getGameMoves = computed(() => {
        return async (id: number | bigint | string) => {
            const key = normalizeId(id);

            if (moves.value[key]) return moves.value[key];

            await fetchGameMoves(id);

            return moves.value[key] || [];
        };
    });



     const getGame = computed(() => {
        return async (collectionID: number | bigint | string, gameID: number | bigint | string): Promise<null | Game> => {

            const collectionGames = games.value[normalizeId(collectionID)];

            const gg = collectionGames?.find(g => g.id == gameID)

            if (gg) {
                return gg
            } 

            return null
        }
     });



    // ---- Actions ----------------------------------------------------

    function normalizeId(id: number | bigint | string): number | string {
        return typeof id === "bigint" ? id.toString() : id;
    }

    function setActiveCollection(id: number) {
        activeCollection.value = id;
    }

    function setActiveGame(id: number) {
        activeGame.value = id;
    }

    async function saveGame(game: Game, moves: Move[]) {
        try {

            const url = `${appStore.getAppUrl}/api/games`;

            const r = await axios.post(url, { game, moves });


        } catch (err) {
            console.log("! chessLogStore.ts saveGame ! \n", err);
        }
    }

    async function fetchProfile() {
        try {
            dashboardStore.setIsLoading(true);

            const url = `${appStore.getAppUrl}/api/user/profile`;

            const r = await axios.get(url);

            if (r.statusText == "OK") {

                if (r.data?.collections) {
                    collections.value = r.data.collections;

                    if (r.data?.games) {
                        games.value = r.data.games
                    }
                }
            }

        } catch (err) {
            console.log("! chessLogStore.ts fetchProfile ! \n", err);
        } finally {
            dashboardStore.setIsLoading(false);
        }
    }


    async function fetchGameMoves(gameID: number | bigint | string) {
        const key = normalizeId(gameID);

        if (moves.value[key]) return moves.value[key];

        try {
            dashboardStore.setIsLoading(true);

            const url = `${appStore.getAppUrl}/api/games/moves/${key}`;
            const r = await axios.get(url);

            if (r.statusText == "OK" && r.data?.moves) {
                moves.value[key] = r.data.moves;
            }
        } finally {
            dashboardStore.setIsLoading(false);
        }
    }


    // - Expose --------------------------------------------------------
    return {
        fetchProfile,
        fetchGameMoves,

        getCollections,
        getGames,

        getActiveCollection,
        getActiveGame,

        setActiveCollection,
        setActiveGame,
        saveGame,

        getGameMoves,
         getGame,
    };
});
