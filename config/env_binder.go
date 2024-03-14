package config

import "github.com/spf13/viper"

func init() {
	envBinder()
}

func envBinder() {
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

	/**
	* Env variables related to AWS S3
	 */
	viper.BindEnv("S3_ID")
	viper.BindEnv("S3_SECRET_KEY")
	viper.BindEnv("S3_BASE_URL")
}
