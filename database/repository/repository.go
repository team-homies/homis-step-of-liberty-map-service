package repository

import (
	"main/database"
	"main/database/repository/patient"
)

type Repository interface {
	patient.PatientRepository
}

func NewRepository() Repository {
	db := database.DB
	return &repository{
		patient.NewPatientRepository(db),
	}
}

type repository struct {
	patient.PatientRepository
}
