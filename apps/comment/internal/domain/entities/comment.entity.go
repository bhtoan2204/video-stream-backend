package entities

type Comment struct {
	AbstractModel
	VideoID    string `json:"video_id" gorm:"type:char(36);not null"`
	UserID     string `json:"user_id" gorm:"type:char(36);not null"`
	Content    string `json:"content" gorm:"type:text;not null"`
	ParentID   string `json:"parent_id" gorm:"type:char(36);default:null"`
	LikeCount  int    `json:"like_count" gorm:"type:int;default:0"`
	ReplyCount int    `json:"reply_count" gorm:"type:int;default:0"`
	Status     string `json:"status" gorm:"type:enum('active', 'inactive');default:'active'"`
}

func NewComment(videoID string, userID string, parentID string, content string) *Comment {
	return &Comment{
		VideoID:    videoID,
		UserID:     userID,
		Content:    content,
		ParentID:   parentID,
		LikeCount:  0,
		ReplyCount: 0,
		Status:     "active",
	}
}
