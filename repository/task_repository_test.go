package repository

import (
	"codebranch/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTaskRepository(t *testing.T) {
	repo := NewTaskRepository()

	assert.NotNil(t, repo)
	assert.GreaterOrEqual(t, len(repo.tasks), 0)
}

func TestGetAll(t *testing.T) {
	repo := NewTaskRepository()

	tasks, err := repo.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, tasks)
	assert.GreaterOrEqual(t, len(tasks), 0)
}

func TestGetByID_Found(t *testing.T) {
	repo := NewTaskRepository()

	task, err := repo.GetByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, 1, task.ID)
}

func TestGetByID_NotFound(t *testing.T) {
	repo := NewTaskRepository()

	task, err := repo.GetByID(999)
	assert.Error(t, err)
	assert.Nil(t, task)
}

func TestCreateTask(t *testing.T) {
	repo := NewTaskRepository()

	newTask := models.Task{
		Title:       "New Task",
		Description: "Test description",
		Completed:   false,
	}

	createdTask, err := repo.Create(newTask)
	assert.NoError(t, err)
	assert.NotZero(t, createdTask.ID)
	assert.Equal(t, newTask.Title, createdTask.Title)
}

func TestUpdateTask_Found(t *testing.T) {
	repo := NewTaskRepository()

	updatedTask := models.Task{
		Title:       "Updated Title",
		Description: "Updated Description",
		Completed:   true,
	}

	task, err := repo.Update(1, updatedTask)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Title", task.Title)
	assert.Equal(t, "Updated Description", task.Description)
	assert.True(t, task.Completed)
}

func TestUpdateTask_NotFound(t *testing.T) {
	repo := NewTaskRepository()

	updatedTask := models.Task{
		Title: "Updated Task",
	}

	task, err := repo.Update(999, updatedTask)
	assert.Error(t, err)
	assert.Equal(t, models.Task{}, task)
}

func TestDeleteTask_Found(t *testing.T) {
	repo := NewTaskRepository()

	err := repo.Delete(1)
	assert.NoError(t, err)

	_, err = repo.GetByID(1)
	assert.Error(t, err)
}

func TestDeleteTask_NotFound(t *testing.T) {
	repo := NewTaskRepository()

	err := repo.Delete(999)
	assert.Error(t, err)
}
