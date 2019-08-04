package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tatsushid/go-prettytable"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		all, _ := cmd.Flags().GetBool("all")

		tasks := taskRepository.GetList(all)

		tbl, err := prettytable.NewTable([]prettytable.Column{
			{Header: "ID", MinWidth: 3},
			{Header: "TASK", MinWidth: 32},
			{Header: "STATUS", MinWidth: 12},
			{Header: "LABEL", MinWidth: 16},
			{Header: "PRIORITY", MinWidth: 1},
			{Header: "DUE_DATE", MinWidth: 10},
		}...)
		if err != nil {
			return err
		}
		tbl.Separator = " "

		onlyListFlg, _ := cmd.Flags().GetBool("only-list")
		if onlyListFlg {
			tbl.NoHeader = true
		}

		for _, task := range *tasks {
			tbl.AddRow(task.ID, task.GetTitle(), task.Status, task.Label, task.Priority, task.GetDueDate())
		}
		tbl.Print()

		return nil
	},
}

func init() {
	listCmd.Flags().BoolP("only-list", "", false, "display list only")
	listCmd.Flags().BoolP("all", "a", false, "display all list")
	rootCmd.AddCommand(listCmd)
}
