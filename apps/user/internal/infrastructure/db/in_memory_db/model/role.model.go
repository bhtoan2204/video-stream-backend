package model_test

type Role struct {
	AbstractModel
	Name        string        `json:"name" gorm:"uniqueIndex;not null"`
	Permissions []*Permission `json:"permissions" gorm:"many2many:role_permissions;"`
}
