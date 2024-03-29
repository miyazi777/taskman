package project

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tatsushid/go-prettytable"
)

// initCmd represents the info command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("project list")
		projects := projectRepository.GetList()

		tbl, err := prettytable.NewTable([]prettytable.Column{
			{Header: "ID", MinWidth: 3},
			{Header: "PROJECT", MinWidth: 32},
			{Header: "CURRENT", MinWidth: 1},
		}...)
		if err != nil {
			return err
		}

		tbl.Separator = " "

		for _, project := range *projects {
			var current string = ""
			if project.CurrentFlg {
				current = "*"
			}
			tbl.AddRow(project.ID, project.Title, current)
		}
		tbl.Print()

		return nil
	},
}

func init() {
	projectCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
