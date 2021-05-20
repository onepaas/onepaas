package database

import (
	"fmt"

	"github.com/onepaas/onepaas/pkg/viper"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	var dialector gorm.Dialector

	switch viper.GetString("database.driver") {
	case "postgres":
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
			viper.GetString("database.host"),
			viper.GetString("database.username"),
			viper.GetString("database.password"),
			viper.GetString("database.name"),
			viper.GetInt("database.port"),
		)

		dialector = postgres.Open(dsn)

	default:
		log.Fatal().Msgf("Unknown database's driver (%s)", viper.GetString("database.driver"))
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("OnePaaS can't initialize a db session.")
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}
