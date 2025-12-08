package routes

import (
	"chess_log/go_backend/internal/api/handlers"
	"chess_log/go_backend/internal/appcore"

	"github.com/go-chi/chi/v5"
)

func AuthRoutes(app *appcore.App) chi.Router {
	r := chi.NewRouter()

	authHandler := &handlers.AuthHandler{
		App: app,
	}

	r.Post("/register", authHandler.RegisterUser)
	r.Post("/verify-email", authHandler.ValidatedEmail)
	r.Post("/resend-verification-email", authHandler.ResendVerificationEmail)
	r.Post("/login", authHandler.LoginUser)
	r.Post("/forgot-password", authHandler.RequestPasswordReset)
	r.Post("/reset-password", authHandler.ResetPassword)

	return r
}
