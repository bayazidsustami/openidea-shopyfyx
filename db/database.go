package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

func InitDBPool() (*pgxpool.Pool, error) {
	applicationDbName := viper.GetString("DB_NAME")
	applicationDbUsername := viper.GetString("DB_USERNAME")
	applicationDbPassword := viper.GetString("DB_PASSWORD")
	applicationDbHost := viper.GetString("DB_HOST")
	applicationDbPort := viper.GetString("DB_PORT")

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", applicationDbUsername, applicationDbPassword, applicationDbHost, applicationDbPort, applicationDbName)
	config, err := pgxpool.ParseConfig(dbUrl)

	config.MaxConnLifetime = 60 * time.Minute
	config.MaxConnIdleTime = 5 * time.Minute
	config.MaxConns = 100
	config.MinConns = 10

	if err != nil {
		return nil, err
	}

	dbPool, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		return nil, err
	}

	return dbPool, nil
}
