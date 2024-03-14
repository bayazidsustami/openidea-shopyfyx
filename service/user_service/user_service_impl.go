package user_service

import (
	"context"
	user_model "openidea-shopyfyx/models/user"
	user_repository "openidea-shopyfyx/repository/user"
	"openidea-shopyfyx/service/auth_service"
	"openidea-shopyfyx/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	Repository  user_repository.UserRepository
	Validator   *validator.Validate
	AuthService auth_service.AuthService
	DBPool      *pgxpool.Pool
}

func New(
	repository user_repository.UserRepository,
	validator *validator.Validate,
	authService auth_service.AuthService,
	dbPool *pgxpool.Pool,
) UserService {
	return &UserServiceImpl{
		Repository:  repository,
		Validator:   validator,
		AuthService: authService,
		DBPool:      dbPool,
	}
}

func (service *UserServiceImpl) Register(context context.Context, request user_model.UserRegisterRequest) (*user_model.UserResponse, error) {

	err := service.Validator.Struct(request)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	conn, err := service.DBPool.Acquire(context)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "something error")
	}
	defer conn.Release()

	tx, err := conn.Begin(context)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "something error")
	}
	defer utils.CommitOrRollback(context, tx)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), viper.GetInt("BCRYPT_SALT"))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "something error")
	}

	user := user_model.User{
		Username: request.Username,
		Password: string(hashedPassword),
		Name:     request.Name,
	}

	userResult, err := service.Repository.Register(context, tx, user)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusConflict, "user already registered")
	}

	validUser, err := service.AuthService.ValidateToken(context, userResult)
	if err != nil {
		tx.Rollback(context)
		return nil, fiber.NewError(fiber.StatusForbidden, err.Error())
	}

	return &user_model.UserResponse{
		Message: "success",
		Data: user_model.UserData{
			Username:    validUser.Username,
			Name:        validUser.Name,
			AccessToken: validUser.AccessToken,
		},
	}, nil
}

func (service *UserServiceImpl) Login(context context.Context, request user_model.UserLoginRequest) (*user_model.UserResponse, error) {
	err := service.Validator.Struct(request)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	conn, err := service.DBPool.Acquire(context)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "something error")
	}
	defer conn.Release()

	tx, err := conn.Begin(context)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "something error")
	}
	defer utils.CommitOrRollback(context, tx)

	user := user_model.User{
		Username: request.Username,
	}

	userResult, err := service.Repository.Login(context, tx, user)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(request.Password))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	validUser, err := service.AuthService.ValidateToken(context, userResult)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusForbidden, err.Error())
	}

	return &user_model.UserResponse{
		Message: "success",
		Data: user_model.UserData{
			Username:    validUser.Username,
			Name:        validUser.Name,
			AccessToken: validUser.AccessToken,
		},
	}, nil
}
