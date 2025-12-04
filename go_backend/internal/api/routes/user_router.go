package routes

import (
	"chess_log/go_backend/internal/api/handlers"
	"chess_log/go_backend/internal/api/middleware"
	"chess_log/go_backend/internal/appcore"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(app *appcore.App) chi.Router {
	r := chi.NewRouter()

	userHandler := &handlers.UserHandler{
		App: app,
	}

	r.Group(func(pr chi.Router) {
		pr.Use(middleware.JWTAuthMiddleware)

		pr.Get("/profile", userHandler.GetUserProfile)
	})

	return r
}
