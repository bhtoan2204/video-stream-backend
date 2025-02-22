package command

import common "github.com/bhtoan2204/user/internal/application/common/command"

type GetUserByIdCommand struct {
	ID uint `json:"id" binding:"required"`
}

type GetUserByIdCommandResult struct {
	Result *common.UserResult `json:"result"`
}

func (*GetUserByIdCommand) CommandName() string {
	return "GetUserByIdCommand"
}
