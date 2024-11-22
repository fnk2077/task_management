package postgres

import (
	"assignment_task/models"
	"fmt"
	"os"

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

	if err := db.AutoMigrate(&models.User{}, models.Task{}); err != nil {
		panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}

	return &Postgres{Db: db}
}
