package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	up "github.com/upper/db/v4"
)

var ErrDuplicateCollection = errors.New("collection with this name already exists")

type Collection struct {
	ID        int64     `db:"id,omitempty" json:"id"`
	UserID    int64     `db:"user_id" json:"user_id"`
	Name      string    `db:"name" json:"name"`
	SortOrder int       `db:"sort_order,omitempty" json:"sort_order"`
	CreatedAt time.Time `db:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at,omitempty" json:"updated_at"`
}

func (m *Collection) TableName() string {
	return "collections"
}

func (m *Collection) Create() (*Collection, error) {

	now := time.Now().UTC()
	m.CreatedAt = now
	m.UpdatedAt = now

	_collection := upperSession.Collection(m.TableName())

	err := _collection.InsertReturning(m)

	if err != nil {
		// Check for duplicate email (SQLSTATE 23505 = unique violation)
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return nil, ErrDuplicateCollection
		}

		// Wrap and return other errors
		return nil, err
	}

	return m, nil
}

func (m *Collection) GetByUserID(userID int64) ([]*Collection, error) {

	var collections []*Collection

	_collection := upperSession.Collection(m.TableName())

	err := _collection.Find(up.Cond{"user_id": userID}).OrderBy("id").All(&collections)

	if err != nil {
		return nil, err
	}

	return collections, nil
}

func RenameCollection(collectionID int64, userID int64, newName string) error {

	col := upperSession.Collection((&Collection{}).TableName())
	now := time.Now().UTC()

	// Ensure the collection belongs to the user
	res := col.Find(up.Cond{
		"id":      collectionID,
		"user_id": userID,
	})

	count, err := res.Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("collection not found or unauthorized")
	}

	// Update ONLY name + updated_at
	return res.Update(map[string]any{
		"name":       newName,
		"updated_at": now,
	})
}

func DeleteCollection(collectionID int64, userID int64) error {

	col := upperSession.Collection((&Collection{}).TableName())

	// Ensure the collection belongs to the user
	res := col.Find(up.Cond{
		"id":      collectionID,
		"user_id": userID,
	})

	count, err := res.Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("collection not found or unauthorized")
	}

	// Delete collection only — PostgreSQL cascades games & moves
	return res.Delete()
}
