package persistent_object

type Video struct {
	BasePO
	Title        string `json:"title" gorm:"type:varchar(255);not null"`
	Description  string `json:"description,omitempty"`
	IsSearchable bool   `json:"is_searchable,omitempty"`
	IsPublic     bool   `json:"visibility,omitempty"`
	VideoURL     string `json:"video_url,omitempty" gorm:"type:varchar(255);not null;uniqueIndex"`
	Bucket       string `json:"bucket" gorm:"type:varchar(255);not null"`
	ObjectKey    string `json:"object_key" gorm:"type:varchar(255);not null;uniqueIndex"`
	UploadedUser string `json:"uploaded_user" gorm:"type:varchar(255);not null"`
}

func (Video) TableName() string {
	return "videos"
}
