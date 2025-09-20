package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crdk8s",
	Short: "Generate the setup for a specific CRD",
	Long:  `Generate the setup to generate & publish the various SDKs for a specific CRD.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
