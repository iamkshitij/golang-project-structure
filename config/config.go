package config

import (
	"os"
	"strconv"
	
	"github.com/gofiber/fiber/v3/log"
	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBName     string
	DBPassword string
	APIPort    string
	BaseURL    string
}

func LoadConfig() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Warn("Warning: .env file not found, using system environment variables")
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}

	apiPort := os.Getenv("API_PORT")
	if apiPort != "" && apiPort[0] != ':' {
		apiPort = ":" + apiPort
	}

	cfg := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     dbPort,
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		APIPort:    apiPort,
		BaseURL:    os.Getenv("BASE_URL"),
	}

	if cfg.DBHost == "" ||
		cfg.DBPort == 0 ||
		cfg.BaseURL == "" ||
		cfg.DBPassword == "" ||
		cfg.APIPort == "" {
		log.Fatal("Missing required env variables")
	}

	return cfg
}

func getEnv(key, fallback string) string {

	val := os.Getenv(key)

	if val == "" {
		return fallback
	}

	return val
}
