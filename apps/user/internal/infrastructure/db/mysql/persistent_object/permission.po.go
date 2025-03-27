package persistent_object

type Permission struct {
	BasePO
	Name        string `json:"name" gorm:"type:varchar(255);uniqueIndex;not null"`
	Description string `json:"description,omitempty"`
}

func (Permission) TableName() string {
	return "permissions"
}
