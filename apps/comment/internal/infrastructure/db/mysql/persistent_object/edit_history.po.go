package persistent_object

type EditHistory struct {
	BasePO
	CommentID string `json:"comment_id" gorm:"type:char(36);not null"`
	UserID    string `json:"user_id" gorm:"type:char(36);not null"`
	Content   string `json:"content" gorm:"type:text;not null"`
}

func (EditHistory) TableName() string {
	return "edit_histories"
}
