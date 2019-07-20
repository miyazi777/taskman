package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("add command.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
