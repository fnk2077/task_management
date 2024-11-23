package task

import (
	"assignment_task/builders"
	"assignment_task/models"
	"assignment_task/requests"
	"assignment_task/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

func New(db Storer) *Handler {
	return &Handler{store: db}
}

type Handler struct {
	store Storer
}

type Storer interface {
	CreateTask(models.Task) error
	GetAllTask() (*[]models.Task, error)
	GetTaskByEmail(email string) (*[]models.Task, error)
	DeleteTaskByID(taskID int) error
	UpdateTask(task models.Task) error
}

func (h *Handler) CraeteTask(c echo.Context) error {
	taskRequest := new(requests.CreateTaskRequest)

	if err := c.Bind(taskRequest); err != nil {
		return err
	}

	newTask := builders.NewTaskBuilder().SetTitle(taskRequest.Title).
		SetDescription(taskRequest.Description).
		SetPriority(taskRequest.Priority).
		SetStatus(taskRequest.Status).
		SetDueDate(taskRequest.DueDate).
		SetAssignedTo(taskRequest.AssignedTo).
		SetCreatedBy(taskRequest.CreatedBy).
		Build()

	h.store.CreateTask(newTask)

	return responses.MessageResponse(c, http.StatusCreated, "Task successfully created")
}

func (h *Handler) GetTask(c echo.Context) error {
	email := c.QueryParam("email")

	if email != "" {
		res, err := h.store.GetTaskByEmail(email)
		if err != nil {
			return err
		}
		return responses.Response(c, http.StatusOK, res)
	}

	res, err := h.store.GetAllTask()
	if err != nil {
		return err
	}

	return responses.Response(c, http.StatusOK, res)
}

func (h *Handler) UpdateTask(c echo.Context) error {
	taskRequest := new(requests.UpdateTaskRequest)

	if err := c.Bind(taskRequest); err != nil {
		return err
	}
	newTask := builders.NewTaskBuilder().SetTitle(taskRequest.Title).
		SetDescription(taskRequest.Description).
		SetPriority(taskRequest.Priority).
		SetStatus(taskRequest.Status).
		SetDueDate(taskRequest.DueDate).
		SetAssignedTo(taskRequest.AssignedTo).
		SetCreatedBy(taskRequest.CreatedBy).
		Build()

	newTask.ID = uint(taskRequest.ID)
	h.store.UpdateTask(newTask)
	return responses.MessageResponse(c, http.StatusOK, "Task successfully update")
}

func (h *Handler) DeleteTaskByID(c echo.Context) error {
	taskRequest := new(requests.DeleteTaskRequest)

	if err := c.Bind(taskRequest); err != nil {
		return err
	}
	h.store.DeleteTaskByID(taskRequest.ID)
	return responses.MessageResponse(c, http.StatusOK, "Task successfully deleted")
}
