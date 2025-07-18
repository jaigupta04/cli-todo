package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/jaigupta04/cli-todo/internal"
)

func init() {
	rootCmd.AddCommand(completeCmd)
}

var completeCmd = &cobra.Command{
	Use:   "complete [task ID]",
	Short: "Mark a task as complete",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: Task ID must be a number")
			return
		}

		tasks, err := internal.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		found := false
		for i := range tasks {
			if tasks[i].ID == id {
				if tasks[i].Completed {
					fmt.Printf("Task #%d is already marked as complete.\n", id)
				} else {
					tasks[i].Completed = true
					fmt.Printf("Task #%d marked as complete.\n", id)
				}
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Task #%d not found.\n", id)
			return
		}

		if err := internal.SaveTasks(tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
	},
}
