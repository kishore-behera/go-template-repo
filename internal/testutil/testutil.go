package testutil

import (
	"os"
	"testing"
)

// TestEnv sets up test environment variables
func TestEnv(t *testing.T) {
	// Set test environment variables
	os.Setenv("APP_ENV", "test")
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "testdb")
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_PASSWORD", "testpass")
}

// CleanupEnv cleans up test environment variables
func CleanupEnv(t *testing.T) {
	// Clean up environment variables
	os.Unsetenv("APP_ENV")
	os.Unsetenv("SERVER_PORT")
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("DB_NAME")
	os.Unsetenv("DB_USER")
	os.Unsetenv("DB_PASSWORD")
}
