package service

import (
	"codebranch/models"
	"codebranch/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAll() ([]models.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) GetTaskByID(id int) (*models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) CreateTask(task models.Task) (models.Task, error) {
	return s.repo.Create(task)
}

func (s *TaskService) UpdateTask(id int, task models.Task) (models.Task, error) {
	return s.repo.Update(id, task)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.Delete(id)
}
