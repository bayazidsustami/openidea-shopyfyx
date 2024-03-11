package user_controller

import "github.com/gofiber/fiber/v2"

type UserController struct {
}

func New() UserController {
	return UserController{}
}

func (controller *UserController) Register(ctx *fiber.Ctx) error {
	//TODO : parsing request body and do register

	return nil
}

func (controller *UserController) Login(ctx *fiber.Ctx) error {
	// TODO : parsing request body and do login

	return nil
}
