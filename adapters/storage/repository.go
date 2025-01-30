package storage

import (
	"encoding/json"
	"fmt"
	"github.com/denner-s/task-tracker-cli/core"
	"os"
)

type FileRepository struct {
	filename string
}

func NewFileRepository(filename string) *FileRepository {
	return &FileRepository{filename: filename}
}

func (r *FileRepository) AddTask(task core.Task) error {
	tasks, err := r.loadTasks()
	if err != nil {
		return err
	}
	tasks = append(tasks, task)
	return r.saveTasks(tasks)
}

func (r *FileRepository) UpdateTask(task core.Task) error {
	tasks, err := r.loadTasks()
	if err != nil {
		return err
	}
	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i] = task
			return r.saveTasks(tasks)
		}
	}
	return fmt.Errorf("task with ID %d not found", task.ID)
}

func (r *FileRepository) DeleteTask(id int) error {
	tasks, err := r.loadTasks()
	if err != nil {
		return err
	}
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return r.saveTasks(tasks)
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func (r *FileRepository) GetTaskByID(id int) (*core.Task, error) {
	tasks, err := r.loadTasks()
	if err != nil {
		return nil, err
	}
	for _, t := range tasks {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, fmt.Errorf("task with ID %d not found", id)
}

func (r *FileRepository) GetAllTasks() ([]core.Task, error) {
	return r.loadTasks()
}

func (r *FileRepository) loadTasks() ([]core.Task, error) {
	data, err := os.ReadFile(r.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []core.Task{}, nil
		}
		return nil, err
	}
	var tasks []core.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *FileRepository) saveTasks(tasks []core.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filename, data, 0644)
}
