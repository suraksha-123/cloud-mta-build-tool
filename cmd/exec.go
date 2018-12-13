package commands

import (
	"os"

	"github.com/spf13/cobra"

	"cloud-mta-build-tool/internal/artifacts"
)

var sourceMtadFlag string
var targetMtadFlag string
var platformMtadFlag string
var sourceMetaFlag string
var targetMetaFlag string
var platformMetaFlag string
var sourceMtarFlag string
var targetMtarFlag string
var sourcePackFlag string
var targetPackFlag string
var sourceBModuleFlag string
var targetBModuleFlag string
var sourceCleanupFlag string
var targetCleanupFlag string
var sourceValidateFlag string

var pPackModuleFlag string
var pBuildModuleNameFlag string
var pValidationFlag string

var descriptorMtadFlag string
var descriptorMtarFlag string
var descriptorMetaFlag string
var descriptorPackFlag string
var descriptorBModuleFlag string
var descriptorCleanupFlag string
var descriptorValidateFlag string

func init() {

	// set source and target path flags of commands
	genMtadCmd.Flags().StringVarP(&sourceMtadFlag, "source", "s", "", "Provide MTA source ")
	genMtadCmd.Flags().StringVarP(&targetMtadFlag, "target", "t", "", "Provide MTA target ")
	genMtadCmd.Flags().StringVarP(&platformMtadFlag, "platform", "p", "", "Provide MTA platform ")
	genMetaCmd.Flags().StringVarP(&sourceMetaFlag, "source", "s", "", "Provide MTA source ")
	genMetaCmd.Flags().StringVarP(&targetMetaFlag, "target", "t", "", "Provide MTA target ")
	genMetaCmd.Flags().StringVarP(&platformMetaFlag, "platform", "p", "", "Provide MTA platform ")
	genMtarCmd.Flags().StringVarP(&sourceMtarFlag, "source", "s", "", "Provide MTA source ")
	genMtarCmd.Flags().StringVarP(&targetMtarFlag, "target", "t", "", "Provide MTA target ")
	packCmd.Flags().StringVarP(&sourcePackFlag, "source", "s", "", "Provide MTA source ")
	packCmd.Flags().StringVarP(&targetPackFlag, "target", "t", "", "Provide MTA target ")
	bModuleCmd.Flags().StringVarP(&sourceBModuleFlag, "source", "s", "", "Provide MTA source ")
	bModuleCmd.Flags().StringVarP(&targetBModuleFlag, "target", "t", "", "Provide MTA target ")
	cleanupCmd.Flags().StringVarP(&sourceCleanupFlag, "source", "s", "", "Provide MTA source ")
	cleanupCmd.Flags().StringVarP(&targetCleanupFlag, "target", "t", "", "Provide MTA target ")
	validateCmd.Flags().StringVarP(&sourceValidateFlag, "source", "s", "", "Provide MTA source  ")

	// set module flags of module related commands
	packCmd.Flags().StringVarP(&pPackModuleFlag, "module", "m", "", "Provide Module name ")
	bModuleCmd.Flags().StringVarP(&pBuildModuleNameFlag, "module", "m", "", "Provide Module name ")

	// set flags of validation command
	validateCmd.Flags().StringVarP(&pValidationFlag, "mode", "m", "", "Provide Validation mode ")

	// set mta descriptor flag of commands
	genMtadCmd.Flags().StringVarP(&descriptorMtadFlag, "desc", "d", "", "Descriptor MTA - dev/dep")
	genMetaCmd.Flags().StringVarP(&descriptorMetaFlag, "desc", "d", "", "Descriptor MTA - dev/dep")
	genMtarCmd.Flags().StringVarP(&descriptorMtarFlag, "desc", "d", "", "Descriptor MTA - dev/dep")
	packCmd.Flags().StringVarP(&descriptorPackFlag, "desc", "d", "", "Descriptor MTA - dev/dep")
	bModuleCmd.Flags().StringVarP(&descriptorBModuleFlag, "desc", "d", "", "Descriptor MTA - dev/dep")
	cleanupCmd.Flags().StringVarP(&descriptorCleanupFlag, "desc", "d", "", "Descriptor MTA - dev/dep")
	validateCmd.Flags().StringVarP(&descriptorValidateFlag, "desc", "d", "", "Descriptor MTA - dev/dep")
}

// Build module
var bModuleCmd = &cobra.Command{
	Use:   "module",
	Short: "Build module",
	Long:  "Build specific module according to the module name",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := artifacts.ExecuteBuild(sourceBModuleFlag, targetBModuleFlag, descriptorBModuleFlag, pBuildModuleNameFlag, os.Getwd)
		logError(err)
		return err
	},
	SilenceUsage:  true,
	SilenceErrors: false,
}

// zip specific module and put the artifacts on the temp folder according
// to the mtar structure, i.e each module has new entry as folder in the mtar folder
// Note - even if the path of the module was changed in the mta.yaml in the mtar the
// the module folder will get the module name
var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "pack module artifacts",
	Long:  "pack the module artifacts after the build process",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := artifacts.ExecutePack(sourcePackFlag, targetPackFlag, descriptorPackFlag, pPackModuleFlag, os.Getwd)
		logError(err)
		return err
	},
}

// Generate metadata info from deployment
var genMetaCmd = &cobra.Command{
	Use:   "meta",
	Short: "generate meta folder",
	Long:  "generate META-INF folder with all the required data",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := artifacts.ExecuteGenMeta(sourceMetaFlag, targetMetaFlag, descriptorMetaFlag, platformMetaFlag, os.Getwd)
		logError(err)
		return err
	},
	SilenceUsage:  true,
	SilenceErrors: false,
}

// Generate mtar from build artifacts
var genMtarCmd = &cobra.Command{
	Use:   "mtar",
	Short: "generate MTAR",
	Long:  "generate MTAR from the project build artifacts",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := artifacts.ExecuteGenMtar(sourceMtarFlag, targetMtarFlag, descriptorMtarFlag, os.Getwd)
		logError(err)
		return err
	},
	SilenceUsage:  true,
	SilenceErrors: false,
}

// Provide mtad.yaml from mta.yaml
var genMtadCmd = &cobra.Command{
	Use:   "mtad",
	Short: "Provide mtad",
	Long:  "Provide deployment descriptor (mtad.yaml) from development descriptor (mta.yaml)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := artifacts.ExecuteGenMtad(sourceMtadFlag, targetMtadFlag, descriptorMtadFlag, platformMtadFlag, os.Getwd)
		logError(err)
		return err
	},
	SilenceUsage:  true,
	SilenceErrors: false,
}

// Validate mta.yaml
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "MBT validation",
	Long:  "MBT validation process",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := artifacts.ExecuteValidation(sourceValidateFlag, descriptorValidateFlag, pValidationFlag, os.Getwd)
		logError(err)
		return err
	},
	SilenceErrors: false,
	SilenceUsage:  true,
}

// Cleanup temp artifacts
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Remove process artifacts",
	Long:  "Remove MTA build process artifacts",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Remove temp folder
		err := artifacts.ExecuteCleanup(sourceCleanupFlag, targetCleanupFlag, descriptorCleanupFlag, os.Getwd)
		logError(err)
		return err
	},
	SilenceUsage:  true,
	SilenceErrors: false,
}