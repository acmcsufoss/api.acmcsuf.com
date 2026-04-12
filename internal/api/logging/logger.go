package logging

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// NewLogger returns the default server logger.
func NewLogger() *slog.Logger {
	opts := &slog.HandlerOptions{AddSource: true}
	return slog.New(slog.NewJSONHandler(os.Stdout, opts)).With("component", "SERVER")
}

// Fatal logs a message at Error level then exits with code 1.
func Fatal(logger *slog.Logger, msg string, args ...any) {
	logger.Error(msg, args...)
	os.Exit(1)
}

// Fatal logs a formatted message at Error level then exits with code 1.
func Fatalf(logger *slog.Logger, format string, args ...any) {
	logger.Error("fatal", "msg", fmt.Sprintf(format, args...))
	os.Exit(1)
}

// RequestLogger returns a gin middleware that logs each request using the provided logger.
func RequestLogger(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		logger.Info("request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency_ms", time.Since(start).Milliseconds(),
			"client_ip", c.ClientIP(),
		)
	}
}
