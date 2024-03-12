package config

import "github.com/spf13/viper"

func EnvBinder() {
	/**
	 * Env variables related to app
	 */
	viper.BindEnv("APP_HOST")
	viper.BindEnv("APP_PORT")

	/**
	 * Env variablees related to database
	 */
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DB_USERNAME")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_PORT")
}
