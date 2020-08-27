package cmd

import (
	"os"
	"time"

	"github.com/onepaas/onepaas/internal/app/onepaas/cmd/migration"
	"github.com/onepaas/onepaas/pkg/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	// LogLevel represents logging level e.g. info, warn, error, debug
	logLevel string
	cfgFile  string
)

// NewRootCommand creates the root command
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "onepaas",
		Short: "One Click to launch your application",
		Long:  `One Click to launch your application`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Prevent showing usage when subcommand return error.
			cmd.SilenceUsage = true

			config.InitConfig(cfgFile)

			lvl, err := zerolog.ParseLevel(logLevel)
			if err != nil {
				return err
			}

			log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: true})
			zerolog.SetGlobalLevel(lvl)

			return nil
		},
	}

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Path to the OnePaaS config file.")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", zerolog.InfoLevel.String(), "Set the logging level (\"trace\"|\"debug\"|\"info\"|\"warn\"|\"error\"|\"fatal\")")

	return rootCmd
}

// Execute run OnePaaS application
func Execute() {
	rootCmd := NewRootCommand()
	rootCmd.AddCommand(NewServeCommand())
	rootCmd.AddCommand(NewVersionCommand())
	rootCmd.AddCommand(migration.NewMigrationCommand(rootCmd))

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().
			Err(err)
	}
}
