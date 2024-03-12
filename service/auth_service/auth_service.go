package auth_service

import (
	"context"
	user_model "openidea-shopyfyx/models/user"
)

type AuthService interface {
	ValidateToken(context context.Context, user user_model.User) (user_model.User, error)
}
