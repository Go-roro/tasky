package cmd

import (
	"github.com/yourusername/tasky/task"
	"testing"
)

func TestDeleteCommand_OnAction(t *testing.T) {
	tests := []struct {
		name  string
		tasks *task.Tasks
		args  []string
		want  string
	}{
		{
			name:  "Delete existing task",
			tasks: task.InitializeWithTasks([]*task.Task{{ID: 1, Title: "Test Task 1", Done: false}}),
			args:  []string{"--id", "1"},
			want:  "🗑️ Task with ID 1 has been deleted successfully",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &DeleteCommand{tasks: tt.tasks}
			result, err := c.OnAction(tt.args)
			if err != nil {
				t.Errorf("DeleteCommand.OnAction() error = %v", err)
				return
			}
			if result != tt.want {
				t.Errorf("unexpected result: got = %q, want = %q", result, tt.want)
			}
		})
	}
}

func TestDeleteCommand_OnAction_Invalid(t *testing.T) {
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
			want:  "failed to delete task with ID 999: task not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &DeleteCommand{tasks: tt.tasks}
			result, err := c.OnAction(tt.args)
			if err == nil || err.Error() != tt.want {
				t.Errorf("expected error = %q, got = %v", tt.want, err)
				return
			}
			if result != "" {
				t.Errorf("expected empty result, got = %q", result)
			}
		})
	}
}
