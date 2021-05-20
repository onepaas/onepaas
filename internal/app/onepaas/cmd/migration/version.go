package migration

import (
	"github.com/onepaas/onepaas/internal/pkg/database"
	"github.com/onepaas/onepaas/internal/pkg/migration"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// NewVersionCommand creates the version sub-command
func NewVersionCommand() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Prints current migration version",
		Long:  "Version returns the currently active migration version.",
		RunE:  runVersion,
	}

	return versionCmd
}

func runVersion(_ *cobra.Command, args []string) (err error) {
	m, err := migration.NewMigrate(database.InitDB())
	if err != nil {
		return err
	}

	v, dirty, err := m.Version()
	if err != nil {
		return err
	}

	if dirty {
		log.Info().Msgf("%v (dirty)\n", v)
	} else {
		log.Info().Msgf("Current version is %d", v)
	}

	return nil
}
