package persistent_object

type Comment struct {
	BasePO
	VideoID  string `json:"video_id" gorm:"type:char(36);not null"`
	UserID   string `json:"user_id" gorm:"type:char(36);not null"`
	Content  string `json:"content" gorm:"type:text;not null"`
	ParentID string `json:"parent_id" gorm:"type:char(36);default:null"`
	Status   string `json:"status" gorm:"type:enum('active', 'inactive');default:'active'"`
}

func (Comment) TableName() string {
	return "comments"
}
