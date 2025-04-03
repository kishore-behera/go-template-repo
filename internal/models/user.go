package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the system
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"`
	FirstName string             `bson:"first_name" json:"first_name"`
	LastName  string             `bson:"last_name" json:"last_name"`
	Active    bool               `bson:"active" json:"active"`
	LastLogin time.Time          `bson:"last_login" json:"last_login"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

// BeforeCreate is a hook that runs before creating a new user
func (u *User) BeforeCreate() error {
	// Hash password before saving
	hashedPassword, err := hashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate is a hook that runs before updating a user
func (u *User) BeforeUpdate() error {
	// Hash password if it was changed
	hashedPassword, err := hashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	u.UpdatedAt = time.Now()
	return nil
}

// Helper function to hash passwords (implement this based on your security requirements)
func hashPassword(password string) (string, error) {
	// TODO: Implement proper password hashing
	return password, nil
}
