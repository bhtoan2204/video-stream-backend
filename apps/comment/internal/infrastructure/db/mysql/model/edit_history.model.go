package model

type EditHistory struct {
	AbstractModel
	CommentID string `json:"comment_id" gorm:"type:char(36);not null"`
	UserID    string `json:"user_id" gorm:"type:char(36);not null"`
	Content   string `json:"content" gorm:"type:text;not null"`
}
