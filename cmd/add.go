package cmd

import (
	"errors"
	"fmt"

	"github.com/miyazi777/taskman/db"
	"github.com/miyazi777/taskman/shell"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		var title string
		var err error
		if len(args) == 0 {
			title, err = shell.EditTextByEditor("")
			if err != nil {
				return err
			}
		} else {
			title = args[0]
		}

		project := projectRepository.GetCurrentProject()
		if project == nil {
			return errors.New("Nothing project.")
		}

		task := db.Task{Title: title, ProjectID: project.ID}
		taskRepository.Insert(&task)

		memoRepository.AddMemo(title)

		fmt.Printf("Add task : %s\n", title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
