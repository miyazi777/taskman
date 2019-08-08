package cmd

import (
	"errors"
	"strconv"

	"github.com/miyazi777/taskman/shell"
	"github.com/spf13/cobra"
)

var titleCmd = &cobra.Command{
	Use:   "title",
	Short: "",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Requires task id.")
		}

		_, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("Numeric error.")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		id, _ := strconv.Atoi(args[0])

		task := taskRepository.FindById(id)
		if task == nil {
			return errors.New("Not found task")
		}
		oldTitle := task.Title

		var err error
		if len(args) == 1 {
			task.Title, err = shell.EditTextByEditor(task.Title)
			if err != nil {
				return err
			}
		} else {
			task.Title = args[1]
		}
		// TODO 入力されたファイル名がすでに存在しているかチェックする

		taskRepository.Update(task)
		memoRepository.ChangeMemo(task.ProjectID, oldTitle, task.Title)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(titleCmd)
}
