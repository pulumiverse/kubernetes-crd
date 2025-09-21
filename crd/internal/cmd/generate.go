package cmd

import (
	"fmt"

	"github.com/pulumiverse/kubernetes-crd/crd/internal/pkg"
	"github.com/spf13/cobra"
)

type generateArguments struct {
	CrdName string
}

var generateArgs generateArguments

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate SDKs for given CRD.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Read the config file `sdks.yaml`
		config, err := pkg.ReadConfig("sdks.yaml")
		if err != nil {
			return err
		}
		allCrds := *config

		// Check the CRD name passed on the CLI
		if generateArgs.CrdName == "" {
			return fmt.Errorf("CRD name must be given via the -n/--name flag")
		}

		// Check if the given CRD name is configured in the config file
		crd, found := allCrds[generateArgs.CrdName]
		if !found {
			return fmt.Errorf("CRD '%v' is not configured in `sdks.yaml` file", generateArgs.CrdName)
		}

		err = pkg.GenerateSDKs(generateArgs.CrdName, crd)
		return err
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&generateArgs.CrdName, "name", "n", "", "name of CRD to generate SDKs for. Name must be configured in `sdks.yaml` file.")
	//generateCmd.Flags().StringVarP(&generateArgs.OutDir, "out", "o", ".", "directory to write generate files to")
	//generateCmd.Flags().StringVarP(&generateArgs.TemplateName, "template", "t", "", "template name to generate (default \"{config.template}\" or otherwise \"bridged-provider\")")
	//generateCmd.Flags().StringVarP(&generateArgs.ConfigPath, "config", "c", ".ci-mgmt.yaml", "local config file to use")
	//generateCmd.Flags().BoolVar(&generateArgs.SkipMigrations, "skip-migrations", false, "skip running migrations")
}
