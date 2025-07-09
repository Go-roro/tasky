package task

import "testing"

func TestPriority_ValidPriority(t *testing.T) {
	tests := []struct {
		name string
		p    Priority
		want bool
	}{
		{"Low", PriorityLow, true},
		{"Medium", PriorityMedium, true},
		{"High", PriorityHigh, true},
		{"Invalid", Priority(0), false},
		{"Invalid", Priority(4), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.ValidPriority(); got != tt.want {
				t.Errorf("Priority.ValidPriority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriority_Symbol(t *testing.T) {
	tests := []struct {
		name string
		p    Priority
		want string
	}{
		{"Low", PriorityLow, "🟢"},
		{"Medium", PriorityMedium, "🟠"},
		{"High", PriorityHigh, "🔴"},
		{"Default", Priority(0), "🟠"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Symbol(); got != tt.want {
				t.Errorf("Priority.Symbol() = %v, want %v", got, tt.want)
			}
		})
	}
}
