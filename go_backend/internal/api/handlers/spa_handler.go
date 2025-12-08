package handlers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"chess_log/go_backend/internal/logger"

	"go.uber.org/zap"
)

type SpaHandler struct {
	tmpl *template.Template // Pre-parsed index.html template
}

// NewSpaHandler:
// Loads and parses index.html ONCE at application startup.
// This avoids re-reading the file on every request.
func NewSpaHandler(staticPath, indexFile string) (*SpaHandler, error) {
	// Build absolute path to dist/index.html
	fullPath := filepath.Join(staticPath, indexFile)

	// Parse the HTML file as a Go template
	t, err := template.ParseFiles(fullPath)
	if err != nil {
		return nil, err
	}

	// Store parsed template inside handler
	return &SpaHandler{tmpl: t}, nil
}

func (h *SpaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Template variables passed into index.html (e.g. VITE-like env)
	data := map[string]string{
		"GO_APP_URL": os.Getenv("APP_URL"),
	}

	// Standard no-cache headers for SPA
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	// Render the pre-loaded template and write it directly to response
	if err := h.tmpl.Execute(w, data); err != nil {
		logger.Error("Failed to execute SPA template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
