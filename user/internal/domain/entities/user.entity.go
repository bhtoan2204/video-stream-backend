package entities

import (
	"errors"
	"time"
)

type Status int

const (
	InActive Status = 0
	Active   Status = 1
	Deleted  Status = 2
)

type User struct {
	AbstractModel
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	FirstName    string     `json:"first_name,omitempty"`
	LastName     string     `json:"last_name,omitempty"`
	Phone        string     `json:"phone,omitempty"`
	BirthDate    *time.Time `json:"birth_date,omitempty"`
	Address      string     `json:"address,omitempty"`
	PasswordHash string     `json:"password_hash"`
	PinCode      string     `json:"pin_code,omitempty"`
	Status       Status     `json:"status"` // Default to Active
	Roles        []*Role    `json:"roles,omitempty"`
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username cannot be empty")
	}
	if u.FirstName == "" || u.LastName == "" {
		return errors.New("first name and last name cannot be empty")
	}
	if u.BirthDate == nil {
		return errors.New("birth date is required")
	}

	return nil
}
