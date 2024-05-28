package repository

import (
	"main/database"
	patient "main/database/repository/map"
)

type Repository interface {
	patient.MapRepository
}

func NewRepository() Repository {
	db := database.DB
	return &repository{
		patient.NewPatientRepository(db),
	}
}

type repository struct {
	patient.MapRepository
}
