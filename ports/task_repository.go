package ports

import "github.com/denner-s/task-tracker-cli/core"

type TaskRepository interface {
	AddTask(task core.Task) error
	UpdateTask(task core.Task) error
	DeleteTask(id int) error
	GetTaskByID(id int) (*core.Task, error)
	GetAllTasks() ([]core.Task, error)
}
