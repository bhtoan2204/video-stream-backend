package model

type Video struct {
	AbstractModel
	Title        string `json:"title" gorm:"type:varchar(255);not null"`
	Description  string `json:"description,omitempty"`
	IsSearchable bool   `json:"is_searchable,omitempty"`
	IsPublic     bool   `json:"visibility,omitempty"`
	VideoURL     string `json:"video_url,omitempty"`
	Bucket       string `json:"bucket" gorm:"type:varchar(255);not null"`
	ObjectKey    string `json:"object_key" gorm:"type:varchar(255);not null"`
	UploadedUser string `json:"uploaded_user" gorm:"type:varchar(255);not null"`
}
