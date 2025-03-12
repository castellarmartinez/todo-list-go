package service

import (
	"codebranch/models"
	"codebranch/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTaskService(t *testing.T) {
	repo := repository.NewTaskRepository()
	service := NewTaskService(repo)

	assert.NotNil(t, service)
	assert.NotNil(t, service.repo)
}

func TestGetAll(t *testing.T) {
	repo := repository.NewTaskRepository()
	service := NewTaskService(repo)

	tasks, err := service.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, tasks)
	assert.GreaterOrEqual(t, len(tasks), 0)
}

func TestGetTaskByID_Found(t *testing.T) {
	repo := repository.NewTaskRepository()
	service := NewTaskService(repo)

	task, err := service.GetTaskByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, 1, task.ID)
}

func TestGetTaskByID_NotFound(t *testing.T) {
	repo := repository.NewTaskRepository()
	service := NewTaskService(repo)

	task, err := service.GetTaskByID(999)
	assert.Error(t, err)
	assert.Nil(t, task)
}

func TestCreateTask(t *testing.T) {
	repo := repository.NewTaskRepository()
	service := NewTaskService(repo)

	newTask := models.Task{
		Title:       "New Task",
		Description: "New description",
		Completed:   false,
	}

	createdTask, err := service.CreateTask(newTask)
	assert.NoError(t, err)
	assert.NotZero(t, createdTask.ID)
	assert.Equal(t, newTask.Title, createdTask.Title)
}

func TestUpdateTask_Found(t *testing.T) {
	repo := repository.NewTaskRepository()
	service := NewTaskService(repo)

	updatedTask := models.Task{
		Title:       "Updated Title",
		Description: "Updated Description",
		Completed:   true,
	}

	task, err := service.UpdateTask(1, updatedTask)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Title", task.Title)
	assert.Equal(t, "Updated Description", task.Description)
	assert.True(t, task.Completed)
}

func TestUpdateTask_NotFound(t *testing.T) {
	repo := repository.NewTaskRepository()
	service := NewTaskService(repo)

	updatedTask := models.Task{
		Title: "Updated Task",
	}

	task, err := service.UpdateTask(999, updatedTask)
	assert.Error(t, err)
	assert.Equal(t, models.Task{}, task)
}

func TestDeleteTask_Found(t *testing.T) {
	repo := repository.NewTaskRepository()
	service := NewTaskService(repo)

	err := service.DeleteTask(1)
	assert.NoError(t, err)

	_, err = service.GetTaskByID(1)
	assert.Error(t, err)
}

func TestDeleteTask_NotFound(t *testing.T) {
	repo := repository.NewTaskRepository()
	service := NewTaskService(repo)

	err := service.DeleteTask(999)
	assert.Error(t, err)
}
