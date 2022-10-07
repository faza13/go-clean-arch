package config

type Database struct {
	Connection string `mapstructure:"DB_CONNECTION"`
	Host       string `mapstructure:"DB_HOST"`
	Database   string `mapstructure:"DB_DATABASE"`
	Username   string `mapstructure:"DB_USERNAME"`
	Password   string `mapstructure:"DB_PASSWORD"`
	Port       string `mapstructure:"DB_PORT"`
	Timezone   string `mapstructure:"DB_TIMEZONE"`
}
