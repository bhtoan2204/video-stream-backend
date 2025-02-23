package common

import "github.com/bhtoan2204/user/internal/domain/entities"

type UserResult struct {
	ID        uint                   `json:"id"`
	Username  string                 `json:"username"`
	Email     string                 `json:"email"`
	FirstName string                 `json:"first_name"`
	LastName  string                 `json:"last_name"`
	Phone     string                 `json:"phone"`
	BirthDate string                 `json:"birth_date"`
	Address   string                 `json:"address"`
	Settings  *entities.UserSettings `json:"settings"`
	Role      string                 `json:"role"`
}
