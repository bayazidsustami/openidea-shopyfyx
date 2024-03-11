package user_controller

import (
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
	//TODO : parsing request body and do register

	return nil
}

func (controller *UserController) Login(ctx *fiber.Ctx) error {
	// TODO : parsing request body and do login

	return nil
}