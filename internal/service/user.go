package service

import (
	"context"
	"errors"

	"github.com/kbehera/go-template-repo/internal/model"
)

// UserService handles user-related business logic
type UserService struct {
	// Add dependencies here (e.g., repositories)
}

// NewUserService creates a new user service
func NewUserService() *UserService {
	return &UserService{}
}

// Create creates a new user
func (s *UserService) Create(ctx context.Context, user *model.User) (*model.User, error) {
	// Validate user data
	if user.Email == "" {
		return nil, errors.New("email is required")
	}

	// Create new user with timestamps
	newUser := model.NewUser(user.Email, user.FirstName, user.LastName)

	// TODO: Add repository call to create user
	// user, err := h.userService.Create(r.Context(), &user)
	// if err != nil {
	//     http.Error(w, err.Error(), http.StatusInternalServerError)
	//     return
	// }

	return newUser, nil
}

// Get retrieves a user by ID
func (s *UserService) Get(ctx context.Context, id string) (*model.User, error) {
	// TODO: Add repository call to get user
	if id == "" {
		return nil, errors.New("id is required")
	}

	return &model.User{
		ID:        id,
		Email:     "example@example.com",
		FirstName: "John",
		LastName:  "Doe",
	}, nil
}
