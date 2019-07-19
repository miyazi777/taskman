package cmd

import (
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Project command.",
	Long:  "Project command.",
}

func init() {
	rootCmd.AddCommand(projectCmd)
}
