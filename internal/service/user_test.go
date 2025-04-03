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
	testutil.SetupTestEnv(t)
	defer testutil.TearDownTestEnv(t)

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
	testutil.SetupTestEnv(t)
	defer testutil.TearDownTestEnv(t)

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

func TestUserService_Get(t *testing.T) {
	// Setup test environment
	testutil.SetupTestEnv(t)
	defer testutil.TearDownTestEnv(t)

	// Arrange
	svc := NewUserService()
	ctx := context.Background()
	id := "test-id"

	// Act
	result, err := svc.Get(ctx, id)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, id, result.ID)
	assert.Equal(t, "example@example.com", result.Email)
	assert.Equal(t, "John", result.FirstName)
	assert.Equal(t, "Doe", result.LastName)
}

func TestUserService_Get_InvalidID(t *testing.T) {
	// Setup test environment
	testutil.SetupTestEnv(t)
	defer testutil.TearDownTestEnv(t)

	// Arrange
	svc := NewUserService()
	ctx := context.Background()
	id := ""

	// Act
	result, err := svc.Get(ctx, id)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "id is required", err.Error())
}

func TestUpdateUser(t *testing.T) {
	// Setup test environment
	testutil.SetupTestEnv(t)
	defer testutil.TearDownTestEnv(t)

	// Test implementation...
}

func TestDeleteUser(t *testing.T) {
	// Setup test environment
	testutil.SetupTestEnv(t)
	defer testutil.TearDownTestEnv(t)

	// Test implementation...
}
