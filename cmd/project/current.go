package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the info command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {

		project := projectRepository.GetCurrentProject()

		fmt.Printf("Current project current ID:%d Title:%s\n", project.ID, project.Title)

		return nil
	},
}

func init() {
	projectCmd.AddCommand(currentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
