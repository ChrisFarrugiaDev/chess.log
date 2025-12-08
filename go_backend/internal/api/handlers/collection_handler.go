package handlers

import (
	"chess_log/go_backend/internal/api/middleware"
	"chess_log/go_backend/internal/appcore"
	"chess_log/go_backend/internal/helpers"
	"chess_log/go_backend/internal/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
		Name string `json:"name"`
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

func (h *CollectionHandler) Rename(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, ok := r.Context().Value(middleware.ContextUserClaims).(jwt.MapClaims)
	if !ok {
		helpers.RespondErrorJSON(w, http.StatusUnauthorized, nil, "Missing auth claims!!")
		return
	}
	userID := int64(claims["UserID"].(float64))

	id := chi.URLParam(r, "id")
	if id == "" {
		helpers.RespondErrorJSON(w, http.StatusBadRequest, nil, "Invalid collection id")
		return
	}

	collectionInt, err := strconv.Atoi(id)
	if err != nil {
		helpers.RespondErrorJSON(w, http.StatusBadRequest, err, "Invalid collection id: must be a number")
		return
	}
	collectionID := int64(collectionInt)

	type Request struct {
		Name string `json:"name"`
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.RespondErrorJSON(w, http.StatusBadRequest, err, "Invalid request payload")
		return
	}

	if req.Name == "" {
		helpers.RespondErrorJSON(w, http.StatusBadRequest, nil, "Collection name is required")
		return
	}

	// Call model
	err = models.RenameCollection(collectionID, userID, req.Name)
	if err != nil {
		helpers.RespondErrorJSON(w, http.StatusInternalServerError, err, "Failed to rename collection")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, map[string]any{
		"message": "Collection renamed successfully",
	})
}

// ---------------------------------------------------------------------

func (h *CollectionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, ok := r.Context().Value(middleware.ContextUserClaims).(jwt.MapClaims)
	if !ok {
		helpers.RespondErrorJSON(w, http.StatusUnauthorized, nil, "Missing auth claims!!")
		return
	}
	userID := int64(claims["UserID"].(float64))

	// Get ID from URL
	id := chi.URLParam(r, "id")
	if id == "" {
		helpers.RespondErrorJSON(w, http.StatusBadRequest, nil, "Invalid collection id")
		return
	}

	// Convert to int64 safely
	collectionInt, err := strconv.Atoi(id)
	if err != nil {
		helpers.RespondErrorJSON(w, http.StatusBadRequest, err, "Invalid collection id: must be a number")
		return
	}
	collectionID := int64(collectionInt)

	// Call model
	err = models.DeleteCollection(collectionID, userID)
	if err != nil {
		helpers.RespondErrorJSON(w, http.StatusInternalServerError, err, "Failed to delete collection")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, map[string]any{
		"message": "Collection deleted successfully",
	})
}
