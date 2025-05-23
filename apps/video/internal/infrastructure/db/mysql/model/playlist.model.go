package model

type Visibility int

const (
	Private  Visibility = 0
	Unlisted Visibility = 1
	Public   Visibility = 2
)

type Playlist struct {
	AbstractModel
	Title      string     `json:"title" gorm:"type:varchar(255);not null"`
	Visibility Visibility `json:"visibility" gorm:"type:varchar(255);not null"`
	UserID     string     `json:"user_id" gorm:"type:varchar(255);not null"`
}
