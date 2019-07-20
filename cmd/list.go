package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("list command.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
