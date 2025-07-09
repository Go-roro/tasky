package task

import "sync/atomic"

var nextID int32
var tasks []*Task

func AddTask(task *Task) {
	task.ID = int(generateID())
	tasks = append(tasks, task)
}

func generateID() int32 {
	return atomic.AddInt32(&nextID, 1)
}
