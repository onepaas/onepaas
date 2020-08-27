package migration

import (
	"github.com/go-pg/migrations/v8"
	"github.com/onepaas/onepaas/internal/pkg/db"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// NewUpCommand creates the up sub-command
func NewUpCommand() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "up",
		Short: "Runs available migrations.",
		RunE:   runUp,
	}

	return initCmd
}

func runUp(_ *cobra.Command, args []string) (err error) {
	oldVersion, newVersion, err := migrations.Run(db.GetDB(), "up")
	if err != nil {
		return
	}

	if newVersion != oldVersion {
		log.Info().Msgf("Migrated from version %d to %d", oldVersion, newVersion)
	} else {
		log.Info().Msg("The all migrations have already migrated")
	}

	return
}
