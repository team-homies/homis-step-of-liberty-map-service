package repository

import (
	"main/database"
	location "main/database/repository/location"
)

type Repository interface {
	location.LocationRepository
}

func NewRepository() Repository {
	db := database.DB
	return &repository{
		location.NewLocationRepository(db),
	}
}

type repository struct {
	location.LocationRepository
}
