package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"strconv"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Requires task id.")
		}

		if len(args) == 1 {
			return errors.New("Requires status.")
		}

		_, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("Numeric error.")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		id, _ := strconv.Atoi(args[0])
		status := args[1]

		task := taskRepository.FindById(id)
		if task == nil {
			return errors.New("Not found task")
		}

		task.Status = status
		taskRepository.Update(task)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
