package bank_account_repository

import (
	"context"
	bank_account_model "openidea-shopyfyx/models/bank_account"
)

type BankAccountRepository interface {
	Create(ctx context.Context, request bank_account_model.BankAccount) bank_account_model.BankAccount
	GetAllByUserId(ctx context.Context, request int) ([]bank_account_model.BankAccount, error)
	Update(ctx context.Context, request bank_account_model.BankAccount) bank_account_model.BankAccount
	Delete(ctx context.Context, request int)
}
