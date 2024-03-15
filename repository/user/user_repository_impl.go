package user_repository

import (
	"context"
	user_model "openidea-shopyfyx/models/user"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type UserRepositoryImpl struct {
}

func New() UserRepository {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) Register(ctx context.Context, tx pgx.Tx, user user_model.User) (user_model.User, error) {

	var userId int
	SQL_INSERT := "INSERT INTO users(username, password, name) values ($1, $2, $3) ON CONFLICT(username) DO NOTHING RETURNING user_id"
	err := tx.QueryRow(ctx, SQL_INSERT, user.Username, user.Password, user.Name).Scan(&userId)
	if err != nil {
		return user_model.User{}, err
	}

	user.UserId = userId
	return user, nil
}

func (repo *UserRepositoryImpl) Login(ctx context.Context, tx pgx.Tx, user user_model.User) (user_model.User, error) {

	var userResult user_model.User
	SQL_GET := "select user_id, username, name, password from users where username=$1"
	err := tx.QueryRow(ctx, SQL_GET, user.Username).Scan(
		&userResult.UserId,
		&userResult.Username,
		&userResult.Name,
		&userResult.Password,
	)
	if err != nil {
		return user_model.User{}, fiber.NewError(fiber.StatusInternalServerError, "something error")
	}

	return userResult, nil
}

func (repo *UserRepositoryImpl) GetSeller(ctx context.Context, tx pgx.Tx, userId int) (user_model.Seller, error) {
	SQL_GET := "SELECT COALESCE(SUM(o.quantity), 0) " +
		"FROM users u " +
		"INNER JOIN products p ON u.user_id = p.user_id " +
		"INNER JOIN orders o ON p.product_id = o.product_id " +
		"WHERE p.deleted_at IS NULL " +
		"AND u.user_id = $1"

	var seller user_model.Seller
	err := tx.QueryRow(ctx, SQL_GET, userId).Scan(&seller.ProductsSoldTotal)
	if err != nil {
		return user_model.Seller{}, fiber.NewError(fiber.StatusInternalServerError, "something error")
	}

	return seller, nil
}
