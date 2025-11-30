package httpserver

import (
	"chess_log/go_backend/internal/api/routes"
	"chess_log/go_backend/internal/appcore"
	"chess_log/go_backend/internal/logger"
	"context"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

// Server wraps the http.Server.
type Server struct {
	HTTPServer *http.Server
	Port       string
}

func NewHttpServer(port string, app *appcore.App) *Server {
	mux := http.NewServeMux()
	mux.Handle("/", routes.Routes(app))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	return &Server{
		HTTPServer: srv,
		Port:       port,
	}
}

func (s *Server) Start() {
	go func() {
		logger.Info("HTTP server started",
			zap.String("address", s.HTTPServer.Addr),
		)

		if err := s.HTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("HTTP server stopped with error", zap.Error(err))
		} else {
			logger.Info("HTTP server stopped gracefully")
		}
	}()
}

func (s *Server) Shutdown(ctx context.Context) error {
	logger.Info("Gracefully shutting down HTTP server...")

	err := s.HTTPServer.Shutdown(ctx)
	if err != nil {
		logger.Error("Error shutting down server", zap.Error(err))
		return err
	}

	logger.Info("HTTP server successfully stopped")
	return nil
}
