package utils

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func CommitOrRollback(ctx context.Context, tx pgx.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback(ctx)
		if errorRollback != nil {
			panic(fiber.NewError(fiber.StatusInternalServerError, "something error"))
		}
	} else {
		errorCommit := tx.Commit(ctx)
		if errorCommit != nil {
			panic(fiber.NewError(fiber.StatusInternalServerError, "something error"))
		}
	}
}
