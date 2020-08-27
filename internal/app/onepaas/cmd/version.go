package cmd

import (
	"fmt"

	"github.com/onepaas/onepaas/pkg/version"
	"github.com/spf13/cobra"
)

var (
	// AppVersion represents OnePaaS version
	AppVersion string
	// GitCommit represents OnePaaS commit hash
	GitCommitHash string
	// BuildTime represents OnePaaS build time
	BuildTime string
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
	v := version.NewVersion(AppVersion, GitCommitHash, BuildTime)

	output, err := v.Render()
	if err != nil {
		return err
	}

	fmt.Println(output)
	return nil
}
