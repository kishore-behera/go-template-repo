package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// Config holds database configuration
type Config struct {
	URI      string
	DBName   string
	Username string
	Password string
}

// NewConfig creates a new database configuration
func NewConfig() *Config {
	return &Config{
		URI:      getEnvOrDefault("MONGODB_URI", "mongodb://localhost:27017"),
		DBName:   getEnvOrDefault("MONGODB_DB", "app"),
		Username: getEnvOrDefault("MONGODB_USER", ""),
		Password: getEnvOrDefault("MONGODB_PASSWORD", ""),
	}
}

// Connect establishes a connection to MongoDB
func Connect(config *Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(config.URI)

	if config.Username != "" && config.Password != "" {
		clientOptions.SetAuth(options.Credential{
			Username: config.Username,
			Password: config.Password,
		})
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	DB = client.Database(config.DBName)
	return nil
}

// Close closes the database connection
func Close() error {
	if DB == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return DB.Client().Disconnect(ctx)
}

// Helper function
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
