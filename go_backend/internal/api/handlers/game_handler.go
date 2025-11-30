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
		Notes        *string `json:"notes"` // nullable
		// SortOrder    int     `json:"sort_order"`
	}

	type Request struct {
		Game  GameReq   `json:"game"`
		Moves []MoveReq `json:"moves"`
	}

	var req Request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.RespondWithError(w, err, "Invalid JSON body.", http.StatusBadRequest)
		return
	}

	// ----- Build Game model -----
	newGame := &models.Game{
		CollectionID: req.Game.CollectionID,
		Name:         req.Game.Name,
		Orientation:  req.Game.Orientation,
		// SortOrder:    req.Game.SortOrder,
		Notes: req.Game.Notes,
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

		helpers.RespondWithError(w, err, "Failed to create game.", http.StatusInternalServerError)
		return
	}

	// ----- Successful response -----
	response := map[string]any{
		"message": "Game created successfully.",
		"data": map[string]any{
			"game": savedGame,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		helpers.RespondWithError(w, err, "Failed to encode response.", http.StatusInternalServerError)
	}
}
