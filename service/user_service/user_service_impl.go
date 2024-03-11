package user_service

import "context"

type UserServiceImpl struct {
}

func New() UserService {
	return &UserServiceImpl{}
}

func (service *UserServiceImpl) Register(context *context.Context) {
	//TODO implement logic
}

func (service *UserServiceImpl) Login(context *context.Context) {
	//TODO implement logic
}
