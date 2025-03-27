package model

import "gorm.io/gorm"

type Video struct {
	AbstractModel
	Title        string `json:"title" gorm:"type:varchar(255);not null"`
	Description  string `json:"description,omitempty"`
	IsSearchable bool   `json:"is_searchable,omitempty"`
	IsPublic     bool   `json:"visibility,omitempty"`
	VideoURL     string `json:"video_url,omitempty" gorm:"type:varchar(255);not null;uniqueIndex"`
	Bucket       string `json:"bucket" gorm:"type:varchar(255);not null"`
	ObjectKey    string `json:"object_key" gorm:"type:varchar(255);not null;uniqueIndex"`
	UploadedUser string `json:"uploaded_user" gorm:"type:varchar(255);not null"`
}

func (v *Video) AfterCreate(tx *gorm.DB) (err error) {
	videoMetadata := VideoMetadata{
		AbstractModel: AbstractModel{
			ID:        v.ID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
		},
		VideoID:   v.ID,
		Duration:  0,
		View:      0,
		Like:      0,
		Dislike:   0,
		Thumbnail: "",
	}
	if err := tx.Create(&videoMetadata).Error; err != nil {
		return err
	}

	return nil
}
