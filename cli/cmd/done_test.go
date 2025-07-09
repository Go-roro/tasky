package cmd

import (
	"github.com/yourusername/tasky/task"
	"testing"
)

func TestDoneCommand_OnAction(t *testing.T) {
	tests := []struct {
		name  string
		tasks *task.Tasks
		args  []string
		want  string
	}{
		{
			name: "Valid command with task ID",
			tasks: task.InitializeWithTasks(
				[]*task.Task{
					{
						Title:    "Test Task 1",
						Priority: task.PriorityHigh,
						Done:     false,
					},
				}),
			args: []string{"--id", "1"},
			want: "✅ Task 1 marked as done: Test Task 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &DoneCommand{tasks: tt.tasks}
			result, err := c.OnAction(tt.args)
			if err != nil {
				t.Errorf("DoneCommand.OnAction() error = %v", err)
				return
			}
			if result != tt.want {
				t.Errorf("unexpected result: got = %q, want = %q", result, tt.want)
			}
		})
	}
}

func TestDoneCommand_OnAction_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		tasks *task.Tasks
		args  []string
		want  string
	}{
		{
			name:  "Missing task ID",
			tasks: task.StaticTasks,
			args:  []string{},
			want:  "task ID is required",
		},
		{
			name:  "Invalid task ID",
			tasks: task.StaticTasks,
			args:  []string{"--id", "999"},
			want:  "task not found",
		},
		{
			name:  "Negative task ID",
			tasks: task.StaticTasks,
			args:  []string{"--id", "-1"},
			want:  "invalid task ID: -1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &DoneCommand{tasks: tt.tasks}
			_, err := c.OnAction(tt.args)
			if err == nil || err.Error() != tt.want {
				t.Errorf("DoneCommand.OnAction() expected error, got nil or different error: %v", err)
				return
			}
		})
	}
}
