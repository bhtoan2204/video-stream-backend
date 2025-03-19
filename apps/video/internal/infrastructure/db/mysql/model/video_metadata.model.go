package model

type VideoMetadata struct {
	AbstractModel
	VideoID   string   `json:"video_id" gorm:"index;not null;uniqueIndex"`
	Duration  int      `json:"duration,omitempty"` // in seconds
	View      int      `json:"view,omitempty"`
	Like      int      `json:"like,omitempty"`
	Dislike   int      `json:"dislike,omitempty"`
	Thumbnail string   `json:"thumbnail,omitempty"`
	Tags      []string `json:"tags,omitempty" gorm:"many2many:video_tags;"`
	IsProcess bool     `json:"is_process,omitempty" gorm:"default:false"`
}
