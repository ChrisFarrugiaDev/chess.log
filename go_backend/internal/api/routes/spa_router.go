package routes

import (
	"chess_log/go_backend/internal/api/handlers"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func SpaRoutes() chi.Router {
	r := chi.NewRouter()

	distPath := "../vue_frontend/dist"

	// Load and parse the Vue index.html template once at startup.
	spaHandler, err := handlers.NewSpaHandler(distPath, "index.html")
	if err != nil {
		panic("Failed to initialize SPA Handler: " + err.Error())
	}

	// Serve static asset files (JS, CSS, images) under /assets.
	// The FileServer serves files directly from dist/assets.
	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir(filepath.Join(distPath, "assets")))))

	// Handle all other GET routes.
	r.Get("/*", func(w http.ResponseWriter, req *http.Request) {

		// Requested URL path (e.g. /favicon.ico, /logo.svg, /index.js)
		reqPath := req.URL.Path

		// Build absolute path inside the dist folder
		fullPath := filepath.Join(distPath, reqPath)

		// Check if file exists on disk and is not a directory
		stat, err := os.Stat(fullPath)

		if err == nil && !stat.IsDir() {
			// If the file exists (e.g. favicon), serve it directly.
			http.ServeFile(w, req, fullPath)
			return
		}

		// Otherwise → it's a Vue route (e.g. /login, /game/create)
		// Serve index.html so the Vue SPA router takes over.
		spaHandler.ServeHTTP(w, req)
	})

	return r
}
