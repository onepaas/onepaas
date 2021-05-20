package migration

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/onepaas/onepaas/internal/pkg/database"
	"github.com/onepaas/onepaas/internal/pkg/migration"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// NewDownCommand creates the down sub-command
func NewDownCommand() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "down",
		Short: "Applying all down migrations",
		Long:  "Down looks at the currently active migration version and will migrate all the way down.",
		RunE:  runDown,
	}

	return initCmd
}

func runDown(_ *cobra.Command, args []string) error {
	downConfirmed := false
	prompt := &survey.Confirm{
		Message: "Are you sure you want to apply all down migrations?",
	}

	survey.AskOne(prompt, &downConfirmed)

	m, err := migration.NewMigrate(database.InitDB())
	if err != nil {
		return err
	}

	if downConfirmed {
		err = m.Down()
		if err != nil {
			return err
		}

		log.Info().Msg("The down migrations have migrated.")

		return nil
	}

	log.Info().Msg("The down migrations haven't migrated.")

	return nil
}
