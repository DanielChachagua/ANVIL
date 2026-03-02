package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var archConfigs = map[string]string{
	"hexagonal": `{
  "paths": {
    "models": "-",
    "schemas": "",
    "ports": "internal/core/ports",
    "services": "internal/core/services",
    "repositories": "internal/adapters/repositories",
    "controllers": "internal/adapters/handlers",
    "routes": "internal/adapters/routes"
  }
}`,
	"clean": `{
  "paths": {
    "models": "domain/models",
    "schemas": "infrastructure/schemas",
    "ports": "domain/ports",
    "services": "application/services",
    "repositories": "infrastructure/repositories",
    "controllers": "interfaces/http/controllers",
    "routes": "interfaces/http/routes"
  }
}`,
	"n-tier": `{
  "paths": {
    "models": "models",
    "schemas": "schemas",
    "ports": "-",
    "services": "services",
    "repositories": "repositories",
    "controllers": "controllers",
    "routes": "routes"
  }
}`,
}

var arch string

var initCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "Initialize an anvil.json file in the specified directory",
	Long:  `Generates a standard anvil.json configuration template file to define the layer architecture routing within your current codebase.`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := "."
		if len(args) > 0 {
			dir = args[0]
		}

		filePath := filepath.Join(dir, "anvil.json")

		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}

		if _, err := os.Stat(filePath); err == nil {
			return fmt.Errorf("anvil.json already exists at %s", filePath)
		}

		configData, ok := archConfigs[arch]
		if !ok {
			return fmt.Errorf("unsupported architecture: %s. Available options: hexagonal, clean, n-tier", arch)
		}

		err := os.WriteFile(filePath, []byte(configData), 0644)
		if err != nil {
			return fmt.Errorf("could not write anvil.json: %w", err)
		}

		fmt.Printf("Successfully created %s for %q architecture\n", filePath, arch)
		return nil
	},
}

func init() {
	initCmd.Flags().StringVarP(&arch, "arch", "a", "hexagonal", "Architecture type (hexagonal, clean, n-tier)")
	rootCmd.AddCommand(initCmd)
}
