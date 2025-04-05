package model

type Transaction struct {
	AbstractModel
	UserId          string  `json:"user_id" gorm:"type:varchar(255);not null"`
	Amount          float64 `json:"amount" gorm:"type:decimal(10,2);not null"`
	CurrencyId      string  `json:"currency_id" gorm:"type:varchar(3);not null"`
	TransactionType string  `json:"transaction_type" gorm:"type:varchar(50);not null"` // donation, refund,
	Status          string  `json:"status" gorm:"type:varchar(50);not null"`
	ReferenceId     string  `json:"reference_id" gorm:"type:varchar(255);not null"`
	PaymentMethod   string  `json:"payment_method" gorm:"type:varchar(50);not null"`
}
