package model

import "time"

type ESUser struct {
	ESAbstractModel
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	BirthDate time.Time `json:"birth_date,omitempty"`
	Address   string    `json:"address,omitempty"`
	PinCode   string    `json:"pin_code,omitempty"`
	Status    int       `json:"status"`
	Roles     []string  `json:"roles,omitempty"`
}
