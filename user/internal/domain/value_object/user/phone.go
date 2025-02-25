package value_object

import (
	"errors"
	"regexp"
)

type Phone struct {
	value string
}

var phoneRegex = regexp.MustCompile(`^\+?[0-9]{9,15}$`)

func NewPhone(phone string) (Phone, error) {
	if !phoneRegex.MatchString(phone) {
		return Phone{}, errors.New("invalid phone number format")
	}
	return Phone{value: phone}, nil
}

func (p Phone) Value() string {
	return p.value
}
