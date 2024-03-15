package bank_account_service

import (
	"context"
	bank_account_model "openidea-shopyfyx/models/bank_account"
	user_model "openidea-shopyfyx/models/user"
)

type BankAccountService interface {
	Create(ctx context.Context, user user_model.User, request bank_account_model.BankAccountRequest) error
	GetAllByUserId(ctx context.Context, user user_model.User) (bank_account_model.BankAccountsByUserIdResponse, error)
	Update(ctx context.Context, user user_model.User, bankAccountId int, request bank_account_model.BankAccountRequest) error
	Delete(ctx context.Context, request int)
}
