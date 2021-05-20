package migration

import (
	"github.com/spf13/cobra"
)

// NewMigrationCommand creates the migration sub-command
func NewMigrationCommand(parent *cobra.Command) *cobra.Command {
	migrationCmd := &cobra.Command{
		Use:   "migration",
		Short: "Print the version number of OnePaaS",
		Long:  "All software has versions. It's mine.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			err := parent.PersistentPreRunE(cmd, args)
			if err != nil {
				return err
			}

			return nil
		},
	}

	migrationCmd.AddCommand(NewUpCommand())
	migrationCmd.AddCommand(NewDownCommand())
	migrationCmd.AddCommand(NewVersionCommand())

	return migrationCmd
}
