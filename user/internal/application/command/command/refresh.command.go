package command

type RefreshTokenCommand struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenCommandResult struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	AccessTokenExpiresAt  int    `json:"access_token_expires_at"`
	RefreshTokenExpiresAt int    `json:"refresh_token_expires_at"`
}

func (*RefreshTokenCommand) CommandName() string {
	return "RefreshTokenCommand"
}
