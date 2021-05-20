package migration

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/onepaas/onepaas/internal/pkg/database"
	"github.com/onepaas/onepaas/internal/pkg/migration"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// NewUpCommand creates the up sub-command
func NewUpCommand() *cobra.Command {
	upCmd := &cobra.Command{
		Use:        "up",
		Short:      "Applying all up migrations",
		Long:       "Up looks at the currently active migration version and will migrate all the way up.",
		RunE: runUp,
	}

	return upCmd
}

func runUp(_ *cobra.Command, args []string) error {
	m, err := migration.NewMigrate(database.InitDB())
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			return err
		}

		log.Info().Msg("The all migrations have already migrated.")

		return nil
	}

	log.Info().Msg("The migrations have migrated.")

	return nil
}
