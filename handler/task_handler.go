package handler

import (
	"codebranch/service"
	"encoding/json"
	"net/http"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
