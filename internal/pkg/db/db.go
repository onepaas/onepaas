package db

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/onepaas/onepaas/pkg/config"
)

var db *pg.DB

func InitDB(c config.Config) *pg.DB {
	db = pg.Connect(&pg.Options{
		Database: c.GetString("database.name"),
		User:     c.GetString("database.username"),
		Password: c.GetString("database.password"),
		Addr: fmt.Sprintf("%s:%d", config.GetString("database.host"), config.GetInt("database.port")),
	})

	return db
}

func GetDB() *pg.DB {
	return db
}
