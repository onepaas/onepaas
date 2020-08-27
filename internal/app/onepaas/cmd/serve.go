package cmd

import (
	"github.com/onepaas/onepaas/internal/app/onepaas"
	"github.com/onepaas/onepaas/internal/pkg/db"
	"github.com/onepaas/onepaas/pkg/config"
	"github.com/spf13/cobra"
)

// NewServeCommand creates the serve sub-command
func NewServeCommand() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "serve",
		Short: "Run OnePaaS",
		Long:  "Run OnePaaS API Server",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			db.InitDB(config.GetConfig())

			return nil
		},
		RunE:   runServe,
	}

	return serverCmd
}

func runServe(_ *cobra.Command, _ []string) error {
	as := onepaas.NewApiServer(config.GetString("address"), config.GetBool("debug"))

	return as.Run()
}
