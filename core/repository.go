package core

type TaskRepository interface {
	AddTask(task Task) error
	UpdateTask(task Task) error
	DeleteTask(id int) error
	GetTaskByID(id int) (*Task, error)
	GetAllTasks() ([]Task, error)
}
