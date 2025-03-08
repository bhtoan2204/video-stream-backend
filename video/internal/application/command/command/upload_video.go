package command

type UploadVideoCommand struct {
}

func (*UploadVideoCommand) CommandName() string {
	return "UploadVideoCommand"
}
