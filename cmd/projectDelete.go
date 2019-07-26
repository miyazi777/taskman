package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var projectDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Project delete command.",
	Long:  "Project delete command.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Requires project id.")
		}

		_, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("Numeric error.")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := strconv.Atoi(args[0])

		project := projectRepository.FindById(id)
		fmt.Println(project)
		if project == nil {
			return errors.New("Nothing project.")
		}

		if project.Current == true {
			return errors.New("Don`t deleted current project")
		}

		projectRepository.DeleteProject(id)
		return nil
	},
}

func init() {
	projectCmd.AddCommand(projectDeleteCmd)
}
