package model

import "time"

type ESUser struct {
	ID           uint           `json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    *time.Time     `json:"deleted_at,omitempty"`
	Username     string         `json:"username"`
	Email        string         `json:"email"`
	FirstName    string         `json:"first_name,omitempty"`
	LastName     string         `json:"last_name,omitempty"`
	Phone        string         `json:"phone,omitempty"`
	BirthDate    *time.Time     `json:"birth_date,omitempty"`
	Address      string         `json:"address,omitempty"`
	PasswordHash string         `json:"password_hash"`
	PinCode      string         `json:"pin_code,omitempty"`
	Status       int            `json:"status"` // Ví dụ: 0: InActive, 1: Active, 2: Deleted
	Roles        []ESRole       `json:"roles,omitempty"`
	Settings     *ESUserSetting `json:"settings,omitempty"`
}
