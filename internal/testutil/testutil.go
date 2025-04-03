package testutil

import (
	"os"
	"testing"
)

// SetupTestEnv sets up test environment variables
func SetupTestEnv(t *testing.T) {
	if err := os.Setenv("APP_ENV", "test"); err != nil {
		t.Fatalf("Failed to set APP_ENV: %v", err)
	}
	if err := os.Setenv("SERVER_PORT", "8080"); err != nil {
		t.Fatalf("Failed to set SERVER_PORT: %v", err)
	}
	if err := os.Setenv("DB_HOST", "localhost"); err != nil {
		t.Fatalf("Failed to set DB_HOST: %v", err)
	}
	if err := os.Setenv("DB_PORT", "5432"); err != nil {
		t.Fatalf("Failed to set DB_PORT: %v", err)
	}
	if err := os.Setenv("DB_NAME", "testdb"); err != nil {
		t.Fatalf("Failed to set DB_NAME: %v", err)
	}
	if err := os.Setenv("DB_USER", "testuser"); err != nil {
		t.Fatalf("Failed to set DB_USER: %v", err)
	}
	if err := os.Setenv("DB_PASSWORD", "testpass"); err != nil {
		t.Fatalf("Failed to set DB_PASSWORD: %v", err)
	}
}

// TearDownTestEnv cleans up test environment variables
func TearDownTestEnv(t *testing.T) {
	if err := os.Unsetenv("APP_ENV"); err != nil {
		t.Errorf("Failed to unset APP_ENV: %v", err)
	}
	if err := os.Unsetenv("SERVER_PORT"); err != nil {
		t.Errorf("Failed to unset SERVER_PORT: %v", err)
	}
	if err := os.Unsetenv("DB_HOST"); err != nil {
		t.Errorf("Failed to unset DB_HOST: %v", err)
	}
	if err := os.Unsetenv("DB_PORT"); err != nil {
		t.Errorf("Failed to unset DB_PORT: %v", err)
	}
	if err := os.Unsetenv("DB_NAME"); err != nil {
		t.Errorf("Failed to unset DB_NAME: %v", err)
	}
	if err := os.Unsetenv("DB_USER"); err != nil {
		t.Errorf("Failed to unset DB_USER: %v", err)
	}
	if err := os.Unsetenv("DB_PASSWORD"); err != nil {
		t.Errorf("Failed to unset DB_PASSWORD: %v", err)
	}
}
