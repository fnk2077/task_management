package builders

import (
	"assignment_task/models"
	"time"
)

type TaskBuilder struct {
	title       string
	description string
	status      string
	priority    string
	dueDate     time.Time
	assignedTo  string
	createdBy   string
}

func NewTaskBuilder() *TaskBuilder {
	return &TaskBuilder{}
}

func (taskBuilder *TaskBuilder) SetTitle(title string) (u *TaskBuilder) {
	taskBuilder.title = title
	return taskBuilder
}

func (taskBuilder *TaskBuilder) SetDescription(description string) (u *TaskBuilder) {
	taskBuilder.description = description
	return taskBuilder
}

func (taskBuilder *TaskBuilder) SetStatus(status string) (u *TaskBuilder) {
	taskBuilder.status = status
	return taskBuilder
}

func (taskBuilder *TaskBuilder) SetPriority(priority string) (u *TaskBuilder) {
	taskBuilder.priority = priority
	return taskBuilder
}

func (taskBuilder *TaskBuilder) SetDueDate(dueDate time.Time) (u *TaskBuilder) {
	taskBuilder.dueDate = dueDate
	return taskBuilder
}

func (taskBuilder *TaskBuilder) SetAssignedTo(assignedTo string) (u *TaskBuilder) {
	taskBuilder.assignedTo = assignedTo
	return taskBuilder
}

func (taskBuilder *TaskBuilder) SetCreatedBy(createdBy string) (u *TaskBuilder) {
	taskBuilder.createdBy = createdBy
	return taskBuilder
}

func (taskBuilder *TaskBuilder) Build() models.Task {
	task := models.Task{
		Title:       taskBuilder.title,
		Description: taskBuilder.description,
		Status:      taskBuilder.status,
		Priority:    taskBuilder.priority,
		DueDate:     taskBuilder.dueDate,
		AssignedTo:  taskBuilder.assignedTo,
		CreatedBy:   taskBuilder.createdBy,
	}

	return task
}
