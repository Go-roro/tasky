package task

import (
	"errors"
	"sync/atomic"
)

type Tasks struct {
	values []*Task
	nextID atomic.Int32
}

var StaticTasks = InitializeTasks()

func (t *Tasks) AddTask(task *Task) {
	task.ID = int(t.generateID())
	t.values = append(t.values, task)
}

func (t *Tasks) generateID() int32 {
	return t.nextID.Add(1)
}

func (t *Tasks) FindTaskByID(id int) (*Task, error) {
	for _, task := range t.values {
		if task.ID == id {
			return task, nil
		}
	}
	return nil, errors.New("task not found")
}

func (t *Tasks) DeleteTaskByID(id int) error {
	for i, task := range t.values {
		if task.ID == id {
			t.values = append(t.values[:i], t.values[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

func InitializeTasks() *Tasks {
	return &Tasks{
		values: make([]*Task, 0),
		nextID: atomic.Int32{},
	}
}

func InitializeWithTasks(tasks []*Task) *Tasks {
	initializeTasks := InitializeTasks()
	for _, task := range tasks {
		initializeTasks.AddTask(task)
	}
	return initializeTasks
}
