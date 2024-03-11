package app

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	DB_NAME     = "openidea_shopifyx"
	DB_USERNAME = "root"
	DB_PASSWORD = "admpassword"
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
)

func GetDBConnection(context context.Context) *pgxpool.Conn {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	config, err := pgxpool.ParseConfig(dbUrl)

	config.MaxConnLifetime = 60 * time.Minute
	config.MaxConnIdleTime = 5 * time.Minute
	config.MaxConns = 100
	config.MinConns = 10

	if err != nil {
		panic(err)
	}

	dbPool, err := pgxpool.NewWithConfig(context, config)

	if err != nil {
		panic(err)
	}

	defer dbPool.Close()

	conn, err := dbPool.Acquire(context)
	if err != nil {
		panic(err)
	}

	defer conn.Release()

	return conn
}
