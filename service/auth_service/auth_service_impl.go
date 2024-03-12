package auth_service

import (
	"context"
	user_model "openidea-shopyfyx/models/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthServiceImpl struct {
}

func New() AuthService {
	return &AuthServiceImpl{}
}

func (service *AuthServiceImpl) ValidateToken(context context.Context, user user_model.User) (user_model.User, error) {

	claims := jwt.MapClaims{
		"username": user.Username,
		"user_id":  user.UserId,
		"exp":      time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	// TODO : jangan lupa ganti JWTsecret key
	signedToken, err := token.SignedString([]byte("ini rahasia"))
	if err != nil {
		return user_model.User{}, err
	}

	user.AccessToken = signedToken

	return user, nil
}
