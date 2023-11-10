package repository

import (
	"context"
	"errors"
	"job-application-api/internal/models"

	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=repository_mock.go -package=repository

type Repo struct {
	DB *gorm.DB
}
type UserRepo interface {
	CreateUser(userData models.User) (models.User, error)
	CheckEmail(ctx context.Context, email string) (models.User, error)
	CreateCompany(ctx context.Context, companyData models.Company) (models.Company, error)
	ViewCompany(ctx context.Context, cid uint64) (models.Company, error)
	ViewAllCompanies(ctx context.Context) ([]models.Company, error)
	CreateJob(ctx context.Context, jobData models.Job) (models.ResponseJob, error)
	ViewJobByCid(ctx context.Context, cid uint64) ([]models.Job, error)
	ViewJobPostings(ctx context.Context) ([]models.Job, error)
	Viewjob(ctx context.Context, cid uint64) (models.Job, error)
}

func NewRepository(db *gorm.DB) (UserRepo, error) {
	if db == nil {
		return nil, errors.New("db cannot be null")
	}
	return &Repo{
		DB: db,
	}, nil
}
