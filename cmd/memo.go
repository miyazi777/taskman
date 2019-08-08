package cmd

import (
	"errors"

	"strconv"

	"github.com/miyazi777/taskman/repository"
	"github.com/miyazi777/taskman/shell"
	"github.com/spf13/cobra"
)

var memoCmd = &cobra.Command{
	Use:   "memo",
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
			return errors.New("Nothing task.")
		}

		memoPath := repository.GetMemoPath(task.ProjectID, task.Title)

		err := shell.StartEditor(memoPath)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(memoCmd)
}
