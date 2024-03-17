package auth_service

import (
	"context"
	user_model "openidea-shopyfyx/models/user"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
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
	signedToken, err := token.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		return user_model.User{}, fiber.NewError(fiber.StatusForbidden, "forbidden")
	}

	user.AccessToken = signedToken

	return user, nil
}

func (service *AuthServiceImpl) GetValidUser(ctx *fiber.Ctx) (user_model.User, error) {
	userInfo := ctx.Locals("userInfo").(*jwt.Token)
	claims := userInfo.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	userId := claims["user_id"].(float64)

	return user_model.User{
		UserId:   int(userId),
		Username: username,
	}, nil
}
