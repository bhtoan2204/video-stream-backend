package value_object

import (
	"time"

	"github.com/bhtoan2204/user/global"
	"go.uber.org/zap"
)

type BirthDate struct {
	value *time.Time
}

func NewBirthDate(birthDate string) (*BirthDate, error) {
	parsedDate, err := time.Parse(time.DateOnly, birthDate)
	if err != nil {
		global.Logger.Error("Failed to parse birth date ", zap.Error(err))
		return nil, err
	}
	return &BirthDate{value: &parsedDate}, nil
}

func (b BirthDate) Value() *time.Time {
	return b.value
}
