package cmd

import (
	"fmt"

	"errors"
	"github.com/spf13/cobra"
	"strconv"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
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

		taskRepository.DeleteTask(id)

		fmt.Println("Delete task id : %d\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
