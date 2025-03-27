package persistent_object

type Tag struct {
	BasePO
	Name   string          `json:"name" gorm:"unique;not null"`
	Videos []VideoMetadata `json:"-" gorm:"many2many:video_tags;"`
}

func (Tag) TableName() string {
	return "tags"
}
