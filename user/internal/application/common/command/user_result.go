package common

import (
	"time"

	"github.com/bhtoan2204/user/internal/domain/entities"
)

type UserResult struct {
	ID        string
	Username  string
	Email     string
	FirstName string
	LastName  string
	Phone     string
	BirthDate *time.Time
	Address   string
	Settings  *entities.UserSettings
	Role      string
}
