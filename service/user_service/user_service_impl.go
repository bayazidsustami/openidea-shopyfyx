package user_service

import (
	"context"
	user_model "openidea-shopyfyx/models/user"
	user_repository "openidea-shopyfyx/repository/user"
	"openidea-shopyfyx/utils"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	Repository user_repository.UserRepository
	Validator  *validator.Validate
}

func New(repository user_repository.UserRepository, validator *validator.Validate) UserService {
	return &UserServiceImpl{
		Repository: repository,
		Validator:  validator,
	}
}

func (service *UserServiceImpl) Register(context context.Context, request user_model.UserRegisterRequest) (*user_model.UserResponse, error) {

	err := service.Validator.Struct(request)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	utils.PanicErr(err)

	user := user_model.User{
		Username: request.Username,
		Password: string(hashedPassword),
		Name:     request.Name,
	}

	userResult := service.Repository.Register(context, user)

	return &user_model.UserResponse{
		Message: "success",
		Data: user_model.UserData{
			Username:    userResult.Username,
			Name:        userResult.Name,
			AccessToken: "",
		},
	}, nil
}

func (service *UserServiceImpl) Login(context context.Context, request user_model.UserLoginRequest) (*user_model.UserResponse, error) {
	err := service.Validator.Struct(request)
	if err != nil {
		return nil, err
	}

	user := user_model.User{
		Username: request.Username,
	}

	userResult := service.Repository.Login(context, user)

	err = bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(request.Password))
	if err != nil {
		return nil, err
	}

	return &user_model.UserResponse{
		Message: "success",
		Data: user_model.UserData{
			Username:    userResult.Username,
			Name:        userResult.Name,
			AccessToken: "",
		},
	}, nil
}
