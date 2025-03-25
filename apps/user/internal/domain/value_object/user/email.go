package value_object

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

var emailRegex = regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)

func NewEmail(email string) (*Email, error) {
	if !emailRegex.MatchString(email) {
		return nil, errors.New("invalid email format")
	}

	return &Email{value: email}, nil
}

func (e *Email) Value() string {
	return e.value
}
