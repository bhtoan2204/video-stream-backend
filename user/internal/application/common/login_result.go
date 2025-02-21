package common

type LoginResult struct {
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  int64
	RefreshTokenExpiresAt int64
}
