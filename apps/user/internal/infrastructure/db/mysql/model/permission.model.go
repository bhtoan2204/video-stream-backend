package model

type Permission struct {
	AbstractModel
	Name        string `json:"name" gorm:"uniqueIndex;not null"` // Unique constraint to prevent duplicates
	Description string `json:"description,omitempty"`
}
