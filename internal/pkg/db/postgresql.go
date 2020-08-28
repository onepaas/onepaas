package db

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/onepaas/onepaas/pkg/viper"
)

var db *pg.DB

func InitDB() *pg.DB {
	db = pg.Connect(&pg.Options{
		Database: viper.GetString("database.name"),
		User:     viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Addr: fmt.Sprintf("%s:%d", viper.GetString("database.host"), viper.GetInt("database.port")),
	})

	return db
}

func GetDB() *pg.DB {
	return db
}
