package bank_account_controller

import (
	bank_account_model "openidea-shopyfyx/models/bank_account"
	"openidea-shopyfyx/service/auth_service"
	"openidea-shopyfyx/service/bank_account_service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BankAccountController struct {
	BankAccountService bank_account_service.BankAccountService
	AuthService        auth_service.AuthService
}

func New(
	bankAccountService bank_account_service.BankAccountService,
	authService auth_service.AuthService,
) BankAccountController {

	return BankAccountController{
		BankAccountService: bankAccountService,
		AuthService:        authService,
	}
}

func (controller *BankAccountController) Create(ctx *fiber.Ctx) error {
	bankAccountRequest := new(bank_account_model.BankAccountRequest)

	err := ctx.BodyParser(bankAccountRequest)
	if err != nil {
		return err
	}

	user, err := controller.AuthService.GetValidUser(ctx)
	if err != nil {
		return err
	}

	err = controller.BankAccountService.Create(ctx.UserContext(), user, *bankAccountRequest)
	if err != nil {
		return err
	}

	return ctx.SendString("success")
}

func (controller *BankAccountController) GetAllByUserId(ctx *fiber.Ctx) error {
	user, err := controller.AuthService.GetValidUser(ctx)
	if err != nil {
		return err
	}

	response, err := controller.BankAccountService.GetAllByUserId(ctx.UserContext(), user)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

func (controller *BankAccountController) Update(ctx *fiber.Ctx) error {
	bankAccountRequest := new(bank_account_model.BankAccountRequest)

	err := ctx.BodyParser(bankAccountRequest)
	if err != nil {
		return err
	}

	user, err := controller.AuthService.GetValidUser(ctx)
	if err != nil {
		return err
	}

	err = controller.BankAccountService.Update(ctx.UserContext(), user, *bankAccountRequest)
	if err != nil {
		return err
	}

	return ctx.SendString("success")
}

func (controller *BankAccountController) Delete(ctx *fiber.Ctx) error {
	bankAccountIdString := ctx.Params("bankAccountId")

	bankAccountId, err := strconv.Atoi(bankAccountIdString)
	if err != nil {
		return err
	}

	_, err = controller.AuthService.GetValidUser(ctx)
	if err != nil {
		return err
	}

	controller.BankAccountService.Delete(ctx.UserContext(), bankAccountId)

	return ctx.SendString("success")
}
