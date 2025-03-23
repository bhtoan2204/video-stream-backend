package model

type Tag struct {
	AbstractModel
	Name string `json:"name" gorm:"unique;not null"`
}
