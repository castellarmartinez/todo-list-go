package main

import (
	"codebranch/handler"
	"codebranch/repository"
	"codebranch/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	rp := repository.NewTaskRepository()
	sv := service.NewTaskService(rp)
	hd := handler.NewTaskHandler(sv)
	rt := chi.NewRouter()

	rt.Route("/tasks", func(r chi.Router) {
		r.Get("/", hd.GetAllTasks)
		r.Get("/{id}", hd.GetTaskByID)
		r.Post("/", hd.CreateTask)
		r.Put("/{id}", hd.UpdateTask)
		r.Delete("/{id}", hd.DeleteTask)
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", rt))
}
