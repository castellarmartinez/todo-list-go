package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Completed   bool   `json:"completed"`
}
