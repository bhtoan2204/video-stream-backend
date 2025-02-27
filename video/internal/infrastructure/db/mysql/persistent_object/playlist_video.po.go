package persistent_object

type PlaylistVideo struct {
	BasePO
	PlaylistID string `json:"playlist_id" gorm:"index;not null"`
	VideoID    string `json:"video_id" gorm:"index;not null"`
}

func (PlaylistVideo) TableName() string {
	return "playlist_videos"
}
