package model

import (
	"encoding/json"
	"time"
)

type ESUser struct {
	ID           string         `json:"id"`
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
	Status       int            `json:"status"`
	Roles        []ESRole       `json:"roles,omitempty"`
	Settings     *ESUserSetting `json:"settings,omitempty"`
}

func (u *ESUser) UnmarshalJSON(data []byte) error {
	type Alias ESUser
	aux := &struct {
		CreatedAt int64  `json:"created_at"`
		UpdatedAt int64  `json:"updated_at"`
		BirthDate int64  `json:"birth_date,omitempty"`
		DeletedAt *int64 `json:"deleted_at,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	u.CreatedAt = time.UnixMilli(aux.CreatedAt)
	u.UpdatedAt = time.UnixMilli(aux.UpdatedAt)
	if aux.BirthDate != 0 {
		t := time.UnixMilli(aux.BirthDate)
		u.BirthDate = &t
	}
	if aux.DeletedAt != nil && *aux.DeletedAt != 0 {
		t := time.UnixMilli(*aux.DeletedAt)
		u.DeletedAt = &t
	}

	return nil
}

func (u ESUser) MarshalJSON() ([]byte, error) {
	type Alias ESUser
	aux := &struct {
		CreatedAt int64  `json:"created_at"`
		UpdatedAt int64  `json:"updated_at"`
		BirthDate int64  `json:"birth_date,omitempty"`
		DeletedAt *int64 `json:"deleted_at,omitempty"`
		*Alias
	}{
		CreatedAt: u.CreatedAt.UnixMilli(),
		UpdatedAt: u.UpdatedAt.UnixMilli(),
		Alias:     (*Alias)(&u),
	}
	if u.BirthDate != nil {
		aux.BirthDate = u.BirthDate.UnixMilli()
	}
	if u.DeletedAt != nil {
		ts := u.DeletedAt.UnixMilli()
		aux.DeletedAt = &ts
	}
	return json.Marshal(aux)
}
