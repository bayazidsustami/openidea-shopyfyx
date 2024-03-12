package user_service

import (
	"context"
	user_model "openidea-shopyfyx/models/user"
)

type UserService interface {
	Register(context context.Context, request user_model.UserRegisterRequest) user_model.UserResponse
	Login(context context.Context)
}
