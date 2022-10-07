package config

type JWT struct {
	JwtSecret     string `mapstructure:"JWT_SECRET"`
	RefreshSecret string `mapstructure:"REFRESH_SECRET"`
}
