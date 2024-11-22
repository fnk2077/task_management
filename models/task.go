package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text"`
	Status      string    `json:"status" gorm:"type:varchar(50);default:'Pending'" `
	Priority    string    `json:"priority" gorm:"type:varchar(50);default:'Medium'"`
	DueDate     time.Time `json:"due_date"`
	AssignedTo  string    `json:"assigned_to" gorm:"type:varchar(200);not null"`
	CreatedBy   string    `json:"created_by" gorm:"type:varchar(200);not null"`
}
