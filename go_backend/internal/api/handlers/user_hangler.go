package handlers

import (
	"chess_log/go_backend/internal/api/middleware"
	"chess_log/go_backend/internal/appcore"
	"chess_log/go_backend/internal/helpers"
	"chess_log/go_backend/internal/models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
	App *appcore.App
}

func (h *UserHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	// Extract JWT claims
	claims, ok := r.Context().Value(middleware.ContextUserClaims).(jwt.MapClaims)
	if !ok {
		helpers.RespondErrorJSON(w, http.StatusUnauthorized, nil, "Missing auth claims!!")
		return
	}

	userID := int64(claims["UserID"].(float64))

	// Fetch games
	games, err := h.App.Models.Game.GetByUserID(userID)
	if err != nil {
		helpers.RespondErrorJSON(w, http.StatusInternalServerError, err, "Failed to fetch games")
		return
	}

	// Fetch collections
	collections, err := h.App.Models.Collection.GetByUserID(userID)
	if err != nil {
		helpers.RespondErrorJSON(w, http.StatusInternalServerError, err, "Failed to fetch collections")
		return
	}

	mappedGames := make(map[int64][]*models.Game)

	for _, c := range collections {
		var gms []*models.Game

		mappedGames[c.ID] = gms
	}

	for _, g := range games {
		mappedGames[g.CollectionID] = append(mappedGames[g.CollectionID], g)
	}

	// Success — use 200 OK
	helpers.RespondJSON(w, http.StatusOK, map[string]any{
		"message":     "User profile retrieved successfully",
		"collections": collections,
		"games":       mappedGames,
	})
}
