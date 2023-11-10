package repository

import (
	"context"
	"errors"
	"job-application-api/internal/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) Viewjob(ctx context.Context, cid uint64) (models.Job, error) {
	var jobData models.Job
	result := r.DB.Preload("Location").
		Preload("TechnologyStack").
		Preload("Qualifications").
		Preload("Shift").
		Where("id = ?", cid).Find(&jobData)

	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Job{}, errors.New("could not create the jobs")
	}
	return jobData, nil
}

func (r *Repo) ViewJobPostings(ctx context.Context) ([]models.Job, error) {
	var jobDetails []models.Job
	result := r.DB.Find(&jobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find jobs")
	}
	return jobDetails, nil

}

func (r *Repo) ViewJobByCid(ctx context.Context, cid uint64) ([]models.Job, error) {
	var jobDetails []models.Job
	result := r.DB.Where("cid = ?", cid).Find(&jobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find job for the cid")
	}
	return jobDetails, nil

}

func (r *Repo) CreateJob(ctx context.Context, jobData models.Job) (models.ResponseJob, error) {
	result := r.DB.Create(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.ResponseJob{}, errors.New("could not create the job")
	}
	return models.ResponseJob{
		Id: jobData.ID,
	}, nil
}
