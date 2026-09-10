package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoUri   string
	MongoDb    string
	ServerPort string
}

func extractEnv(key string) (string, error) {
	val := os.Getenv(key)

	if val == "" {
		return "", fmt.Errorf("Missing req env")
	}

	return val, nil
}

func Load() (Config, error) {

	if err := godotenv.Load("../../.env"); err != nil {
		return Config{}, fmt.Errorf("Failed to load env")
	}

	mongoUri, err := extractEnv("MONGO_URI")

	if err != nil {
		return Config{}, err
	}

	mongoDbName, err := extractEnv("MONGO_DB_NAME")

	if err != nil {
		return Config{}, err
	}

	port, err := extractEnv("PORT")

	if err != nil {
		return Config{}, err
	}

	return Config{
		MongoUri:   mongoUri,
		MongoDb:    mongoDbName,
		ServerPort: port,
	}, nil
}
