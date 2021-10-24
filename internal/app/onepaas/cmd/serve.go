package cmd

import (
	"github.com/onepaas/onepaas/internal/app/onepaas"
	"github.com/onepaas/onepaas/internal/pkg/database"
	"github.com/onepaas/onepaas/pkg/viper"
	"github.com/spf13/cobra"
)

// NewServeCommand creates the serve sub-command
func NewServeCommand() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "serve",
		Short: "Run OnePaaS",
		Long:  "Run OnePaaS API Server",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			database.InitRedis()

			return nil
		},
		RunE:   runServe,
	}

	return serverCmd
}

func runServe(_ *cobra.Command, _ []string) error {
	as := onepaas.NewApiServer(viper.GetString("address"), viper.GetBool("debug"))

	return as.Run()
}
