package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "taskmanager",
	Short: "Taskmanager is a CLI for managing tasks",
}
