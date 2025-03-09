package command

type GetVideoByUserIdCommand struct {
}

type GetVideoByUserIdCommandResult struct {
}

func (*GetVideoByUserIdCommand) CommandName() string {
	return "GetVideoByUserIdCommand"
}