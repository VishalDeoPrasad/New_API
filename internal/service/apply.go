package service

import (
	"context"
	"errors"
	"fmt"
	"job-application-api/internal/models"
	"sync"
)

func (s *Service) SelectApplications(ctx context.Context, jobApply []models.RespondJApplicant) ([]models.RespondJApplicant, error) {
	var SelectedJobData []models.RespondJApplicant
	jobdetail, err := s.UserRepo.Viewjob(ctx, uint64(1))
	fmt.Println("[][]]]", jobdetail, "[][][][[]]")
	if err != nil {
		return nil, errors.New("problem in fetching jobs from database")
	}

	ch := make(chan models.RespondJApplicant)
	wg := new(sync.WaitGroup)

	for _, val := range jobApply {
		wg.Add(1)
		go func(val models.RespondJApplicant) {
			defer wg.Done()
			bool, singleApplication := matchApplicants(val, jobdetail)
			if bool {
				ch <- singleApplication
			}

		}(val)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for data := range ch {
		SelectedJobData = append(SelectedJobData, data)
	}

	return SelectedJobData, nil

}

func matchApplicants(form models.RespondJApplicant, jobdetail models.Job) (bool, models.RespondJApplicant) {
	if form.Jobs.Experience != jobdetail.MaxExperience {
		return false, models.RespondJApplicant{}
	}

	if form.Jobs.Jobtype != jobdetail.Jobtype {
		return false, models.RespondJApplicant{}
	}

	if form.Jobs.NoticePeriod != jobdetail.MaxNoticePeriod {
		return false, models.RespondJApplicant{}
	}
	// MinNoticePeriod, err := strconv.Atoi(jobdetail.MinNoticePeriod)
	// if err != nil {
	// 	return false, models.RespondJApplicant{}
	// }
	// MaxNoticePeriod, err := strconv.Atoi(jobdetail.MaxNoticePeriod)
	// if err != nil {
	// 	return false, models.RespondJApplicant{}
	// }
	// applicantNoticePeriod, err := strconv.Atoi(v.Jobs.NoticePeriod)
	// if err != nil {
	// 	return false, models.RespondJApplicant{}
	// }

	// if (applicantNoticePeriod < MinNoticePeriod) || (applicantNoticePeriod > MaxNoticePeriod) {
	// 	return false, models.RespondJApplicant{}
	// }
	// MinExperience, err := strconv.Atoi(jobdetail.MinExperience)
	// if err != nil {
	// 	return false, models.RespondJApplicant{}
	// }
	// MaxExperience, err := strconv.Atoi(jobdetail.MaxExperience)
	// if err != nil {
	// 	return false, models.RespondJApplicant{}
	// }
	// applicantExperience, err := strconv.Atoi(v.Jobs.Experience)
	// if err != nil {
	// 	return false, models.RespondJApplicant{}

	// }
	// if (applicantExperience < MinExperience) || (applicantExperience > MaxExperience) {
	// 	return false, models.RespondJApplicant{}
	// }

	count := 0
	for _, v1 := range form.Jobs.Location {
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
	for _, v1 := range form.Jobs.Qualifications {
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

	// count = 0
	// for _, v1 := range form.Jobs.Shift {
	// 	count = 0
	// 	for _, v2 := range jobdetail.Shift {
	// 		if v1 == v2.ID {
	// 			count++
	// 		}

	// 	}
	// }
	// if count == 0 {
	// 	return false, models.RespondJApplicant{}
	// }

	// count = 0
	// for _, v1 := range form.Jobs.TechnologyStack {
	// 	count = 0
	// 	for _, v2 := range jobdetail.TechnologyStack {
	// 		if v1 == v2.ID {
	// 			count++
	// 		}

	// 	}
	// }
	return true, form
}
