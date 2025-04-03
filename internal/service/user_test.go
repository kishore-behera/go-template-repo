package service

import (
	"context"
	"testing"

	"github.com/kbehera/go-template-repo/internal/model"
	"github.com/kbehera/go-template-repo/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Create(t *testing.T) {
	// Setup test environment
	testutil.TestEnv(t)
	defer testutil.CleanupEnv(t)

	// Arrange
	svc := NewUserService()
	ctx := context.Background()
	user := &model.User{
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
	}

	// Act
	result, err := svc.Create(ctx, user)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.Email, result.Email)
	assert.Equal(t, user.FirstName, result.FirstName)
	assert.Equal(t, user.LastName, result.LastName)
	assert.NotZero(t, result.CreatedAt)
	assert.NotZero(t, result.UpdatedAt)
}

func TestUserService_Create_InvalidEmail(t *testing.T) {
	// Setup test environment
	testutil.TestEnv(t)
	defer testutil.CleanupEnv(t)

	// Arrange
	svc := NewUserService()
	ctx := context.Background()
	user := &model.User{
		Email:     "",
		FirstName: "Test",
		LastName:  "User",
	}

	// Act
	result, err := svc.Create(ctx, user)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "email is required", err.Error())
}
