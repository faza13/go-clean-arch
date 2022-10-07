package config

import "github.com/spf13/viper"

type App struct {
	Name     string `mapstructure:"APP_NAME"`
	Port     string `mapstructure:"APP_PORT"`
	TimeZone string `mapstructure:"APP_TIMEZONE"`
	Database *Database
	JWT      *JWT
	Cache    *Cache
}

func NewApp() *App {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var config App

	// set cache App
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	// set cache database
	err = viper.Unmarshal(&config.Database)
	if err != nil {
		panic(err)
	}

	// set cache config
	err = viper.Unmarshal(&config.Cache)
	if err != nil {
		panic(err)
	}

	// set jwt config
	err = viper.Unmarshal(&config.JWT)
	if err != nil {
		panic(err)
	}

	return &config
}
