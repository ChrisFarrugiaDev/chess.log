package routes

import (
	"chess_log/go_backend/internal/api/handlers"
	"chess_log/go_backend/internal/appcore"

	"github.com/go-chi/chi/v5"
)

func AuthRoutes(app *appcore.App) chi.Router {
	r := chi.NewRouter()

	gameHandler := &handlers.AuthHandler{
		App: app,
	}

	// r.Use(middleware.JWTAuthMiddleware)

	r.Post("/register", gameHandler.RegisterUser)

	return r
}
