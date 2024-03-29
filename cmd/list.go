package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/tatsushid/go-prettytable"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		sort, _ := cmd.Flags().GetString("sort")

		if sort != "" {
			items := [...]string{"id", "task", "status", "label", "priority", "due_date"}
			existFlg := false
			for _, item := range items {
				if item == sort {
					existFlg = true
				}
			}
			if !existFlg {
				return errors.New("Error sort element")
			}
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		all, _ := cmd.Flags().GetBool("all")
		label, _ := cmd.Flags().GetString("label")
		sort, _ := cmd.Flags().GetString("sort")

		if sort == "task" {
			sort = "title"
		}

		tasks := taskRepository.GetList(all, label, sort)

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
	listCmd.Flags().StringP("label", "l", "", "label filter")
	listCmd.Flags().BoolP("all", "a", false, "display all list")
	listCmd.Flags().StringP("sort", "s", "", "sort list")
	rootCmd.AddCommand(listCmd)
}
