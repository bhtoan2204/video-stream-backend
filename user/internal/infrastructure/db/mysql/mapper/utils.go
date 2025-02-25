package mapper

import (
	"time"

	"gorm.io/gorm"
)

func deletedAtToTimePointer(deletedAt gorm.DeletedAt) *time.Time {
	if deletedAt.Valid {
		return &deletedAt.Time
	}
	return nil
}

// Convert *time.Time to gorm.DeletedAt
func timePointerToDeletedAt(t *time.Time) gorm.DeletedAt {
	if t != nil {
		return gorm.DeletedAt{Time: *t, Valid: true}
	}
	return gorm.DeletedAt{Valid: false}
}
