package payload

type ImageResizePayload struct {
	SourcePath string `json:"source_path"`
	TargetPath string `json:"target_path"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
}
