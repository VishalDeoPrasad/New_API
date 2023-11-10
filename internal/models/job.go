package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Company         Company           `json:"-" gorm:"ForeignKey:cid"`
	Cid             uint              `json:"cid"`
	Jobname         string            `json:"jobname"`
	MinNoticePeriod string            `json:"minNoticePeriod" validate:"required"`
	MaxNoticePeriod string            `json:"maxNoticePeriod" validate:"required"`
	Description     string            `json:"description" validate:"required"`
	MinExperience   string            `json:"minExperience" validate:"required"`
	MaxExperience   string            `json:"maxExperience" validate:"required"`
	Location        []Location        `json:"location" gorm:"many2many:job_location;"`
	TechnologyStack []TechnologyStack `json:"technologyStacks" gorm:"many2many:job_techstack;"`
	Qualifications  []Qualification   `json:"qualifications" gorm:"many2many:job_qualification;"`
	Shift           []Shift           `json:"shifts" gorm:"many2many:job_shift;" `
	Jobtype         string            `json:"jobType" validate:"required"`
}
type ResponseJob struct {
	Id uint `json:"id"`
}

type Location struct {
	gorm.Model
	PlaceName string `json:"placeName"`
}

type TechnologyStack struct {
	gorm.Model
	StackName string `json:"stackName"`
}

type Qualification struct {
	gorm.Model
	QualificationRequired string `json:"qualificationRequired"`
}

type Shift struct {
	gorm.Model
	ShiftType string `json:"shiftTypes"`
}

type NewJob struct {
	Company         Company `json:"-" `
	Cid             uint    `json:"cid"`
	Jobname         string  `json:"jobname"`
	MinNoticePeriod string  `json:"minNoticePeriod"`
	MaxNoticePeriod string  `json:"maxNoticePeriod"`
	Description     string  `json:"description"`
	MinExperience   string  `json:"minExperience"`
	MaxExperience   string  `json:"maxExperience"`
	Location        []uint  `json:"location"`
	TechnologyStack []uint  `json:"technologyStacks"`
	Qualifications  []uint  `json:"qualifications"`
	Shift           []uint  `json:"shifts"`
	Jobtype         string  `json:"jobType"`
}
