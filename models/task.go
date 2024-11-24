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
	DueDate     time.Time `json:"dueDate"`
	AssignedTo  string    `json:"assignedTo" gorm:"type:varchar(200);not null"`
	CreatedBy   string    `json:"createdBy" gorm:"type:varchar(200);not null"`
}
