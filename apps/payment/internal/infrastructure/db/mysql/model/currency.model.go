package model

type Currency struct {
	AbstractModel
	Code   string `json:"code" gorm:"primaryKey;type:varchar(3);not null;unique"`
	Name   string `json:"name" gorm:"type:varchar(255);not null"`
	Symbol string `json:"symbol" gorm:"type:varchar(10);not null"`
}
