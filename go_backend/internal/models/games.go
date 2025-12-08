package models

import (
	"context"
	"fmt"

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

func CreateGameWithMoves(

	ctx context.Context,
	game *Game,
	moves []*GameMove,
	createCollection bool,
	collectionName string,

) (*Game, error) {

	err := upperSession.Tx(func(tx up.Session) error {

		now := time.Now().UTC()

		// ---------------------------------------------------
		// 0. Create collection if requested
		// ---------------------------------------------------
		if createCollection {
			newCollection := &Collection{
				UserID:    game.UserID,
				Name:      collectionName,
				SortOrder: 0, // optional: auto-assign later
				CreatedAt: now,
				UpdatedAt: now,
			}

			col := tx.Collection((&Collection{}).TableName())
			if err := col.InsertReturning(newCollection); err != nil {
				return err // rollback
			}

			// assign to game
			game.CollectionID = newCollection.ID
		}

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

// ---------------------------------------------------------------------

func RenameGame(gameID int64, userID int64, newName string) error {
	col := upperSession.Collection((&Game{}).TableName())

	now := time.Now().UTC()

	// Only allow renaming if the game belongs to the user
	res := col.Find(up.Cond{
		"id":      gameID,
		"user_id": userID,
	})

	// Ensure game exists
	count, err := res.Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("game not found or unauthorized")
	}

	// Update
	return res.Update(map[string]any{
		"name":       newName,
		"updated_at": now,
	})
}

// ---------------------------------------------------------------------

func DeleteGame(gameID int64, userID int64) error {

	gameCol := upperSession.Collection((&Game{}).TableName())

	// Ensure ownership
	res := gameCol.Find(up.Cond{
		"id":      gameID,
		"user_id": userID,
	})

	count, err := res.Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("game not found or unauthorized")
	}

	// Delete the game — Postgres deletes its moves automatically
	return res.Delete()
}

// ---------------------------------------------------------------------
