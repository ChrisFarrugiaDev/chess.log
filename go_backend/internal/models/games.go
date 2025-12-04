package models

import (
	"context"

	"time"

	up "github.com/upper/db/v4"
)

type Game struct {
	ID           int64     `db:"id,omitempty" json:"id"`
	CollectionID int64     `db:"collection_id" json:"collection_id"`
	UserID       int64     `db:"user_id" json:"user_id"`
	Name         string    `db:"name" json:"name"`
	Orientation  string    `db:"orientation" json:"orientation"`
	SortOrder    int       `db:"sort_order" json:"sort_order"`
	Notes        *string   `db:"notes" json:"notes"` // nullable
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

// ---------------------------------------------------------------------

func (m *Game) TableName() string {
	return "games"
}

// ---------------------------------------------------------------------

func (m *Game) Create() (*Game, error) {

	now := time.Now().UTC()
	m.CreatedAt = now
	m.UpdatedAt = now

	_collection := upperSession.Collection(m.TableName())

	err := _collection.InsertReturning(m)

	if err != nil {

		return nil, err
	}

	return m, nil
}

// ---------------------------------------------------------------------

func CreateGameWithMoves(ctx context.Context, game *Game, moves []*GameMove) (*Game, error) {

	err := upperSession.Tx(func(tx up.Session) error {

		now := time.Now().UTC()

		// -------------------------------------------
		// 1. Insert game
		// -------------------------------------------
		game.CreatedAt = now
		game.UpdatedAt = now

		gameCol := tx.Collection(game.TableName())

		if err := gameCol.InsertReturning(game); err != nil {
			return err // rollback
		}

		// -------------------------------------------
		// 2. Insert moves one-by-one (safe + correct)
		// -------------------------------------------
		moveCol := tx.Collection((&GameMove{}).TableName())

		for _, mv := range moves {
			mv.GameID = game.ID
			mv.CreatedAt = now
			mv.UpdatedAt = now

			// Use InsertReturning if you want ID, Insert if not needed
			if err := moveCol.InsertReturning(mv); err != nil {
				return err // rollback
			}
		}

		// commit tx
		return nil
	})

	if err != nil {
		return nil, err
	}

	return game, nil
}

func (m *Game) GetByUserID(userID int64) ([]*Game, error) {

	var games []*Game

	_collection := upperSession.Collection(m.TableName())

	err := _collection.Find(up.Cond{"user_id": userID}).OrderBy("name").All(&games)

	if err != nil {
		return nil, err
	}

	return games, nil
}
