package user_repository

import (
	"context"
	user_model "openidea-shopyfyx/models/user"
)

type UserRepository interface {
	Register(ctx context.Context, request user_model.User) user_model.User
	Login(ctx context.Context, request user_model.User) user_model.User
}
