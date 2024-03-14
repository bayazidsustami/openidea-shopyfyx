package bank_account_service

import (
	"context"
	bank_account_model "openidea-shopyfyx/models/bank_account"
	bank_account_repository "openidea-shopyfyx/repository/bank_account"
	"openidea-shopyfyx/service/auth_service"

	"github.com/go-playground/validator/v10"
)

type BankAccountServiceImpl struct {
	Repository  bank_account_repository.BankAccountRepository
	Validator   *validator.Validate
	AuthService *auth_service.AuthService
}

func New(
	repository bank_account_repository.BankAccountRepository,
	validator *validator.Validate,
	authService auth_service.AuthService,
) BankAccountService {
	return &BankAccountServiceImpl{
		Repository:  repository,
		Validator:   validator,
		AuthService: authService,
	}
}

func (service *BankAccountServiceImpl) Create(context context.Context, request bank_account_model.BankAccountRequest) (*bank_account_model.BankAccountResponse, error) {
	err := service.Validator.Struct(request)
	if err != nil {
		return nil, err
	}

	bankAccount := bank_account_model.BankAccount{
		BankName:          request.BankName,
		BankAccountName:   request.BankAccountName,
		BankAccountNumber: request.BankAccountNumber,
	}
}

func (service *BankAccountServiceImpl) Get(context context.Context, request int) (*bank_account_model.BankAccountResponse, error) {

}
