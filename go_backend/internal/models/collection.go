package models

import (
	"errors"
	"strings"
	"time"

	up "github.com/upper/db/v4"
)

var ErrDuplicateCollection = errors.New("collection with this name already exists")

type Collection struct {
	ID        int64     `db:"id,omitempty" json:"id"`
	UserID    int64     `db:"user_id" json:"user_id"`
	Name      string    `db:"name" json:"name"`
	Color     string    `db:"color" json:"color"`
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
