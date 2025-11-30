package models

import "time"

type GameMove struct {
	ID         int64  `db:"id,omitempty" json:"id"`
	GameID     int64  `db:"game_id" json:"game_id"`
	MoveNumber int    `db:"move_number" json:"move_number"`
	Color      string `db:"color" json:"color"`

	BeforeFEN string `db:"before_fen" json:"before"`
	AfterFEN  string `db:"after_fen" json:"after"`

	SAN        string `db:"san" json:"san"`
	LAN        string `db:"lan" json:"lan"`
	Piece      string `db:"piece" json:"piece"`
	FromSquare string `db:"from_square" json:"from"`
	ToSquare   string `db:"to_square" json:"to"`
	Flags      string `db:"flags" json:"flags"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (m *GameMove) TableName() string {
	return "game_moves"
}

func (m *GameMove) Create() (*GameMove, error) {

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
