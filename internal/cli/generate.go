package cli

import (
	"github.com/DanielChachagua/anvil/internal/generator"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate [entity]",
	Short: "Generate scaffolding for a specific domain entity",
	Long:  `Create all the necessary base files such as models, services, repositories and HTTP controllers for a new entity in your project based on your anvil.json routing config.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		entity := args[0]
		return generator.GenerateModule(entity)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
