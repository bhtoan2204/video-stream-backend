package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status int

const (
	InActive Status = 0
	Active   Status = 1
	Deleted  Status = 2
)

type User struct {
	AbstractModel
	Username     string        `json:"username" gorm:"uniqueIndex;not null"`
	Email        string        `json:"email" gorm:"uniqueIndex;not null"`
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

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	defaultSettings := UserSettings{
		AbstractModel: AbstractModel{
			ID:        uuid.NewString(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
		UserID:               u.ID,
		Language:             "en",
		Theme:                ThemeLight,
		NotificationsEnabled: true,
		Privacy: &PrivacySettings{
			ShowEmail:       false,
			ShowSubscribers: true,
		},
	}
	if err := tx.Create(&defaultSettings).Error; err != nil {
		return err
	}

	return nil
}
