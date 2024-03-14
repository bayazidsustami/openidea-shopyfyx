package bank_account_service

import (
	"context"
	bank_account_model "openidea-shopyfyx/models/bank_account"
)

type BankAccountService interface {
	Create(context context.Context, request bank_account_model.BankAccountRequest) (*bank_account_model.BankAccount, error)
	Get(context context.Context, request int) (*bank_account_model.BankAccountResponse, error)
}
