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
		title, _ := cmd.PersistentFlags().GetString("title")

		fmt.Println("project add command. : " + title)
		return nil
	},
}

func init() {
	projectAddCmd.PersistentFlags().StringP("title", "t", "", "project title")
	projectCmd.AddCommand(projectAddCmd)
}
