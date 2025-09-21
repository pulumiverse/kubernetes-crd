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
		//	config, err := pkg.LoadLocalConfig(generateArgs.ConfigPath)
		//	if err != nil {
		//		return err
		//	}
		//
		//	// Template name priority: CLI flag > config file
		//	if generateArgs.TemplateName == "" {
		//		if config.Template != "" {
		//			generateArgs.TemplateName = config.Template
		//		}
		//	}
		//
		//	// Name priority: CLI flag > config file ("repository", then "name" field)
		//	if generateArgs.RepositoryName == "" {
		//		if config.Repository != "" {
		//			generateArgs.RepositoryName = config.Repository
		//		} else {
		//			providerName := config.Provider
		//			organizationName := config.Organization
		//			if providerName != "" && organizationName != "" {
		//				generateArgs.RepositoryName = fmt.Sprintf("%s/pulumi-%s", organizationName, providerName)
		//			}
		//		}
		//	}
		//
		//	if generateArgs.RepositoryName == "" {
		//		return fmt.Errorf("repository name must be set either in the config file or via the --name flag")
		//	}
		//
		//	err = pkg.GeneratePackage(pkg.GenerateOpts{
		//		RepositoryName: generateArgs.RepositoryName,
		//		OutDir:         generateArgs.OutDir,
		//		TemplateName:   generateArgs.TemplateName,
		//		Config:         config,
		//		SkipMigrations: generateArgs.SkipMigrations,
		//	})
		//	return err
		return nil
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
