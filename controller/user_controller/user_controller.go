package user_controller

import (
	user_model "openidea-shopyfyx/models/user"
	"openidea-shopyfyx/service/user_service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Service user_service.UserService
}

func New(service user_service.UserService) UserController {
	return UserController{
		Service: service,
	}
}

func (controller *UserController) Register(ctx *fiber.Ctx) error {
	userRequest := new(user_model.UserRegisterRequest)

	err := ctx.BodyParser(userRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request")
	}

	result, err := controller.Service.Register(ctx.UserContext(), *userRequest)
	if err != nil {
		return err
	}

	ctx.Status(201)
	return ctx.JSON(result)
}

func (controller *UserController) Login(ctx *fiber.Ctx) error {
	userRequest := new(user_model.UserLoginRequest)
	err := ctx.BodyParser(userRequest)
	if err != nil {
		return err
	}

	result, err := controller.Service.Login(ctx.UserContext(), *userRequest)
	if err != nil {
		return err
	}

	return ctx.JSON(result)
}
