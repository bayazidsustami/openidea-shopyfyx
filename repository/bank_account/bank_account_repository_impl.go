package bank_account_repository

import (
	"context"
	"log"
	bank_account_model "openidea-shopyfyx/models/bank_account"
	"openidea-shopyfyx/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BankAccountRepositoryImpl struct {
	DBPool *pgxpool.Pool
}

func New(DBPool *pgxpool.Pool) BankAccountRepository {
	return &BankAccountRepositoryImpl{
		DBPool: DBPool,
	}
}

func (repo *BankAccountRepositoryImpl) Create(ctx context.Context, bank_account bank_account_model.BankAccount) bank_account_model.BankAccount {
	conn, err := repo.DBPool.Acquire(ctx)
	utils.PanicErr(err)
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	utils.PanicErr(err)
	defer utils.CommitOrRollback(ctx, tx)

	var bankAccountId int
	SQL_INSERT := "INSERT INTO bank_accounts(bank_name, bank_account_name, bank_account_number, user_id) values ($1, $2, $3, $4) RETURNING bank_account_id"
	err = tx.QueryRow(ctx, SQL_INSERT, bank_account.BankName, bank_account.BankAccountName, bank_account.BankAccountNumber, bank_account.UserId).Scan(&bankAccountId)
	log.Print(err)

	utils.PanicErr(err)

	bank_account.BankAccountId = bankAccountId
	return bank_account
}

func (repo *BankAccountRepositoryImpl) GetAllByUserId(ctx context.Context, user_id int) ([]bank_account_model.BankAccount, error) {
	conn, err := repo.DBPool.Acquire(ctx)
	utils.PanicErr(err)
	defer conn.Release()

	SQL_GETALL := "SELECT bank_account_id, bank_name, bank_account_name, bank_account_number from bank_accounts WHERE user_id=$1"
	rows, err := conn.Query(ctx, SQL_GETALL, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bankAccounts []bank_account_model.BankAccount
	for rows.Next() {
		bankAccount := bank_account_model.BankAccount{}
		err := rows.Scan(
			&bankAccount.BankAccountId,
			&bankAccount.BankName,
			&bankAccount.BankAccountName,
			&bankAccount.BankAccountNumber,
		)
		if err != nil {
			return nil, err
		}

		bankAccounts = append(bankAccounts, bankAccount)
	}

	return bankAccounts, nil
}

func (repo *BankAccountRepositoryImpl) Update(ctx context.Context, bank_account bank_account_model.BankAccount) bank_account_model.BankAccount {
	conn, err := repo.DBPool.Acquire(ctx)
	utils.PanicErr(err)
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	utils.PanicErr(err)
	defer utils.CommitOrRollback(ctx, tx)

	var bankAccountResult bank_account_model.BankAccount
	SQL_UPDATE := "update bank_accounts set bank_name = $2, bank_account_name = $3, bank_account_number = $4 WHERE bank_account_id = $1;"
	err = tx.QueryRow(ctx, SQL_UPDATE, bank_account.BankAccountId, bank_account.BankName, bank_account.BankAccountName, bank_account.BankAccountNumber).Scan(
		&bankAccountResult.BankAccountId,
		&bankAccountResult.BankName,
		&bankAccountResult.BankAccountName,
		&bankAccountResult.BankAccountNumber,
		&bankAccountResult.UserId,
	)
	utils.PanicErr(err)

	return bankAccountResult
}

func (repo *BankAccountRepositoryImpl) Delete(ctx context.Context, bank_account_id int) {
	conn, err := repo.DBPool.Acquire(ctx)
	utils.PanicErr(err)
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	utils.PanicErr(err)
	defer utils.CommitOrRollback(ctx, tx)

	SQL_DELETE := "delete from bank_accounts where bank_account_id = $1"
	_, err = tx.Exec(ctx, SQL_DELETE, bank_account_id)
	utils.PanicErr(err)

	return
}
