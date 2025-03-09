package command

type DeleteVideoCommand struct {
}

type DeleteVideoCommandResult struct {
}

func (*DeleteVideoCommand) CommandName() string {
	return "DeleteVideoCommand"
}