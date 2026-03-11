package main

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"golang.project.structure/config"
	"golang.project.structure/database"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	defer cancel()

	cfg := config.LoadConfig()
	//configure logger here

	// configure database here
	err := database.Initialize()
	if err != nil {
		log.Fatal("Error while initializing the database:", err)
	}
	app := fiber.New()
	//app.Use(middle)

	go func() {
		if err := app.Listen(cfg.APIPort); err != nil {
			log.Fatal(err.Error())
		}
	}()
	<-ctx.Done()
	log.Println("Shutting down server....")
	_ = app.Shutdown()
}
