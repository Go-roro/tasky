package cmd

import (
	"errors"
	"flag"
	"fmt"

	"github.com/yourusername/tasky/task"
)

type DoneCommand struct {
	tasks *task.Tasks
}

func (c *DoneCommand) Name() string {
	return "done"
}

func (c *DoneCommand) CommandDescription() string {
	return "Mark a task as done"
}
func (c *DoneCommand) CommandHelp() string {
	return `done - Mark a task as done

Usage:
  done --id <task_id>

Options:
  --id       ID of the task to mark as done (required)

Examples:
  done --id 1
`
}

func (c *DoneCommand) OnAction(args []string) (string, error) {
	id, err := parseDoneArgs(args)
	if err != nil {
		return "", err
	}

	taskByID, err := c.tasks.FindTaskByID(id)
	if err != nil {
		return "", err
	}
	taskByID.MarkAsDone()
	return generateResult(taskByID), nil
}

func generateResult(task *task.Task) string {
	return fmt.Sprintf("✅ Task %d marked as done: %s", task.ID, task.Title)
}

func parseDoneArgs(args []string) (int, error) {
	flagSet := flag.NewFlagSet("done", flag.ContinueOnError)
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
