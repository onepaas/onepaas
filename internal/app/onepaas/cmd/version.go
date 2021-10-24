package cmd

import (
	"fmt"

	"github.com/onepaas/onepaas/pkg/version"
	"github.com/spf13/cobra"
)

// NewVersionCommand creates the version sub-command
func NewVersionCommand() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of OnePaaS",
		Long:  "All software has versions. It's mine.",
		RunE:   runVersion,
	}

	return versionCmd
}

func runVersion(_ *cobra.Command, _ []string) error {
	v := version.NewVersion(version.AppVersion, version.GitCommitHash, version.BuildTime)

	output, err := v.Render()
	if err != nil {
		return err
	}

	fmt.Println(output)
	return nil
}
