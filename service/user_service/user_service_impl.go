package user_service

import (
	"context"
	user_model "openidea-shopyfyx/models/user"
	user_repository "openidea-shopyfyx/repository/user"
	"openidea-shopyfyx/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	Repository user_repository.UserRepository
}

func New(repository user_repository.UserRepository) UserService {
	return &UserServiceImpl{
		Repository: repository,
	}
}

func (service *UserServiceImpl) Register(context context.Context, request user_model.UserRegisterRequest) user_model.UserResponse {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	utils.PanicErr(err)

	user := user_model.User{
		Username: request.Username,
		Password: string(hashedPassword),
		Name:     request.Name,
	}

	service.Repository.Register(context, user)

	return user_model.UserResponse{}
}

func (service *UserServiceImpl) Login(context context.Context) {
	//TODO implement logic
}
