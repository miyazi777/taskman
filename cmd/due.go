package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

var dueCmd = &cobra.Command{
	Use:   "due",
	Short: "",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Requires task id.")
		}

		if len(args) == 1 {
			return errors.New("Requires due.")
		}

		_, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("Numeric error.")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := strconv.Atoi(args[0])

		dueDate, err := time.Parse("2006/01/02", args[1])
		if err != nil {
			return errors.New("Due date format error. Format is yyyy/MM/dd")
		}

		task := taskRepository.FindById(id)
		if task == nil {
			return errors.New("Not found task")
		}

		task.DueDate = dueDate
		taskRepository.Update(task)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(dueCmd)
}
