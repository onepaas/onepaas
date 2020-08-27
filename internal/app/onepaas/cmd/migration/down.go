package migration

import (
	"github.com/go-pg/migrations/v8"
	"github.com/onepaas/onepaas/internal/pkg/db"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// NewDownCommand creates the down sub-command
func NewDownCommand() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "down",
		Short: "Reverts last migration",
		RunE:   runDown,
	}

	return initCmd
}

func runDown(_ *cobra.Command, args []string) (err error) {
	oldVersion, newVersion, err := migrations.Run(db.GetDB(), "down")
	if err != nil {
		return
	}

	if newVersion != oldVersion {
		log.Info().Msgf("Migrated from version %d to %d", oldVersion, newVersion)
	} else {
		log.Info().Msg("The last version has already reverted")
	}

	return
}
