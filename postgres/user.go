package postgres

import (
	"assignment_task/models"
	"fmt"
)

func (p *Postgres) RegisterUser(user models.User) error {
	if err := p.Db.Create(&user).Error; err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	return nil
}

func (p *Postgres) GetUserByEmail(email string) models.User {
	var user models.User
	p.Db.Where("email = ?", email).Find(&user)
	return user
}
