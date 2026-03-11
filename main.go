package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"golang.project.structure/api/routes"
	"golang.project.structure/config"
	"golang.project.structure/database"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// Configure Fiber logging
	log.SetLevel(log.LevelInfo) // Set to LevelDebug for more detailed logs
	
	cfg := config.LoadConfig()
	
	// Initialize database
	err := database.Initialize()
	if err != nil {
		log.Fatal("Error while initializing the database:", err)
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Go Fiber API v1.0.0",
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			log.Errorf("Request error: %v", err)
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Setup routes
	routes.SetupRoutes(app, database.OraDB)

	// Start server in a goroutine
	go func() {
		if err := app.Listen(cfg.APIPort); err != nil {
			log.Fatal(err.Error())
		}
	}()

	log.Infof("🚀 Server started on %s", cfg.APIPort)
	log.Info("📊 Press Ctrl+C to shutdown")
	
	// Wait for interrupt signal
	<-ctx.Done()
	log.Info("🛑 Shutting down server....")
	_ = app.Shutdown()
	log.Info("✅ Server shutdown complete")
}
