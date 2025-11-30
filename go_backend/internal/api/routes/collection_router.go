package routes

import (
	"chess_log/go_backend/internal/api/handlers"
	"chess_log/go_backend/internal/appcore"

	"github.com/go-chi/chi/v5"
)

func CollectionRoutes(app *appcore.App) chi.Router {
	r := chi.NewRouter()

	colletionHandler := &handlers.CollectionHandler{
		App: app,
	}

	// r.Use(middleware.JWTAuthMiddleware)

	r.Post("/", colletionHandler.Store)

	return r
}
