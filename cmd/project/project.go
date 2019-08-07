package project

import (
	"github.com/miyazi777/taskman/db"
	"github.com/spf13/cobra"
)

var projectRepository = db.ProjectRepositoryImpl{}

// initCmd represents the info command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "",
	Long:  "",
}

func GetProjectCmd() *cobra.Command {
	return projectCmd
}

//func GetProjectRepository() *db.ProjectRepository {
//	return projectRepository
//}
