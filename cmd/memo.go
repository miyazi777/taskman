package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var memoCmd = &cobra.Command{
	Use:   "memo",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("memo command.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(memoCmd)
}
