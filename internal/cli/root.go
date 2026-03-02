package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "anvil",
	Short: "Anvil is a powerful scaffolding CLI tool for Go",
	Long:  `Anvil helps you rapidly bootstrap application layers such as patterns for ports and adapters and basic CRUD endpoints with standard Go and frameworks like Fiber.`,
}

// Execute runs the root command configured by cobra
func Execute() error {
	return rootCmd.Execute()
}
