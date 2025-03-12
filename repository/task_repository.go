package repository

import (
	"codebranch/data"
	"codebranch/models"
)

type TaskRepository struct {
	tasks map[int]models.Task
}

func NewTaskRepository() *TaskRepository {
	tasks := make(map[int]models.Task)

	for _, task := range data.InitialTasks {
		tasks[task.ID] = task
	}

	return &TaskRepository{
		tasks: tasks,
	}
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	var taskList []models.Task

	for _, task := range r.tasks {
		taskList = append(taskList, task)
	}

	return taskList, nil
}
