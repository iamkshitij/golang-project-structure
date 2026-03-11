package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/jmoiron/sqlx"
	"golang.project.structure/api/middleware"
	"golang.project.structure/api/resources/user"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App, db *sqlx.DB) {
	// Setup logging configuration
	middleware.SetupLogging()
	
	// Global middleware
	app.Use(middleware.RequestLogger())
	app.Use(middleware.CORS())
	
	// Root endpoint
	app.Get("/", func(c fiber.Ctx) error {
		log.Info("Root endpoint accessed")
		return c.JSON(fiber.Map{
			"message": "Welcome to Go Fiber API",
			"version": "1.0.0",
			"docs":    "/api/v1/health",
		})
	})
	
	// API version 1 group
	api := app.Group("/api/v1")
	
	// Health check endpoint
	api.Get("/health", func(c fiber.Ctx) error {
		log.Debug("Health check endpoint accessed")
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "API is running",
			"version": "1.0.0",
		})
	})
	
	// User routes
	user.SetupRoutes(api, db)
	
	log.Info("🛣️  All routes configured successfully")
	
	// Add more resource routes here as you expand
	// product.SetupRoutes(api, db)
	// order.SetupRoutes(api, db)
}