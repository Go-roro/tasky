package cmd

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/yourusername/tasky/task"
)

type AddCommand struct{}

func (c *AddCommand) Name() string {
	return "add"
}

func (c *AddCommand) CommandHelp() string {
	return `add - Add a new task to your task list

Usage:
  add --title "Task title" [--priority 1|2|3] [--due "YYYY-MM-DD HH:MM"]

Options:
  --title     Title of the task (required)
  --priority  Priority level: 1 = High, 2 = Medium (default), 3 = Low
  --due       Optional due date, format: "2006-01-02 15:04"

Examples:
  add --title "Write blog post"
  add --title "Buy milk" --priority 1
  add --title "Finish homework" --due "2025-07-10 18:00"
  add --title "Submit report" --priority 1 --due "2025-07-15 09:00"
`
}

func (c *AddCommand) CommandDescription() string {
	return "Add a new task to your task list"
}

func (c *AddCommand) OnAction(args []string) (string, error) {
	title, priorityInput, due, err := parseAddOptions(args)
	if err != nil {
		return "", fmt.Errorf("failed to parse flags: %w", err)
	}

	if *title == "" {
		return "", errors.New("title is required")
	}

	priority := task.Priority(*priorityInput)
	if !priority.ValidPriority() {
		return "", fmt.Errorf("priority must be between %d and %d", task.PriorityLow, task.PriorityHigh)
	}

	dueDate, err := parseDue(*due)
	if err != nil {
		return "", err
	}

	newTask := task.Task{
		Title:     *title,
		Priority:  priority,
		Due:       dueDate,
		Done:      false,
		CreatedAt: time.Now(),
	}
	task.AddTask(&newTask)
	return generateAddTaskResult(newTask), nil
}

func parseAddOptions(args []string) (title *string, priorityInput *int, due *string, err error) {
	flagSet := flag.NewFlagSet("add", flag.ExitOnError)
	title = flagSet.String("title", "", "Title of the task")
	priorityInput = flagSet.Int("priority", 2, "Priority of the task (1-3)")
	due = flagSet.String("due", "", "Due date of the task (YYYY-MM-DD HH:MM)")

	if err := flagSet.Parse(args); err != nil {
		return nil, nil, nil, err
	}
	return title, priorityInput, due, nil
}

func parseDue(due string) (*time.Time, error) {
	if due == "" {
		return nil, nil
	}
	parsedDate, err := time.Parse("2006-01-02 15:04", due)
	if err != nil {
		return nil, fmt.Errorf("invalid due date format: %v", err)
	}
	return &parsedDate, nil
}

func generateAddTaskResult(newTask task.Task) string {
	return fmt.Sprintf("✅ Task ID-%d is added: %s (Priority: %s, Due: %s)", newTask.ID, newTask.Title, newTask.Priority.Symbol(), newTask.Due)
}
