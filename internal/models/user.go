package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system
type User struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  *string    `json:"-"`
	DOB       *time.Time `json:"dob"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// create a new user
func NewUser(fields map[string]interface{}) (*User, error) {

	// get default timestamp
	now := time.Now()

	// initializing new user
	newUser := &User{
		ID:        uuid.New(),
		Name:      fields["name"].(string),
		Username:  fields["username"].(string),
		Email:     fields["email"].(string),
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	// handle password field
	if password, ok := fields["password"].(string); ok {
		newUser.Password = &password
	}

	// handle dob field
	if bd, ok := fields["dob"].(time.Time); ok {
		newUser.DOB = &bd
	} else {
		return nil, fmt.Errorf("date of birth is invalid")
	}

	return newUser, nil

}
