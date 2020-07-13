package cmd

import (
	"fmt"
	"strconv"

	"github.com/drofloh/gophercises/taskmanager/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as done.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to print the argument:", arg)

			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue

			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark task \"%d\" as done. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked task \"%d\" as done.\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
