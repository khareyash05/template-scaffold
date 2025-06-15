package cmd

import (
	"github.com/khareyash05/scaffold-templates/pkg/registry"
	"github.com/khareyash05/scaffold-templates/pkg/renderer"
	"github.com/khareyash05/scaffold-templates/pkg/prompt"
	"github.com/spf13/cobra"
)

var (
	tplSource string
	outputDir string
	dryRun    bool
	diffMode  bool
)

func init() {
	initCmd := &cobra.Command{
		Use:   "init [template]",
		Short: "Generate a new project from a template",
		Args:  cobra.ExactArgs(1),
		RunE:  runInit,
	}

	initCmd.Flags().StringVarP(&tplSource, "source", "s", "./templates", "template source (path or repo@version)")
	initCmd.Flags().StringVarP(&outputDir, "output", "o", ".", "where to generate files")
	initCmd.Flags().BoolVar(&dryRun, "dry-run", false, "do not write files")
	initCmd.Flags().BoolVar(&diffMode, "diff", false, "show diff in dry-run")

	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) error {
	name := args[0]
	// 1. Fetch template (local or Git) -> unpack into temp dir
	tplPath, err := registry.Fetch(tplSource, name)
	if err != nil {
		return err
	}

	// 2. Collect variables from user (interactive or via config)
	vars, err := prompt.Collect(tplPath)
	if err != nil {
		return err
	}

	// 3. Render into outputDir (or show dry-run)
	if dryRun {
		return renderer.DryRun(tplPath, vars, outputDir, diffMode)
	}
	return renderer.Render(tplPath, vars, outputDir)
}
