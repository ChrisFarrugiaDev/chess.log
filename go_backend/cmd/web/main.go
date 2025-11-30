package main

import (
	"chess_log/go_backend/internal/appcore"
	"chess_log/go_backend/internal/httpserver"
	"chess_log/go_backend/internal/logger"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"go.uber.org/zap"
)

var app appcore.App

func main() {

	// -----------------------------------------------------------------
	// Load enviermental variables
	loadEnv()

	// -----------------------------------------------------------------
	// Init Logger
	logger.InitLogger()

	// -----------------------------------------------------------------
	// Initialize database connection

	initializeDatabase()

	// -----------------------------------------------------------------
	// Create a context that will be cancelled when an interrupt or termination signal is received.

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	// stop() will stop the context from listening for further OS signals
	defer stop()

	// -----------------------------------------------------------------

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		logger.Warn("HTTP_PORT environment variable not set, using default 4567")
		httpPort = "4567"
	}
	httpserver := httpserver.NewHttpServer(httpPort, &app)

	httpserver.Start()

	// -----------------------------------------------------------------

	// Block main goroutine until context is cancelled by an OS signal (e.g. CTRL+C).
	// This keeps the main function alive while the HTTP server runs in the background.
	<-ctx.Done()

	// -----------------------------------------------------------------
	// Gracefully shutdown

	// Gracefully shutdown the HTTP server
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := httpserver.Shutdown(shutdownCtx)
	if err != nil {
		logger.Error("Error during server shutdown", zap.Error(err))

	}

	// Close DB
	if app.DB != nil {
		app.DB.Close()
		logger.Log.Info("Database connection pool closed gracefully.")
	}

	// Ensure logs are flushed on exit
	if err := logger.Log.Sync(); err != nil && !isInvalidSyncError(err) {
		fmt.Fprintf(os.Stderr, "Logger sync failed: %v\n", err)
	}
}

// isInvalidSyncError returns true if the error is the harmless
// "invalid argument" sync error for /dev/stderr (common in dev)
//
//	Zap has a known harmless error on many systems.
func isInvalidSyncError(err error) bool {
	return strings.Contains(err.Error(), "invalid argument") && strings.Contains(err.Error(), "/dev/stderr")
}
