package value_object

import (
	"errors"
	"unicode"

	"github.com/bhtoan2204/user/pkg/encrypt_password"
)

type Password struct {
	hashedPassword string
}

func NewPassword(password string) (Password, error) {
	if len(password) < 8 {
		return Password{}, errors.New("password must be at least 8 characters")
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return Password{}, errors.New("password must contain uppercase, lowercase, number, and special character")
	}

	return Password{hashedPassword: password}, nil
}

func (p Password) Hash() string {
	hashedPassword, _ := encrypt_password.HashPassword(p.hashedPassword)
	return hashedPassword
}
