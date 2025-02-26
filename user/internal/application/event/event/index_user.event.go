package event

import "time"

type Status int

const (
	InActive Status = 0
	Active   Status = 1
	Deleted  Status = 2
)

type IndexUserEvent struct {
	ID           string     `json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	Phone        string     `json:"phone"`
	BirthDate    *time.Time `json:"birth_date"`
	Address      string     `json:"address"`
	PasswordHash string     `json:"password_hash"`
	PinCode      string     `json:"pin_code"`
	Status       Status     `json:"status"`
}

func (*IndexUserEvent) EventName() string {
	return "IndexUserEvent"
}
