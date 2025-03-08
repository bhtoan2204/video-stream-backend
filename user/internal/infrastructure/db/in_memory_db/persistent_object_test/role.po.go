package persistent_object_test

type Role struct {
	BasePO
	Name        string        `json:"name" gorm:"type:varchar(255);uniqueIndex;not null"`
	Permissions []*Permission `json:"permissions" gorm:"many2many:role_permissions;"`
}

func (Role) TableName() string {
	return "roles"
}
