package handlers

import (
	"chess_log/go_backend/internal/api/middleware"
	"chess_log/go_backend/internal/appcore"
	"chess_log/go_backend/internal/helpers"
	"chess_log/go_backend/internal/models"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type CollectionHandler struct {
	App *appcore.App
}

func (h *CollectionHandler) Store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, ok := r.Context().Value(middleware.ContextUserClaims).(jwt.MapClaims)
	if !ok {
		helpers.RespondErrorJSON(w, http.StatusUnauthorized, nil, "Missing auth claims!!")
		return
	}

	userID := int64(claims["UserID"].(float64))

	type Request struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	var req Request

	// Parse JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.RespondErrorJSON(w, http.StatusBadRequest, err, "Invalid request payload")
		return
	}

	// Validate input
	if req.Name == "" {
		helpers.RespondErrorJSON(w, http.StatusBadRequest, nil, "Collection name is required")
		return
	}

	// Create model
	collection := &models.Collection{
		Name:   req.Name,
		Color:  req.Color,
		UserID: userID,
	}

	// Insert into DB
	created, err := collection.Create()
	if err != nil {
		if errors.Is(err, models.ErrDuplicateCollection) {
			helpers.RespondErrorJSON(w, http.StatusConflict, nil, "A collection with this name already exists")
			return
		}

		helpers.RespondErrorJSON(w, http.StatusInternalServerError, err, "Failed to create collection")
		return
	}

	// Success
	helpers.RespondJSON(w, http.StatusCreated, map[string]any{
		"message":    "Collection created successfully",
		"collection": created,
	})
}
