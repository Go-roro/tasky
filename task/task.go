package task

import "time"

type Task struct {
	ID        int
	Title     string
	Priority  Priority
	Due       *time.Time
	Done      bool
	CreatedAt time.Time
}

func (t *Task) MarkAsDone() {
	if t.Done {
		return
	}
	t.Done = true
}
