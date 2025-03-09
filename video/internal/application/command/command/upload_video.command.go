package command

type UploadVideoCommand struct {
}

type UploadVideoCommandResult struct {
}

func (*UploadVideoCommand) CommandName() string {
	return "UploadVideoCommand"
}
