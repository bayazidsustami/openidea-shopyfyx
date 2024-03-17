package bank_account_service

import (
	"context"
	bank_account_model "openidea-shopyfyx/models/bank_account"
	user_model "openidea-shopyfyx/models/user"
	bank_account_repository "openidea-shopyfyx/repository/bank_account"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type BankAccountServiceImpl struct {
	BankAccountRepository bank_account_repository.BankAccountRepository
	Validator             *validator.Validate
}

func New(
	bankAccountRepository bank_account_repository.BankAccountRepository,
	validator *validator.Validate,
) BankAccountService {
	return &BankAccountServiceImpl{
		BankAccountRepository: bankAccountRepository,
		Validator:             validator,
	}
}

func (service *BankAccountServiceImpl) Create(ctx context.Context, user user_model.User, request bank_account_model.BankAccountRequest) error {
	err := service.Validator.Struct(request)
	if err != nil {
		return err
	}

	bankAccount := bank_account_model.BankAccount{
		BankName:          request.BankName,
		BankAccountName:   request.BankAccountName,
		BankAccountNumber: request.BankAccountNumber,
		UserId:            user.UserId,
	}

	service.BankAccountRepository.Create(ctx, bankAccount)

	return nil
}

func (service *BankAccountServiceImpl) GetAllByUserId(ctx context.Context, user user_model.User) (bank_account_model.BankAccountsByUserIdResponse, error) {
	bankAccounts, err := service.BankAccountRepository.GetAllByUserId(ctx, user.UserId)
	if err != nil {
		return bank_account_model.BankAccountsByUserIdResponse{}, err
	}

	var bankAccountsByUserIdResponse bank_account_model.BankAccountsByUserIdResponse
	for _, bankAccount := range bankAccounts {
		bankAccountData := bank_account_model.BankAccountData{
			BankAccountId:     bankAccount.BankAccountId,
			BankName:          bankAccount.BankName,
			BankAccountName:   bankAccount.BankAccountName,
			BankAccountNumber: bankAccount.BankAccountNumber,
		}

		bankAccountsByUserIdResponse.Data = append(bankAccountsByUserIdResponse.Data, bankAccountData)
	}
	bankAccountsByUserIdResponse.Message = "success"

	return bankAccountsByUserIdResponse, nil
}

func (service *BankAccountServiceImpl) Update(ctx context.Context, user user_model.User, bankAccountId int, request bank_account_model.BankAccountRequest) error {
	err := service.Validator.Struct(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid input")
	}

	bankAccount := bank_account_model.BankAccount{
		BankAccountId:     bankAccountId,
		BankName:          request.BankName,
		BankAccountName:   request.BankAccountName,
		BankAccountNumber: request.BankAccountNumber,
		UserId:            user.UserId,
	}

	service.BankAccountRepository.Update(ctx, bankAccount)

	return nil
}

func (service *BankAccountServiceImpl) Delete(ctx context.Context, request int) {
	service.BankAccountRepository.Delete(ctx, request)
}
