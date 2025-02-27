package model

type UserVideoMetadata struct {
	AbstractModel
	UserID    string `json:"user_id" gorm:"index;not null"`
	VideoID   string `json:"video_id" gorm:"index;not null"`
	Watched   bool   `json:"watched,omitempty"`
	WatchTime int    `json:"watch_time,omitempty"`
	Liked     *bool  `json:"liked,omitempty"` // nil: not rated, true: liked, false: disliked
}
