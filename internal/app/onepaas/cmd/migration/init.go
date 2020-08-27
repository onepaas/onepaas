package migration

import (
	"github.com/go-pg/migrations/v8"
	"github.com/onepaas/onepaas/internal/pkg/db"
	"github.com/spf13/cobra"
)

// NewInitCommand creates the init sub-command
func NewInitCommand() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Creates migration info table in the database.",
		RunE:   runInit,
	}

	return initCmd
}

func runInit(_ *cobra.Command, args []string) error {
	_, _, err := migrations.Run(db.GetDB(), "init")

	return err
}
