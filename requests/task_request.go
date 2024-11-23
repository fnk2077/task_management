package requests

import (
	"time"
)

type CreateTaskRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	DueDate     time.Time `json:"due_date"`
	AssignedTo  string    `json:"assigned_to"`
	CreatedBy   string    `json:"created_by"`
}

type DeleteTaskRequest struct {
	ID int `json:"id"`
}

type UpdateTaskRequest struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	DueDate     time.Time `json:"due_date"`
	AssignedTo  string    `json:"assigned_to"`
	CreatedBy   string    `json:"created_by"`
}
