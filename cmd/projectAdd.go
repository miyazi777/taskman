package cmd

import (
	"fmt"
	"github.com/miyazi777/taskman/db"
	"github.com/miyazi777/taskman/shell"
	"github.com/spf13/cobra"
)

var projectAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Project add command.",
	Long:  "Project add command.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var title string
		var err error
		title, _ = cmd.PersistentFlags().GetString("title")
		if title == "" {
			title, err = shell.EditTextByEditor("")
			if err != nil {
				return err
			}
		}

		project := db.Project{Title: title}
		projectRepository.Insert(&project)

		fmt.Println("project add command. : " + title)
		return nil
	},
}

func init() {
	projectAddCmd.PersistentFlags().StringP("title", "t", "", "project title")
	projectCmd.AddCommand(projectAddCmd)
}
