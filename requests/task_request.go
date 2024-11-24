package requests

import (
	"time"
)

type CreateTaskRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	DueDate     time.Time `json:"dueDate"`
	AssignedTo  string    `json:"assignedTo"`
	CreatedBy   string    `json:"createdBy"`
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
	DueDate     time.Time `json:"dueDate"`
	AssignedTo  string    `json:"assignedTo"`
	CreatedBy   string    `json:"createdBy"`
}
