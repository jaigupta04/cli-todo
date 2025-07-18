package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/jaigupta04/cli-todo/internal"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Delete a task by its ID",
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

		index := -1
		for i, task := range tasks {
			if task.ID == id {
				index = i
				break
			}
		}
		if index == -1 {
			fmt.Printf("Task #%d not found.\n", id)
			return
		}

		tasks = append(tasks[:index], tasks[index+1:]...)
		if err := internal.SaveTasks(tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
		fmt.Printf("Task #%d deleted.\n", id)
	},
}
