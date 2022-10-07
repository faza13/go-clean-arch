package config

type Cache struct {
	Host     string `mapstructure:"REDIS_HOST"`
	Database string `mapstructure:"REDIS_DATABASE"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	Port     string `mapstructure:"REDIS_PORT"`
}
