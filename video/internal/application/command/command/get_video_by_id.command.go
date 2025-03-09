package command

type GetVideoByIdCommand struct {
}

type GetVideoByIdCommandResult struct {
}

func (*GetVideoByIdCommand) CommandName() string {
	return "GetVideoByIdCommand"
}