package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "Project list command.",
	Long:  "Project list command.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("project list command.")
		return nil
	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)
}
