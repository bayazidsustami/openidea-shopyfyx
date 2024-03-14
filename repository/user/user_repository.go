package user_repository

import (
	"context"
	user_model "openidea-shopyfyx/models/user"

	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	Register(ctx context.Context, tx pgx.Tx, request user_model.User) (user_model.User, error)
	Login(ctx context.Context, tx pgx.Tx, request user_model.User) (user_model.User, error)
}
