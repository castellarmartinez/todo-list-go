package repository

import (
	"codebranch/data"
	"codebranch/models"
	"errors"
)

type TaskRepository struct {
	tasks  map[int]models.Task
	nextID int
}

func NewTaskRepository() *TaskRepository {
	tasks := make(map[int]models.Task)

	for _, task := range data.InitialTasks {
		tasks[task.ID] = task
	}

	return &TaskRepository{
		tasks:  tasks,
		nextID: len(data.InitialTasks) + 1,
	}
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	var taskList []models.Task

	for _, task := range r.tasks {
		taskList = append(taskList, task)
	}

	return taskList, nil
}

func (r *TaskRepository) GetByID(id int) (*models.Task, error) {
	task, exists := r.tasks[id]

	if !exists {
		return nil, errors.New("task not found")
	}

	return &task, nil
}

func (r *TaskRepository) Create(task models.Task) (models.Task, error) {
	task.ID = r.nextID
	r.tasks[r.nextID] = task
	r.nextID++

	return task, nil
}

func (r *TaskRepository) Update(id int, updatedTask models.Task) (models.Task, error) {
	currentTask, exists := r.tasks[id]

	if !exists {
		return models.Task{}, errors.New("task not found")
	}

	if updatedTask.Title != "" {
		currentTask.Title = updatedTask.Title
	}

	if updatedTask.Description != "" {
		currentTask.Description = updatedTask.Description
	}

	currentTask.Completed = updatedTask.Completed

	r.tasks[id] = currentTask
	return currentTask, nil
}

func (r *TaskRepository) Delete(id int) error {
	_, exists := r.tasks[id]

	if !exists {
		return errors.New("task not found")
	}

	delete(r.tasks, id)
	return nil
}
