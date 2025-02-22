package command

type RefreshTokenCommand struct {
	RefreshToken string `json:"refresh_token"`
}

func (*RefreshTokenCommand) CommandName() string {
	return "RefreshTokenCommand"
}
