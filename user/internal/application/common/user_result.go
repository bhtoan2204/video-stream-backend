package common

import "github.com/bhtoan2204/user/internal/domain/entities"

type UserResult struct {
	ID        uint
	Username  string
	Email     string
	FirstName string
	LastName  string
	Phone     string
	BirthDate string
	Address   string
	Settings  *entities.UserSettings
	Role      string
}
