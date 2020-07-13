package cmd

import (
	"fmt"
	"strings"

	"github.com/drofloh/gophercises/taskmanager/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong: ", err.Error())
			return
		}
		fmt.Printf("Added task \"%s\" to task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
