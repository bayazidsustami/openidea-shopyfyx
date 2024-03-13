package auth_service

import (
	"context"
	user_model "openidea-shopyfyx/models/user"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	ValidateToken(context context.Context, user user_model.User) (user_model.User, error)
	GetValidUser(ctx *fiber.Ctx) (user_model.User, error)
}
