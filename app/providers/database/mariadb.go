package database

import (
	"base/app/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
)

func NewMariaDBProvider(dbConfig *config.Database) *Provider {
	// set dsn mariadb
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
		url.QueryEscape(dbConfig.Timezone))

	conn, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}))
	if err != nil {
		panic("error: " + err.Error())
	}

	return &Provider{DB: conn}
}
