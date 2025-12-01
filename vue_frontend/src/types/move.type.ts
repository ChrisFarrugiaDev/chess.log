export type Move = {
    game_id?: string;
    move_number?: number;
    color?: string;
    before?: string;
    after?: string;
    san?: string;
    lan?: string;
    piece?: string;
    from?: string;
    to?: string;
    flags?: string;
}