package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("delete command.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
