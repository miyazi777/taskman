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

		tasks := taskRepository.GetList()

		tbl, err := prettytable.NewTable([]prettytable.Column{
			{Header: "ID", MinWidth: 3},
			{Header: "TASK", MinWidth: 32},
			{Header: "STATUS", MinWidth: 12},
			{Header: "LABEL", MinWidth: 64},
			{Header: "DUE_DATE", MinWidth: 10},
		}...)
		if err != nil {
			return err
		}
		tbl.Separator = " "

		onlyListFlg, _ := cmd.PersistentFlags().GetBool("only-list")
		if onlyListFlg {
			tbl.NoHeader = true
		}

		for _, task := range *tasks {
			d := task.DueDate
			tbl.AddRow(task.ID, task.Title, task.Status, task.Label, d.Format("2006/01/02"))
		}
		tbl.Print()

		return nil
	},
}

func init() {
	listCmd.PersistentFlags().BoolP("only-list", "", false, "display list only")
	rootCmd.AddCommand(listCmd)
}
