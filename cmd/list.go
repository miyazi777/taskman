package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {

		tasks := taskRepository.GetList()

		for _, task := range *tasks {
			fmt.Printf("%d %s %s\n", task.ID, task.Title, task.Status)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
