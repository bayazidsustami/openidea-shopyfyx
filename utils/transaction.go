package utils

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CommitOrRollback(ctx context.Context, tx pgx.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback(ctx)
		PanicErr(errorRollback)
	} else {
		errorCommit := tx.Commit(ctx)
		PanicErr(errorCommit)
	}
}
