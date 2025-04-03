package service

import (
	"testing"
)

// TestMain provides setup and teardown for all tests in the package
func TestMain(m *testing.M) {
	// Setup
	setup()

	// Run tests
	m.Run()

	// Teardown
	teardown()
}

func setup() {
	// Add any test setup here
	// For example: database connections, test data, etc.
}

func teardown() {
	// Add any cleanup here
	// For example: closing connections, removing test data, etc.
}
