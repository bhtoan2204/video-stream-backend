package persistent_object_test

type Permission struct {
	BasePO
	Name        string `json:"name" gorm:"type:varchar(255);uniqueIndex;not null"`
	Description string `json:"description,omitempty"`
}

func (Permission) TableName() string {
	return "permissions"
}
