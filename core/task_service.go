package core

import (
	"github.com/denner-s/task-tracker-cli/ports"
	"time"
)

type TaskService struct {
	repo ports.TaskRepository
}

func NewTaskService(repo ports.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) AddTask(description string) (*Task, error) {
	tasks, err := s.repo.GetAllTasks()
	if err != nil {
		return nil, err
	}
	newID := 1
	if len(tasks) > 0 {
		for _, t := range tasks {
			if t.ID >= newID {
				newID = t.ID + 1
			}
		}
	}
	task := Task{
		ID:          newID,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := s.repo.AddTask(task); err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) UpdateTask(id int, description string) (*Task, error) {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	task.Description = description
	task.UpdatedAt = time.Now()
	if err := s.repo.UpdateTask(*task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.DeleteTask(id)
}

func (s *TaskService) MarkTaskInProgress(id int) (*Task, error) {
	return s.UpdateStatus(id, StatusInProgress)
}

func (s *TaskService) MarkTaskDone(id int) (*Task, error) {
	return s.UpdateStatus(id, StatusDone)
}

func (s *TaskService) UpdateStatus(id int, status TaskStatus) (*Task, error) {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	task.Status = status
	task.UpdatedAt = time.Now()
	if err := s.repo.UpdateTask(*task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) ListTasks(filter string) ([]Task, error) {
	tasks, err := s.repo.GetAllTasks()
	if err != nil {
		return nil, err
	}
	if filter == "" {
		return tasks, nil
	}
	var filtered []Task
	for _, t := range tasks {
		if string(t.Status) == filter {
			filtered = append(filtered, t)
		}
	}
	return filtered, nil
}
