package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

// SetupLogging configures the global logging settings
func SetupLogging() {
	// Set log level based on environment
	env := os.Getenv("LOG_LEVEL")
	switch env {
	case "DEBUG":
		log.SetLevel(log.LevelDebug)
	case "INFO":
		log.SetLevel(log.LevelInfo)
	case "WARN":
		log.SetLevel(log.LevelWarn)
	case "ERROR":
		log.SetLevel(log.LevelError)
	default:
		log.SetLevel(log.LevelInfo) // Default to INFO
	}
	
	log.Info("🔧 Logging configured successfully")
}

// RequestLogger returns an enhanced request logger middleware
func RequestLogger() fiber.Handler {
	return logger.New(logger.Config{
		Format: "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	})
}

// FileLogger returns a file-based logger middleware
func FileLogger(filename string) fiber.Handler {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Errorf("Failed to open log file: %v", err)
		return RequestLogger() // Fallback to stdout logger
	}

	return logger.New(logger.Config{
		Format: "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
		Done: func(c fiber.Ctx, logString []byte) {
			// Write to file manually since Output field is not available
			file.Write(logString)
		},
	})
}

// APILogger logs API-specific events
func LogAPIEvent(event, message string, data ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	if len(data) > 0 {
		log.Infof("[%s] %s: %s - Data: %+v", timestamp, event, message, data)
	} else {
		log.Infof("[%s] %s: %s", timestamp, event, message)
	}
}

// LogError logs errors with context
func LogError(context, message string, err error) {
	log.Errorf("[ERROR] %s: %s - %v", context, message, err)
}

// LogDebug logs debug information (only shown when debug level is enabled)
func LogDebug(context, message string, data ...interface{}) {
	if len(data) > 0 {
		log.Debugf("[DEBUG] %s: %s - %+v", context, message, data)
	} else {
		log.Debugf("[DEBUG] %s: %s", context, message)
	}
}