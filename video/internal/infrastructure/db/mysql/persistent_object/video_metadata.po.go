package persistent_object

type VideoMetadata struct {
	BasePO
	VideoID   string `json:"video_id" gorm:"index;not null;uniqueIndex"`
	Duration  int    `json:"duration,omitempty"`
	View      int    `json:"view,omitempty"`
	Like      int    `json:"like,omitempty"`
	Dislike   int    `json:"dislike,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

func (VideoMetadata) TableName() string {
	return "video_metadata"
}
