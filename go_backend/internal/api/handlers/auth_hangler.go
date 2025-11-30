package handlers

import (
	"chess_log/go_backend/internal/appcore"
	"chess_log/go_backend/internal/helpers"
	"chess_log/go_backend/internal/mailer"
	"chess_log/go_backend/internal/models"
	"encoding/json"
	"errors"
	"net/http"
)

type AuthHandler struct {
	App *appcore.App
}

func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

	type Request struct {
		Name      string `db:"name" json:"name"`
		Email     string `db:"email" json:"email"`
		Password1 string `db:"password" json:"password1"`
		Password2 string `db:"password" json:"password2"`
	}

	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		helpers.RespondWithError(w, err, "Failed to create user.", http.StatusInternalServerError)
	}

	if req.Password1 != req.Password2 {
		http.Error(w, "Passwords do not match.", http.StatusBadRequest)
		return
	}

	if len(req.Password1) < 6 {
		http.Error(w, "Password must be at least 6 characters long.", http.StatusBadRequest)
		return
	}

	hash, err := helpers.HashPassword(req.Password1)

	if err != nil {
		helpers.RespondWithError(w, err, "Failed to create user.", http.StatusInternalServerError)
	}

	newUser := &models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: hash,
	}

	err = mailer.TwilioEmailSender("Register User", req.Email)

	if err != nil {
		helpers.RespondWithError(w, err, "Failed to create user!!", http.StatusInternalServerError)
	}

	// Attempt to create the user
	createdUser, err := newUser.Create()

	if err != nil {
		if errors.Is(err, models.ErrDuplicateUser) {
			http.Error(w, "User with this email already exists.", http.StatusConflict)
			return
		}
		helpers.RespondWithError(w, err, "Failed to create user.", http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"message": "User created successfully.",
		"data": map[string]any{
			"user": createdUser,
		},
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		helpers.RespondWithError(w, err, "Failed to encode response.", http.StatusInternalServerError)
	}

}
