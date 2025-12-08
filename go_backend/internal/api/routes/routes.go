package routes

import (
	"chess_log/go_backend/internal/appcore"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Routes(app *appcore.App) http.Handler {
	mux := chi.NewRouter()

	// Middleware
	mux.Use(middleware.Recoverer)

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	mux.Route("/api", func(r chi.Router) {
		r.Mount("/collections", CollectionRoutes(app))
		r.Mount("/games", GameRoutes(app))
		r.Mount("/auth", AuthRoutes(app))
		r.Mount("/user", UserRoutes(app))
		// r.Mount("/game", GameRoutes())
	})

	mux.Mount("/", SpaRoutes())

	return mux
}
