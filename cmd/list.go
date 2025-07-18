package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/jaigupta04/cli-todo/internal"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := internal.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		fmt.Println("Your To-Do List:")
		for _, task := range tasks {
			status := " "
			if task.Completed {
				status = "x"
			}
			var priority string
			switch task.Priority {
			case internal.Low:
				priority = "Low"
			case internal.Medium:
				priority = "Medium"
			case internal.High:
				priority = "High"
			}
			fmt.Printf("[%s] #%d: %s | Priority: %s | Category: %s | Due: %s\n", status, task.ID, task.Title, priority, task.Category, task.DueDate)
		}
	},
}
