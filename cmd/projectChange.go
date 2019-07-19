package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var projectChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "Project change command.",
	Long:  "Project change command.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("project change command.")
		return nil
	},
}

func init() {
	projectCmd.AddCommand(projectChangeCmd)
}
