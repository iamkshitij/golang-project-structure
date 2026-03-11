package config

import (
	"log"
	"os"
	"strconv"
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

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Invalid DB_PORT")
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
