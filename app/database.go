package app

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDBConnection(context context.Context) *pgxpool.Pool {
	config, err := pgx.ParseConfig("postgres://user:password@localhost:5432/database_name")
	if err != nil {
		panic(err)
	}

	dbPool, err := pgxpool.NewWithConfig(context, &pgxpool.Config{
		ConnConfig:      config,
		MaxConnLifetime: 60 * time.Minute,
		MaxConnIdleTime: 5 * time.Minute,
		MaxConns:        100,
		MinConns:        10,
	})

	if err != nil {
		panic(err)
	}

	defer dbPool.Close()

	return dbPool
}
