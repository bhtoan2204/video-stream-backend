package global

type TaskType struct {
	TypeEmailDelivery    string
	TypeImageResize      string
	TypeVideoTranscoding string
}

var TaskTypeInstance = TaskType{
	TypeEmailDelivery:    "email:deliver",
	TypeImageResize:      "image:resize",
	TypeVideoTranscoding: "video:transcoding",
}
