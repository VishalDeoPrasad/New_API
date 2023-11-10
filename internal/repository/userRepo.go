package repository

import (
	"context"
	"errors"
	"job-application-api/internal/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateUser(userData models.User) (models.User, error) {
	result := r.DB.Create(&userData)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return userData, nil
}
func (r *Repo) CheckEmail(ctx context.Context, email string) (models.User, error) {
	var userDetails models.User
	result := r.DB.Where("email = ?", email).First(&userDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.User{}, errors.New("email not found")
	}
	return userDetails, nil
}
