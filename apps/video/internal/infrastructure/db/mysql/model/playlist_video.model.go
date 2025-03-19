package model

type PlaylistVideo struct {
	AbstractModel
	PlaylistID string `json:"playlist_id" gorm:"index;not null"`
	VideoID    string `json:"video_id" gorm:"index;not null"`
}
