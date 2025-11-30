package helpers

import (
	"chess_log/go_backend/internal/logger"
	"fmt"
	"net/http"
	"os"
	"runtime"

	"go.uber.org/zap"
)

func RespondWithError(w http.ResponseWriter, err error, message string, statusCode int, depth ...int) {

	// Set the default depth to 1 if none is provided
	callDepth := 1
	if len(depth) > 0 {
		callDepth = depth[0]
	}

	_, file, line, _ := runtime.Caller(callDepth)
	var userError string

	if os.Getenv("DEBUG") == "true" || os.Getenv("DEBUG") == "1" {
		userError = fmt.Sprintf("%s: at %s:%d: \n%v", message, file, line, err)
	} else {
		userError = message
	}

	logger.Error(message, zap.Error(err), zap.Int("call_deep", callDepth))

	http.Error(w, userError, statusCode)
}
