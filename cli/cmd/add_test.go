package cmd

import (
	"strings"
	"testing"
)

func TestAddCommand_OnAction(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "Valid command with all options",
			args: []string{"--title", "Test Task", "--priority", "1", "--due", "2023-10-01 12:00"},
			want: "✅ Task added",
		},
		{
			name: "Valid command with title and priority only",
			args: []string{"--title", "Test Task", "--priority", "1"},
			want: "✅ Task added",
		},
		{
			name: "Valid command with title only",
			args: []string{"--title", "Test Task"},
			want: "✅ Task added",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &AddCommand{}
			result, err := c.OnAction(tt.args)
			if err != nil {
				t.Errorf("AddCommand.OnAction() error = %v, want nil", err)
			}
			if result == "" || !strings.Contains(result, tt.want) {
				t.Errorf("AddCommand.OnAction() got = %v, want substring %v", result, tt.want)
			}
		})
	}
}

func TestAddCommand_OnAction_Invalid(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		errorMessage string
	}{
		{
			name:         "Missing title",
			args:         []string{"--priority", "1"},
			errorMessage: "title is required",
		},
		{
			name:         "Invalid priority",
			args:         []string{"--title", "Test Task", "--priority", "4"},
			errorMessage: "priority must be between 1 and 3",
		},
		{
			name:         "Invalid due date format",
			args:         []string{"--title", "Test Task", "--due", "invalid-date"},
			errorMessage: "invalid due date format",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &AddCommand{}
			_, err := c.OnAction(tt.args)
			if err == nil || !strings.Contains(err.Error(), tt.errorMessage) {
				t.Errorf("AddCommand.OnAction() expected error, got nil")
				return
			}
		})
	}
}
