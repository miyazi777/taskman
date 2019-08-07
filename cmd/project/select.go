package project

import (
	"errors"
	"strconv"

	"github.com/spf13/cobra"
)

// initCmd represents the info command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
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

		oldCurrentProject := projectRepository.GetCurrentProject()
		if oldCurrentProject == nil {
			return errors.New("Notfound current project.")
		}

		newCurreentProject := projectRepository.FindById(id)
		if newCurreentProject == nil {
			return errors.New("Notfound project by id.")
		}

		oldCurrentProject.CurrentFlg = false
		projectRepository.Update(oldCurrentProject)

		newCurreentProject.CurrentFlg = true
		projectRepository.Update(newCurreentProject)

		return nil
	},
}

func init() {
	projectCmd.AddCommand(selectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
