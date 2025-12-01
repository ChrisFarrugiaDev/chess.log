package handlers

import (
	"chess_log/go_backend/internal/appcore"
	"chess_log/go_backend/internal/helpers"
	"chess_log/go_backend/internal/mailer"
	"chess_log/go_backend/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type AuthHandler struct {
	App *appcore.App
}

func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type Request struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		Password1 string `json:"password1"`
		Password2 string `json:"password2"`
	}

	var req Request

	// Parse JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.RespondErrorJSON(w, http.StatusBadRequest, err, "Invalid request payload")
		return
	}

	// Validate passwords
	if req.Password1 != req.Password2 {
		helpers.RespondErrorJSON(w, http.StatusBadRequest, nil, "Passwords do not match")
		return
	}

	if len(req.Password1) < 6 {
		helpers.RespondErrorJSON(w, http.StatusBadRequest, nil, "Password must be at least 6 characters")
		return
	}

	// Hash password
	hash, err := helpers.HashPassword(req.Password1)
	if err != nil {
		helpers.RespondErrorJSON(w, http.StatusInternalServerError, err, "Failed to hash password")
		return
	}

	// Build user model
	user := &models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: hash,
	}

	// Save user
	createdUser, err := user.Create()
	if err != nil {
		if errors.Is(err, models.ErrDuplicateUser) {
			helpers.RespondErrorJSON(w, http.StatusConflict, nil, "User with this email already exists")
			return
		}
		helpers.RespondErrorJSON(w, http.StatusInternalServerError, err, "Failed to create user")
		return
	}

	// Generate email verification token

	jwtPayload := map[string]any{
		"Subject": "email_validation", // use underscore for consistency
		"Email":   createdUser.Email,
		"UserID":  createdUser.ID,
	}

	exp := time.Now().Add(15 * time.Minute)
	token, err := helpers.GenerateJWT(jwtPayload, exp)
	if err != nil {
		helpers.RespondErrorJSON(w, http.StatusInternalServerError, err, "Failed to generate verification token")
		return
	}

	baseURL := os.Getenv("APP_URL")
	link := fmt.Sprintf("%s/verify-email?token=%s", baseURL, token)

	plainBody, htmlBody := mailer.GenerateEmailValidationTemp(link)

	// Send verification email

	msg, err := h.App.Mailer.CreateMessage(
		createdUser.Email,
		"Verify your email - ChessLog",
		plainBody,
		htmlBody,
	)
	if err != nil {
		helpers.RespondErrorJSON(w, http.StatusInternalServerError, err, "Failed to compose verification email")
		return
	}

	if err := h.App.Mailer.Send(msg); err != nil {
		helpers.RespondErrorJSON(w, http.StatusInternalServerError, err, "Failed to send verification email")
		return
	}

	// Success

	helpers.RespondJSON(w, http.StatusOK, map[string]any{
		"message": "User created successfully. Please verify your email.",
		"user":    createdUser,
	})
}

func (h *AuthHandler) ValidatedEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type Request struct {
		Token string `json:"token"`
	}

	var req Request

	// Parse JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.RespondErrorJSON(w, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	// Validate JWT
	claims, err := helpers.ValidateJWT(req.Token)
	if err != nil {
		helpers.RespondErrorJSON(w, http.StatusUnauthorized, err, "Invalid or expired token")
		return
	}

	userID, ok := claims["UserID"].(float64)
	if !ok {
		helpers.RespondErrorJSON(w, http.StatusUnauthorized, nil, "Invalid token payload")
		return
	}

	email, ok := claims["Email"].(string)
	if !ok {
		helpers.RespondErrorJSON(w, http.StatusUnauthorized, nil, "Invalid token payload")
		return
	}

	tokenType, ok := claims["Subject"].(string)
	if !ok || tokenType != "email_validation" {
		helpers.RespondErrorJSON(w, http.StatusUnauthorized, nil, "Invalid token type")
		return
	}

	// Get user

	user, err := models.GetUserByID(int64(userID))
	if err != nil {
		helpers.RespondErrorJSON(w, http.StatusNotFound, err, "User not found")
		return
	}

	if user.Email != email {
		helpers.RespondErrorJSON(w, http.StatusUnauthorized, nil, "Token does not match user")
		return
	}

	// Already verified?
	if user.EmailVerified {
		helpers.RespondJSON(w, http.StatusOK, map[string]any{
			"message": "Email already verified",
			"user":    user,
		})
		return
	}

	// Mark verified

	user.EmailVerified = true
	if err := user.Update(); err != nil {
		helpers.RespondErrorJSON(w, http.StatusInternalServerError, err, "Failed to update user")
		return
	}

	// Success

	helpers.RespondJSON(w, http.StatusOK, map[string]any{
		"message": "Email successfully verified",
		"user":    user,
	})
}

func (h *AuthHandler) ResendVerificationEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type Request struct {
		Token string `json:"token"`
	}

	var req Request

	// Parse JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.RespondErrorJSON(w, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	// Extract claims even if expired
	claims, err := helpers.ValidateJWTAllowExpired(req.Token)
	if err != nil {
		helpers.RespondErrorJSON(w, http.StatusUnauthorized, err, "Invalid token")
		return
	}

	userID, ok := claims["UserID"].(float64)
	if !ok {
		helpers.RespondErrorJSON(w, http.StatusUnauthorized, nil, "Invalid token payload")
		return
	}

	email, ok := claims["Email"].(string)
	if !ok {
		helpers.RespondErrorJSON(w, http.StatusUnauthorized, nil, "Invalid token payload")
		return
	}

	tokenType, ok := claims["Subject"].(string)
	if !ok || tokenType != "email_validation" {
		helpers.RespondErrorJSON(w, http.StatusUnauthorized, nil, "Invalid token type")
		return
	}

	// Get user
	user, err := models.GetUserByID(int64(userID))
	if err != nil {
		helpers.RespondErrorJSON(w, http.StatusNotFound, err, "User not found")
		return
	}

	if user.Email != email {
		helpers.RespondErrorJSON(w, http.StatusUnauthorized, nil, "Token does not match user")
		return
	}

	// Already verified
	if user.EmailVerified {
		helpers.RespondJSON(w, http.StatusOK, map[string]any{
			"message": "Email already verified",
			"user":    user,
		})
		return
	}

	// --- Generate NEW verification token ---
	jwtPayload := map[string]any{
		"Subject": "email_validation",
		"Email":   user.Email,
		"UserID":  user.ID,
	}

	exp := time.Now().Add(15 * time.Minute)
	newToken, err := helpers.GenerateJWT(jwtPayload, exp)
	if err != nil {
		helpers.RespondErrorJSON(w, http.StatusInternalServerError, err, "Failed to generate new verification token")
		return
	}

	baseURL := os.Getenv("APP_URL")
	link := fmt.Sprintf("%s/verify-email?token=%s", baseURL, newToken)

	plainBody, htmlBody := mailer.GenerateEmailValidationTemp(link)

	// Send email
	msg, err := h.App.Mailer.CreateMessage(
		user.Email,
		"Resend Email Verification - ChessLog",
		plainBody,
		htmlBody,
	)
	if err != nil {
		helpers.RespondErrorJSON(w, http.StatusInternalServerError, err, "Failed to compose verification email")
		return
	}

	if err := h.App.Mailer.Send(msg); err != nil {
		helpers.RespondErrorJSON(w, http.StatusInternalServerError, err, "Failed to send verification email")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, map[string]any{
		"message": "Verification email resent successfully",
	})
}
