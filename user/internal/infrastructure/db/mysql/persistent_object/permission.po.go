package persistent_object

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(255);uniqueIndex;not null"`
	Description string `json:"description,omitempty"`
}

func (Permission) TableName() string {
	return "permissions"
}
