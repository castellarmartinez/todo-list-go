package data

import "codebranch/models"

var InitialTasks = []models.Task{
	{ID: 1, Title: "Clean the house", Description: "I have to do some cleaning weekly", Completed: false},
	{ID: 2, Title: "Workout", Description: "I have to be in shape", Completed: false},
}
