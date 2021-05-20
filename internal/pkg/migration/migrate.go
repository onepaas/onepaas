package migration

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/httpfs"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/onepaas/onepaas/internal/app/onepaas/migrations"
	"gorm.io/gorm"
)

func NewMigrate(database *gorm.DB) (*migrate.Migrate, error) {
	db, err := database.DB()
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"embed://",
		"pgx",
		driver,
	)
	if err != nil {
		return nil, err
	}

	return m, nil
}
