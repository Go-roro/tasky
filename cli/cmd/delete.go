package cmd

import (
	"errors"
	"flag"
	"fmt"

	"github.com/yourusername/tasky/task"
)

type DeleteCommand struct {
	tasks *task.Tasks
}

func (c *DeleteCommand) Name() string {
	return "delete"
}

func (c *DeleteCommand) CommandDescription() string {
	return "Delete a task by its ID"
}

func (c *DeleteCommand) CommandHelp() string {
	return `delete - Delete a task by its ID
Usage:
	
	  delete --id <task_id>
Options:
	  --id       ID of the task to delete (required)
Examples:
	  delete --id 1
`
}

func (c *DeleteCommand) OnAction(args []string) (string, error) {
	deleteId, err := parseDeleteArgs(args)
	if err != nil {
		return "", err
	}
	if err := c.tasks.DeleteTaskByID(deleteId); err != nil {
		return "", fmt.Errorf("failed to delete task with ID %d: %w", deleteId, err)
	}
	return fmt.Sprintf("🗑️ Task with ID %d has been deleted successfully", deleteId), nil
}

func parseDeleteArgs(args []string) (int, error) {
	flagSet := flag.NewFlagSet("delete", flag.ContinueOnError)
	id := flagSet.Int("id", 0, "ID of the task to mark as done")

	if err := flagSet.Parse(args); err != nil {
		return 0, fmt.Errorf("failed to parse flags")
	}

	if *id == 0 {
		return 0, errors.New("task ID is required")
	}

	if *id < 0 {
		return 0, fmt.Errorf("invalid task ID: %d", *id)
	}

	return *id, nil
}
