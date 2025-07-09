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

var tasks []*Task

func AddTask(task *Task) {
	tasks = append(tasks, task)
}
