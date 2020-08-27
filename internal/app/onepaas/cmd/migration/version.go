package migration

import (
	"github.com/go-pg/migrations/v8"
	"github.com/onepaas/onepaas/internal/pkg/db"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// NewVersionCommand creates the version sub-command
func NewVersionCommand() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "version",
		Short: "Prints current db version",
		RunE:   runVersion,
	}

	return initCmd
}

func runVersion(_ *cobra.Command, args []string) (err error) {
	oldVersion, _, err := migrations.Run(db.GetDB(), "version")
	if err != nil {
		return
	}

	log.Info().Msgf("Version is %d", oldVersion)

	return
}
