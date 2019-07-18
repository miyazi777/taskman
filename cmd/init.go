package cmd

import (
	"fmt"

	"github.com/miyazi777/taskman/db"
	"github.com/spf13/cobra"
)

// initCmd represents the info command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {

		db.InitDb()
		fmt.Println("init command.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
