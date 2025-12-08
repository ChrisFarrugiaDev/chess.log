package routes

import (
	"chess_log/go_backend/internal/api/handlers"
	"chess_log/go_backend/internal/api/middleware"
	"chess_log/go_backend/internal/appcore"

	"github.com/go-chi/chi/v5"
)

func CollectionRoutes(app *appcore.App) chi.Router {
	r := chi.NewRouter()

	collectionHandler := &handlers.CollectionHandler{
		App: app,
	}

	r.Group(func(pr chi.Router) {
		pr.Use(middleware.JWTAuthMiddleware)

		// Create collection
		pr.Post("/", collectionHandler.Store)

		// Rename collection
		pr.Put("/{id}", collectionHandler.Rename)

		// Delete collection
		pr.Delete("/{id}", collectionHandler.Delete)
	})

	return r
}
