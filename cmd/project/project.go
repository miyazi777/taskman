package project

import (
	"github.com/spf13/cobra"
)

// initCmd represents the info command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "",
	Long:  "",
}

func GetProjectCmd() *cobra.Command {
	return projectCmd
}
