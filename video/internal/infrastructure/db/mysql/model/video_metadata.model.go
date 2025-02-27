package model

type VideoMetadata struct {
	AbstractModel
	VideoID   string `json:"video_id" gorm:"index;not null"`
	Duration  int    `json:"duration,omitempty"`
	View      int    `json:"view,omitempty"`
	Like      int    `json:"like,omitempty"`
	Dislike   int    `json:"dislike,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}
