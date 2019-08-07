package project

import (
	"fmt"

	"github.com/miyazi777/taskman/db"
	"github.com/miyazi777/taskman/shell"
	"github.com/spf13/cobra"
)

// initCmd represents the info command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		var title string
		var err error
		if len(args) == 0 {
			title, err = shell.EditTextByEditor("")
			if err != nil {
				return err
			}
		} else {
			title = args[0]
		}

		var currentFlg bool = false
		if !projectRepository.CheckExistsCurrentProject() {
			currentFlg = true
		}

		project := db.Project{Title: title, CurrentFlg: currentFlg}
		projectRepository.Insert(&project)

		fmt.Printf("Add project : %s\n", title)
		return nil
	},
}

func init() {
	projectCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
