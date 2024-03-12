package user_repository

import (
	"context"
	user_model "openidea-shopyfyx/models/user"
	"openidea-shopyfyx/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryImpl struct {
	DBPool *pgxpool.Pool
}

func New(DBPool *pgxpool.Pool) UserRepository {
	return &UserRepositoryImpl{
		DBPool: DBPool,
	}
}

func (repo *UserRepositoryImpl) Register(ctx context.Context, user user_model.User) user_model.User {
	conn, err := repo.DBPool.Acquire(ctx)
	utils.PanicErr(err)
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	utils.PanicErr(err)
	defer utils.CommitOrRollback(ctx, tx)

	SQL := "insert into users(username, password, name) values ($1, $2, $3)"

	_, err = tx.Exec(ctx, SQL, user.Username, user.Password, user.Name)
	utils.PanicErr(err)

	return user
}

func (repo *UserRepositoryImpl) Login(ctx context.Context, user user_model.User) user_model.User {
	//todo
	return user_model.User{}
}
