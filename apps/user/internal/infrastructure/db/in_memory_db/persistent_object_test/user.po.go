package persistent_object_test

import (
	"time"
)

type Status int

const (
	InActive Status = 0
	Active   Status = 1
	Deleted  Status = 2
)

type User struct {
	BasePO
	Username     string        `json:"username" gorm:"type:varchar(255);uniqueIndex;not null"`
	Email        string        `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	FirstName    string        `json:"first_name,omitempty"`
	LastName     string        `json:"last_name,omitempty"`
	Phone        string        `json:"phone,omitempty"`
	BirthDate    *time.Time    `json:"birth_date,omitempty"`
	Address      string        `json:"address,omitempty"`
	PasswordHash string        `json:"password_hash" gorm:"not null"`
	PinCode      string        `json:"pin_code,omitempty"`
	Status       Status        `json:"status" gorm:"default:1"` // Default to Active
	Roles        []*Role       `json:"roles" gorm:"many2many:user_roles;"`
	Settings     *UserSettings `json:"settings,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

func (User) TableName() string {
	return "users"
}
