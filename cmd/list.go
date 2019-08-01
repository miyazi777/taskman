package cmd

import (
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {

		tasks := taskRepository.GetList()

		t := tabby.New()
		t.AddHeader("ID", "TASK", "STATUS")
		for _, task := range *tasks {
			t.AddLine(task.ID, task.Title, task.Status)
		}
		t.Print()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
