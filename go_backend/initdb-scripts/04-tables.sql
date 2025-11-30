DROP TABLE IF EXISTS game_moves CASCADE;
DROP TABLE IF EXISTS games CASCADE;
DROP TABLE IF EXISTS collections CASCADE;
DROP TABLE IF EXISTS users CASCADE;

-- ---------------------------------------------------------------------

CREATE TABLE users (
    id            BIGSERIAL PRIMARY KEY,
    name          TEXT NOT NULL,
    email         TEXT UNIQUE NOT NULL,
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    password_hash TEXT NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER trg_users_updated
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE FUNCTION update_timestamp();

-- ---------------------------------------------------------------------

CREATE TABLE collections (
    id          BIGSERIAL PRIMARY KEY,
    user_id     BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name        TEXT NOT NULL,
    sort_order  INT NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER trg_collections_updated
BEFORE UPDATE ON collections
FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE INDEX idx_collections_user_id
ON collections (user_id);

-- ---------------------------------------------------------------------

CREATE TABLE games (
    id              BIGSERIAL PRIMARY KEY,
    collection_id   BIGINT NOT NULL REFERENCES collections(id) ON DELETE CASCADE,
    user_id     BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name            TEXT NOT NULL,
    orientation     TEXT NOT NULL DEFAULT 'white',

    sort_order      INT NOT NULL DEFAULT 0,
    notes           TEXT,                -- optional
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER trg_games_updated
BEFORE UPDATE ON games
FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE INDEX idx_games_collection_id
ON games (collection_id);


-- ---------------------------------------------------------------------

CREATE TABLE game_moves (
    id              BIGSERIAL PRIMARY KEY,
    game_id         BIGINT NOT NULL REFERENCES games(id) ON DELETE CASCADE,

    move_number     INT NOT NULL,
    color           VARCHAR(1) NOT NULL CHECK (color IN ('w','b')),

    before_fen      TEXT NOT NULL,
    after_fen       TEXT NOT NULL,

    san             VARCHAR(32) NOT NULL,
    lan             VARCHAR(32) NOT NULL,
    piece           VARCHAR(8) NOT NULL,
    from_square     VARCHAR(4) NOT NULL,
    to_square       VARCHAR(4) NOT NULL,
    flags           VARCHAR(8) NOT NULL,

    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);




CREATE TRIGGER set_updated_at
BEFORE UPDATE ON game_moves
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE INDEX idx_game_moves_order
ON game_moves (game_id, move_number, color);

-- ---------------------------------------------------------------------