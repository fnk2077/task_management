package postgres

import (
	"assignment_task/models"
	"fmt"
)

func (p *Postgres) CreateTask(task models.Task) error {
	if err := p.Db.Create(&task).Error; err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}
	return nil
}

func (p *Postgres) EditTask(task models.Task) error {
	if err := p.Db.Model(&models.Task{}).Where("id = ?", task.ID).Updates(&task).Error; err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}
	return nil
}

func (p *Postgres) DeleteTask(taskID int) error {
	if err := p.Db.Delete(&models.Task{}, taskID).Error; err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}

func (p *Postgres) GetTaskByEmail(email string) (*models.Task, error) {
	var task models.Task
	p.Db.Where("email = ?", email).First(&task)
	return &task, nil
}
