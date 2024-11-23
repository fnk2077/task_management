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

func (p *Postgres) UpdateTask(task models.Task) error {
	if err := p.Db.Model(&models.Task{}).Where("id = ?", task.ID).Updates(&task).Error; err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}
	return nil
}

func (p *Postgres) DeleteTaskByID(taskID int) error {
	if err := p.Db.Delete(&models.Task{}, taskID).Error; err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}

func (p *Postgres) GetAllTask() (*[]models.Task, error) {
	var tasks []models.Task
	if err := p.Db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return &tasks, nil
}


func (p *Postgres) GetTaskByEmail(email string) (*models.Task, error) {
	var task models.Task
	p.Db.Where("email = ?", email).First(&task)
	return &task, nil
}
