package helpers

import (
	"chess_log/go_backend/internal/logger"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"

	"go.uber.org/zap"
)

func RespondErrorJSON(w http.ResponseWriter, status int, err error, message string, depth ...int) {
	// Determine call depth
	callDepth := 1
	if len(depth) > 0 {
		callDepth = depth[0]
	}

	// Only log internal server errors
	if status == http.StatusInternalServerError {
		_, file, line, _ := runtime.Caller(callDepth)

		logger.Error(
			message,
			zap.Error(err),
			zap.String("file", file),
			zap.Int("line", line),
		)
	}

	// Build user-facing message
	var userMessage string

	if os.Getenv("DEBUG") == "true" || os.Getenv("DEBUG") == "1" {
		userMessage = fmt.Sprintf("%s: %v", message, err)
	} else {
		userMessage = message
	}

	RespondJSON(w, status, map[string]any{
		"success": false,
		"message": userMessage,
	})
}

func RespondJSON(w http.ResponseWriter, status int, data map[string]any, depth ...int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// If "success" key not present, add it as true
	if _, ok := data["success"]; !ok {
		data["success"] = true
	}

	_ = json.NewEncoder(w).Encode(data)
}
