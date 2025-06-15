package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/khareyash05/template-scaffold/pkg/prompt"
	"github.com/khareyash05/template-scaffold/pkg/registry"
	"github.com/khareyash05/template-scaffold/pkg/renderer"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	tplSource  string
	outputDir  string
	dryRun     bool
	diffMode   bool
	configFile string
)

type ScaffoldConfig struct {
	Variables map[string]interface{} `json:"variables" yaml:"variables"`
}

func init() {
	initCmd := &cobra.Command{
		Use:   "init [template]",
		Short: "Generate a new project from a template",
		Long: `Generate a new project from a template.
You can provide variables either interactively or via a config file.

Example config.yaml:
variables:
  author: John Doe
  version: 1.0.0
  description: A sample project
  license: MIT`,
		Args: cobra.ExactArgs(1),
		RunE: runInit,
	}

	initCmd.Flags().StringVarP(&tplSource, "source", "s", "./templates", "template source (path or repo@version)")
	initCmd.Flags().StringVarP(&outputDir, "output", "o", ".", "where to generate files")
	initCmd.Flags().BoolVar(&dryRun, "dry-run", false, "do not write files")
	initCmd.Flags().BoolVar(&diffMode, "diff", false, "show diff in dry-run")
	initCmd.Flags().StringVarP(&configFile, "config", "c", "", "path to YAML/JSON file with variable values (non-interactive)")

	rootCmd.AddCommand(initCmd)
}

func readConfig(filePath string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	config := &ScaffoldConfig{}

	// Try YAML first
	err = yaml.Unmarshal(data, config)
	if err == nil {
		return config.Variables, nil
	}

	// If YAML fails, try JSON
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("file is neither valid YAML nor JSON: %v", err)
	}

	return config.Variables, nil
}

func runInit(cmd *cobra.Command, args []string) error {
	name := args[0]
	// 1. Fetch template (local or Git) -> unpack into temp dir
	tplPath, err := registry.Fetch(tplSource, name)
	if err != nil {
		return err
	}

	var vars map[string]interface{}

	// 2. Collect variables from user (interactive or via config)
	if configFile != "" {
		vars, err = readConfig(configFile)
		if err != nil {
			return fmt.Errorf("error reading config file: %v", err)
		}
	} else {
		vars, err = prompt.Collect(tplPath)
		if err != nil {
			return err
		}
	}

	// 3. Render into outputDir (or show dry-run)
	if dryRun {
		return renderer.DryRun(tplPath, vars, outputDir, diffMode)
	}
	return renderer.Render(tplPath, vars, outputDir)
}
