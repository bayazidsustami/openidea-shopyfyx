package user_service

import "context"

type UserService interface {
	Register(context *context.Context)
	Login(context *context.Context)
}
