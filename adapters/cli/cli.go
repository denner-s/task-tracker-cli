package cli

import (
	"fmt"
	"github.com/denner-s/task-tracker-cli/core"
	"os"
	"strconv"
)

type CLI struct {
	service *core.TaskService
}

func NewCLI(service *core.TaskService) *CLI {
	return &CLI{service: service}
}

func (c *CLI) HandleCommand(args []string) {
	if len(args) == 0 {
		fmt.Println("Available commands: add, update, delete, mark-in-progress, mark-done, list")
		os.Exit(1)
	}

	switch args[0] {
	case "add":
		c.handleAdd(args[1:])
	case "update":
		c.handleUpdate(args[1:])
	case "delete":
		c.handleDelete(args[1:])
	case "mark-in-progress":
		c.handleMarkInProgress(args[1:])
	case "mark-done":
		c.handleMarkDone(args[1:])
	case "list":
		c.handleList(args[1:])
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}

func (c *CLI) handleAdd(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: add <description>")
		os.Exit(1)
	}

	task, err := c.service.AddTask(args[0])
	if err != nil {
		fmt.Printf("Error adding task: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
}

func (c *CLI) handleUpdate(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: update <id> <new_description>")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Invalid ID: %v\n", err)
		os.Exit(1)
	}

	task, err := c.service.UpdateTask(id, args[1])
	if err != nil {
		fmt.Printf("Error updating task: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Task updated successfully (ID: %d)\n", task.ID)
}

func (c *CLI) handleDelete(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: delete <id>")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Invalid ID: %v\n", err)
		os.Exit(1)
	}

	if err := c.service.DeleteTask(id); err != nil {
		fmt.Printf("Error deleting task: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Task deleted successfully (ID: %d)\n", id)
}

func (c *CLI) handleMarkInProgress(args []string) {
	c.handleMarkStatus(args, "in-progress", core.StatusInProgress)
}

func (c *CLI) handleMarkDone(args []string) {
	c.handleMarkStatus(args, "done", core.StatusDone)
}

func (c *CLI) handleMarkStatus(args []string, status string, targetStatus core.TaskStatus) {
	if len(args) < 1 {
		fmt.Printf("Usage: mark-%s <id>\n", status)
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Invalid ID: %v\n", err)
		os.Exit(1)
	}

	task, err := c.service.UpdateStatus(id, targetStatus)
	if err != nil {
		fmt.Printf("Error marking task as %s: %v\n", status, err)
		os.Exit(1)
	}
	fmt.Printf("Task marked as %s (ID: %d)\n", status, task.ID)
}

func (c *CLI) handleList(args []string) {
	var filter string
	if len(args) > 0 {
		filter = args[0]
	}

	tasks, err := c.service.ListTasks(filter)
	if err != nil {
		fmt.Printf("Error listing tasks: %v\n", err)
		os.Exit(1)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	for _, task := range tasks {
		fmt.Printf("ID: %d | %s | Status: %s | Created: %s | Updated: %s\n",
			task.ID,
			task.Description,
			task.Status,
			task.CreatedAt.Format("2006-01-02 15:04"),
			task.UpdatedAt.Format("2006-01-02 15:04"),
		)
	}
}
