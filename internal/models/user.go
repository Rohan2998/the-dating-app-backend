package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system
type User struct {
	ID            uuid.UUID  `json:"id" db:"id"`
	Name          string     `json:"name" db:"name"`
	Username      string     `json:"username" db:"username"`
	Email         string     `json:"email" db:"email"`
	Password      *string    `json:"-" db:"password"`
	CountryAbbr   *string    `json:"countryAbbr" db:"country_abbr"`
	CountryCode   *string    `json:"countryCode" db:"country_code"`
	ContactNumber *string    `json:"contactNumber" db:"contact_number"`
	DOB           *time.Time `json:"dob" db:"dob"`
	CreatedAt     *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt     *time.Time `json:"updatedAt" db:"updated_at"`
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

	// handle country abbreviation
	if countryAbbr, ok := fields["countryAbbr"].(string); ok {
		newUser.CountryAbbr = &countryAbbr
	}

	// handle country code
	if countryCode, ok := fields["countryCode"].(string); ok {
		newUser.CountryCode = &countryCode
	}

	// handle country code
	if contactNumber, ok := fields["contactNumber"].(string); ok {
		newUser.ContactNumber = &contactNumber
	}

	// handle dob field
	if bd, ok := fields["dob"].(time.Time); ok {
		newUser.DOB = &bd
	} else {
		return nil, fmt.Errorf("date of birth is invalid")
	}

	return newUser, nil
}
