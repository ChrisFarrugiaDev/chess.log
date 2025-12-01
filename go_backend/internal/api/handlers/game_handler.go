package handlers

import (
	"chess_log/go_backend/internal/appcore"
	"chess_log/go_backend/internal/helpers"
	"chess_log/go_backend/internal/models"
	"encoding/json"
	"net/http"
)

type GameHandler struct {
	App *appcore.App
}

// Store creates a game and its moves in a single transaction.
func (h *GameHandler) Store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// ----- Request body structure -----
	type MoveReq struct {
		MoveNumber int    `json:"move_number"`
		Color      string `json:"color"`
		BeforeFEN  string `json:"before"`
		AfterFEN   string `json:"after"`
		SAN        string `json:"san"`
		LAN        string `json:"lan"`
		Piece      string `json:"piece"`
		From       string `json:"from"`
		To         string `json:"to"`
		Flags      string `json:"flags"`
	}

	type GameReq struct {
		CollectionID int64   `json:"collection_id"`
		Name         string  `json:"name"`
		Orientation  string  `json:"orientation"`
		Notes        *string `json:"notes"`
	}

	type Request struct {
		Game  GameReq   `json:"game"`
		Moves []MoveReq `json:"moves"`
	}

	var req Request

	// ----- Decode JSON -----
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.RespondErrorJSON(w, http.StatusBadRequest, err, "Invalid JSON payload")
		return
	}

	// ----- Validation -----
	if req.Game.CollectionID == 0 {
		helpers.RespondErrorJSON(w, http.StatusBadRequest, nil, "collection_id is required")
		return
	}

	if req.Game.Name == "" {
		helpers.RespondErrorJSON(w, http.StatusBadRequest, nil, "game name is required")
		return
	}

	if len(req.Moves) == 0 {
		helpers.RespondErrorJSON(w, http.StatusBadRequest, nil, "moves array cannot be empty")
		return
	}

	// Default orientation
	if req.Game.Orientation == "" {
		req.Game.Orientation = "white"
	}

	// ----- Build Game model -----
	newGame := &models.Game{
		CollectionID: req.Game.CollectionID,
		Name:         req.Game.Name,
		Orientation:  req.Game.Orientation,
		Notes:        req.Game.Notes,
	}

	// ----- Build GameMove models -----
	moveModels := make([]*models.GameMove, 0, len(req.Moves))

	for _, mv := range req.Moves {
		moveModels = append(moveModels, &models.GameMove{
			MoveNumber: mv.MoveNumber,
			Color:      mv.Color,
			BeforeFEN:  mv.BeforeFEN,
			AfterFEN:   mv.AfterFEN,
			SAN:        mv.SAN,
			LAN:        mv.LAN,
			Piece:      mv.Piece,
			FromSquare: mv.From,
			ToSquare:   mv.To,
			Flags:      mv.Flags,
		})
	}

	// ----- Save game + moves using transaction -----
	ctx := r.Context()
	savedGame, err := models.CreateGameWithMoves(ctx, newGame, moveModels)
	if err != nil {
		helpers.RespondErrorJSON(w, http.StatusInternalServerError, err, "Failed to create game")
		return
	}

	// ----- Successful response -----
	helpers.RespondJSON(w, http.StatusCreated, map[string]any{
		"message": "Game created successfully",
		"game":    savedGame,
	})
}
