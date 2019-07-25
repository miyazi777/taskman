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
		projects := projectRepository.GetList()

		for _, project := range *projects {
			var current string = " "
			if project.Current == true {
				current = "*"
			}
			fmt.Printf("%s %s\n", current, project.Title)
		}

		return nil
	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)
}
