package routes

import (
	"chess_log/go_backend/internal/api/handlers"
	"chess_log/go_backend/internal/api/middleware"
	"chess_log/go_backend/internal/appcore"

	"github.com/go-chi/chi/v5"
)

func GameRoutes(app *appcore.App) chi.Router {
	r := chi.NewRouter()

	gameHandler := &handlers.GameHandler{
		App: app,
	}

	r.Group(func(pr chi.Router) {
		pr.Use(middleware.JWTAuthMiddleware)

		pr.Post("/", gameHandler.Store)
		pr.Get("/moves/{game_id}", gameHandler.GetMoves)
	})

	return r
}
