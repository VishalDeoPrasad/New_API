package service

import (
	"context"
	"job-application-api/internal/models"

	"gorm.io/gorm"
)

func (s *Service) AddJobDetails(ctx context.Context, jobData models.NewJob) (models.ResponseJob, error) {

	createjobdetails := models.Job{
		Cid:             jobData.Cid,
		Jobname:         jobData.Jobname,
		MinExperience:   jobData.MinExperience,
		MaxExperience:   jobData.MaxExperience,
		MinNoticePeriod: jobData.MinNoticePeriod,
		MaxNoticePeriod: jobData.MaxNoticePeriod,
		Jobtype:         jobData.Jobtype,
		Description:     jobData.Description,
	}

	for _, v := range jobData.Location {
		tempCreateJobDetails := models.Location{
			Model: gorm.Model{
				ID: v,
			},
		}
		createjobdetails.Location = append(createjobdetails.Location, tempCreateJobDetails)
	}

	for _, v := range jobData.Qualifications {
		tempCreateJobDetails := models.Qualification{
			Model: gorm.Model{
				ID: v,
			},
		}
		createjobdetails.Qualifications = append(createjobdetails.Qualifications, tempCreateJobDetails)
	}
	for _, v := range jobData.Shift {
		tempCreateJobDetails := models.Shift{
			Model: gorm.Model{
				ID: v,
			},
		}
		createjobdetails.Shift = append(createjobdetails.Shift, tempCreateJobDetails)
	}
	for _, v := range jobData.TechnologyStack {
		tempCreateJobDetails := models.TechnologyStack{
			Model: gorm.Model{
				ID: v,
			},
		}
		createjobdetails.TechnologyStack = append(createjobdetails.TechnologyStack, tempCreateJobDetails)
	}

	job, err := s.UserRepo.CreateJob(ctx, createjobdetails)
	if err != nil {
		return models.ResponseJob{}, err
	}
	return job, nil
}
func (s *Service) ViewJobDetailsById(ctx context.Context, cid uint64) (models.Job, error) {
	jobData, err := s.UserRepo.Viewjob(ctx, cid)
	if err != nil {
		return models.Job{}, err
	}
	return jobData, nil

}

func (s *Service) ViewAllJobPostings(ctx context.Context) ([]models.Job, error) {
	jobData, err := s.UserRepo.ViewJobPostings(ctx)
	if err != nil {
		return nil, err
	}
	return jobData, nil

}

func (s *Service) ViewJobDetails(ctx context.Context, cid uint64) ([]models.Job, error) {
	jobData, err := s.UserRepo.ViewJobByCid(ctx, cid)
	if err != nil {
		return nil, err
	}
	return jobData, nil
}
