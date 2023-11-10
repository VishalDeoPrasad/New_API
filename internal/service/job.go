package service

import (
	"context"
	"errors"
	"job-application-api/internal/models"
	"strconv"
	"sync"

	"gorm.io/gorm"
)

func (s *Service) FilterApplications(ctx context.Context, jobApplication []models.RespondJApplicant) ([]models.RespondJApplicant, error) {
	var FilterJobData []models.RespondJApplicant
	jobdetail, err := s.UserRepo.Viewjob(ctx, uint64(1))
	if err != nil {
		return nil, errors.New("not able to get  jobs from database")
	}

	ch := make(chan models.RespondJApplicant)
	wg := new(sync.WaitGroup)

	for _, v := range jobApplication {
		wg.Add(1)
		go func(v models.RespondJApplicant) {
			defer wg.Done()
			bool, singleApplication := checkApplicantsCriteria(v, jobdetail)
			if bool {
				ch <- singleApplication
			}

		}(v)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for data := range ch {
		FilterJobData = append(FilterJobData, data)
	}

	return FilterJobData, nil

}
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
func checkApplicantsCriteria(v models.RespondJApplicant, jobdetail models.Job) (bool, models.RespondJApplicant) {
	MinNoticePeriod, err := strconv.Atoi(jobdetail.MinNoticePeriod)
	if err != nil {
		return false, models.RespondJApplicant{}
	}
	MaxNoticePeriod, err := strconv.Atoi(jobdetail.MaxNoticePeriod)
	if err != nil {
		return false, models.RespondJApplicant{}
	}
	applicantNoticePeriod, err := strconv.Atoi(v.Jobs.NoticePeriod)
	if err != nil {
		return false, models.RespondJApplicant{}
	}

	if (applicantNoticePeriod < MinNoticePeriod) || (applicantNoticePeriod > MaxNoticePeriod) {
		return false, models.RespondJApplicant{}
	}
	MinExperience, err := strconv.Atoi(jobdetail.MinExperience)
	if err != nil {
		return false, models.RespondJApplicant{}
	}
	MaxExperience, err := strconv.Atoi(jobdetail.MaxExperience)
	if err != nil {
		return false, models.RespondJApplicant{}
	}
	applicantExperience, err := strconv.Atoi(v.Jobs.Experience)
	if err != nil {
		return false, models.RespondJApplicant{}

	}
	if (applicantExperience < MinExperience) || (applicantExperience > MaxExperience) {
		return false, models.RespondJApplicant{}
	}

	count := 0
	for _, v1 := range v.Jobs.Location {
		count = 0
		for _, v2 := range jobdetail.Location {
			if v1 == v2.ID {
				count++

			}
		}
	}
	if count == 0 {
		return false, models.RespondJApplicant{}
	}

	if count == 0 {
		return false, models.RespondJApplicant{}
	}
	count = 0
	for _, v1 := range v.Jobs.Qualifications {
		count = 0
		for _, v2 := range jobdetail.Qualifications {
			if v1 == v2.ID {
				count++
			}

		}
	}
	if count == 0 {
		return false, models.RespondJApplicant{}
	}

	count = 0
	for _, v1 := range v.Jobs.Shift {
		count = 0
		for _, v2 := range jobdetail.Shift {
			if v1 == v2.ID {
				count++
			}

		}
	}
	if count == 0 {
		return false, models.RespondJApplicant{}
	}

	count = 0
	for _, v1 := range v.Jobs.TechnologyStack {
		count = 0
		for _, v2 := range jobdetail.TechnologyStack {
			if v1 == v2.ID {
				count++
			}

		}
	}
	return true, v
}
