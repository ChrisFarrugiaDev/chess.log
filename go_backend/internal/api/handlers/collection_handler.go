package handlers

import (
	"chess_log/go_backend/internal/appcore"
	"chess_log/go_backend/internal/helpers"
	"chess_log/go_backend/internal/models"
	"encoding/json"
	"errors"
	"net/http"
)

type CollectionHandler struct {
	App *appcore.App
}

func (h *CollectionHandler) Store(w http.ResponseWriter, r *http.Request) {

	type Request struct {
		UserID int64  `db:"user_id" json:"user_id"`
		Name   string `db:"name" json:"name"`
	}

	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		helpers.RespondWithError(w, err, "Failed to create collection.", http.StatusInternalServerError)
	}

	newCollection := &models.Collection{
		Name:   req.Name,
		UserID: req.UserID,
	}

	newCollection, err = newCollection.Create()

	if err != nil {
		if errors.Is(err, models.ErrDuplicateCollection) {
			http.Error(w, "Collection with this name already exists.", http.StatusConflict)
			return
		}
		helpers.RespondWithError(w, err, "Failed to create collection.", http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"message": "Collection updated successfully.",
		"data": map[string]any{
			"collection": newCollection,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		helpers.RespondWithError(w, err, "Failed to encode response.", http.StatusInternalServerError)
	}

}
