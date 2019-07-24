package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var projectCurrentCmd = &cobra.Command{
	Use:   "current",
	Short: "Project current command.",
	Long:  "Project current command.",
	RunE: func(cmd *cobra.Command, args []string) error {
		project, err := projectRepository.FindByCurrent()
		if err != nil {
			return err
		}

		fmt.Println(project.Title)
		return nil
	},
}

func init() {
	projectCmd.AddCommand(projectCurrentCmd)
}
