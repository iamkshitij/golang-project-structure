package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

// Logger returns a logger middleware with custom format
func Logger() fiber.Handler {
	return logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} - ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05",
	})
}