package config

import "os"

type Config struct {
	Port             string
	DatabaseURL      string
	SessionSecretKey string
}

func Load() Config {
	return Config{
		Port:             getEnv("PORT", "8080"),
		DatabaseURL:      getEnv("DATABASE_URL", "postgres://pizzauser:pizzapass@localhost:5432/pizza_tracker?sslmode=disable"),
		SessionSecretKey: getEnv("SESSION_SECRET_KEY", "pizza-order-secret-key"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
