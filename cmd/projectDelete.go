package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var projectDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Project delete command.",
	Long:  "Project delete command.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("project delete command.")
		return nil
	},
}

func init() {
	projectCmd.AddCommand(projectDeleteCmd)
}
