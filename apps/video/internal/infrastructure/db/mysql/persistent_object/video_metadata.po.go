package persistent_object

type VideoMetadata struct {
	BasePO
	VideoID   string `json:"video_id" gorm:"index;not null;"`
	Duration  int    `json:"duration,omitempty"`
	View      int    `json:"view,omitempty"`
	Like      int    `json:"like,omitempty"`
	Dislike   int    `json:"dislike,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
	Tags      []Tag  `json:"tags,omitempty" gorm:"many2many:video_tags;"`
	IsProcess bool   `json:"is_process,omitempty" gorm:"default:false"`
}

func (VideoMetadata) TableName() string {
	return "video_metadata"
}
