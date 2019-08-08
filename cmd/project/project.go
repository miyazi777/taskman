package project

import (
	"github.com/miyazi777/taskman/db"
	"github.com/miyazi777/taskman/repository"
	"github.com/spf13/cobra"
)

var projectRepository = db.ProjectRepositoryImpl{}
var memoRepository = repository.MemoRepositoryImpl{}

// initCmd represents the info command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "",
	Long:  "",
}

func GetProjectCmd() *cobra.Command {
	return projectCmd
}
