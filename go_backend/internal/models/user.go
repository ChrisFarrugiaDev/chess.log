package models

import (
	"errors"
	"strings"
	"time"
)

var ErrDuplicateUser = errors.New("user with this name already exists")

type User struct {
	ID            int64     `db:"id,omitempty" json:"id"`
	Name          string    `db:"name" json:"name"`
	Email         string    `db:"email" json:"email"`
	EmailVerified bool      `db:"email_verified" json:"email_verified"`
	PasswordHash  string    `db:"password_hash" json:"-"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

func (m *User) TableName() string {
	return "users"
}

func (m *User) Create() (*User, error) {

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
