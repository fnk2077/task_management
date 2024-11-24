package postgres

import (
	"assignment_task/models"
	"fmt"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Db *gorm.DB
}

func New() *Postgres {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	if err := db.AutoMigrate(&models.User{},models.Task{}); err != nil {
		panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}

	user1 := models.User{
		Email:    "a1@a.com",
		Password: "$2a$10$g/2qT.2URShprAteorbdKez1QdPwFoQcZMvysjAJ0UF0WPqWtJp6S",
		Name:     "a1",
	}
	
	user2 := models.User{
		Email:    "test01@a.com",
		Password: "$2a$10$5jCZSH4/AfHRNZjjz6ClruNr/RAbPOYnyckVOyb/mTBsIaBL8xIeG",
		Name:     "test01",
	}
	
	task1 := models.Task{
		Title:       "Sleep",
		Description: "Go Sleep",
		Status:      "Pending",
		Priority:    "Low",
		DueDate:     time.Now().Add(time.Hour),
		AssignedTo:  "a1@a.com",
		CreatedBy:   "test01@a.com",
	}
	
	task2 := models.Task{
		Title:       "Eat",
		Description: "Eat rice",
		Status:      "Pending",
		Priority:    "Low",
		DueDate:     time.Now().Add(time.Hour),
		AssignedTo:  "test01@a.com",
		CreatedBy:   "a1@a.com",
	}

	if err := db.FirstOrCreate(&user1, models.User{Email: user1.Email}).Error; err != nil {
		panic(fmt.Sprintf("Failed to create or find user1: %v", err))
	}
	
	if err := db.FirstOrCreate(&user2, models.User{Email: user2.Email}).Error; err != nil {
		panic(fmt.Sprintf("Failed to create or find user2: %v", err))
	}
	
	if err := db.FirstOrCreate(&task1, models.Task{Title: task1.Title, AssignedTo: task1.AssignedTo}).Error; err != nil {
		panic(fmt.Sprintf("Failed to create or find task1: %v", err))
	}
	
	if err := db.FirstOrCreate(&task2, models.Task{Title: task2.Title, AssignedTo: task2.AssignedTo}).Error; err != nil {
		panic(fmt.Sprintf("Failed to create or find task2: %v", err))
	}
	

	return &Postgres{Db: db}
}
