package entities

type Video struct {
	AbstractModel
	Title        string `json:"title"`
	Description  string `json:"description,omitempty"`
	IsSearchable bool   `json:"is_searchable,omitempty"`
	IsPublic     bool   `json:"visibility,omitempty"`
	VideoURL     string `json:"video_url,omitempty"`
	Bucket       string `json:"bucket"`
	ObjectKey    string `json:"object_key"`
	UploadedUser string `json:"uploaded_user"`
}
