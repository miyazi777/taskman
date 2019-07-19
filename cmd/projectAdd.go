package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var projectAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Project add command.",
	Long:  "Project add command.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("project add command.")
		return nil
	},
}

func init() {
	projectCmd.AddCommand(projectAddCmd)
}
