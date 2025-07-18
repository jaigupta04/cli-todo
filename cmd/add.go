package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/jaigupta04/cli-todo/internal"
)

var (
	addTitle    string
	addCategory string
	addDue      string
	addPriority string
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&addTitle, "title", "t", "", "Title of the task")
	addCmd.Flags().StringVarP(&addCategory, "category", "c", "", "Category for the task")
	addCmd.Flags().StringVarP(&addDue, "due", "d", "", "Due date (YYYY-MM-DD)")
	addCmd.Flags().StringVarP(&addPriority, "priority", "p", "medium", "Task priority (low, medium, high)")
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if addTitle == "" {
			fmt.Println("Error: --title is required")
			return
		}

		priority := internal.Medium // default
		switch addPriority {
		case "low":
			priority = internal.Low
		case "high":
			priority = internal.High
		}

		var dueDate string
		if addDue != "" {
			if _, err := time.Parse("2006-01-02", addDue); err != nil {
				fmt.Println("Error: invalid due date format, use YYYY-MM-DD")
				return
			}
			dueDate = addDue
		}

		tasks, err := internal.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		newID := 1
		if len(tasks) > 0 {
			newID = tasks[len(tasks)-1].ID + 1
		}

		task := internal.Task{
			ID:        newID,
			Title:     addTitle,
			Completed: false,
			Priority:  priority,
			Category:  addCategory,
			DueDate:   dueDate,
		}

		tasks = append(tasks, task)
		if err := internal.SaveTasks(tasks); err != nil {
			fmt.Println("Error saving task:", err)
			return
		}
		fmt.Printf("Task added [ID: %v]: %s\n", task.ID, task.Title)
	},
}
