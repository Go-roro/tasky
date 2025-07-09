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
		ID:        0, // ID will be set by the database or task manager
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
	return fmt.Sprintf("✅ Task added: %s (Priority: %s, Due: %s)", newTask.Title, newTask.Priority.Symbol(), newTask.Due)
}
