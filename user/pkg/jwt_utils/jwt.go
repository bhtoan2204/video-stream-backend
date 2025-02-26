package jwt_utils

import (
	"time"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/domain/entities"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *entities.User) (string, string, int64, int64, error) {
	accessSecret := []byte(global.Config.SecurityConfig.JWTAccessSecret)
	accessExpiration := time.Now().Add(time.Duration(global.Config.SecurityConfig.JWTAccessExpiration) * time.Second).Unix()
	refreshSecret := []byte(global.Config.SecurityConfig.JWTRefreshSecret)
	refreshExpiration := time.Now().Add(time.Duration(global.Config.SecurityConfig.JWTRefreshExpiration) * time.Second).Unix()

	// Access Token Claims
	accessClaims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      accessExpiration,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccessToken, err := accessToken.SignedString(accessSecret)

	if err != nil {
		return "", "", 0, 0, err
	}

	// Refresh Token Claims
	refreshClaims := jwt.MapClaims{
		"id":  user.ID,
		"exp": refreshExpiration,
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString(refreshSecret)
	if err != nil {
		return "", "", 0, 0, err
	}

	return signedAccessToken, signedRefreshToken, int64(accessExpiration), int64(refreshExpiration), nil
}

func VerifyToken(tokenString string, secret string) (*jwt.Token, jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})

	if err != nil || !parsedToken.Valid {
		return nil, nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, jwt.ErrInvalidKey
	}

	return parsedToken, claims, nil
}

func ExtractTokenClaims(tokenString string, secret string) (jwt.MapClaims, error) {
	_, claims, err := VerifyToken(tokenString, secret)
	return claims, err
}

func RefreshNewToken(user *entities.User, refreshTokenString string) (string, string, int64, int64, error) {
	refreshSecret := []byte(global.Config.SecurityConfig.JWTRefreshSecret)
	_, claims, err := VerifyToken(refreshTokenString, string(refreshSecret))
	if err != nil {
		return "", "", 0, 0, err
	}

	refreshExpiration := int64(claims["exp"].(float64))
	if time.Now().Unix() > int64(refreshExpiration) {
		return "", "", 0, 0, jwt.ErrTokenExpired
	}

	if err != nil {
		return "", "", 0, 0, err
	}

	accessToken, refreshToken, accessExpiration, refreshExpiration, err := GenerateToken(user)
	if err != nil {
		return "", "", 0, 0, err
	}

	return accessToken, refreshToken, accessExpiration, refreshExpiration, nil
}
