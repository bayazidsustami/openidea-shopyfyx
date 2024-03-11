package user

import "github.com/gofiber/fiber/v2"

type UserController struct {
}

func (controller *UserController) register(ctx *fiber.Ctx) error {
	//TODO : parsing request body and do register

	return nil
}

func (controller *UserController) login(ctx *fiber.Ctx) error {
	// TODO : parsing request body and do login

	return nil
}
