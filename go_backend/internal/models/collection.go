package models

import (
	"errors"
	"strings"
	"time"
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
