package location

import (
	"main/database/entity"

	"gorm.io/gorm"
)

type LocationRepository interface {
	GetByID(id uint) (*entity.Map, error)
}

type gormLocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &gormLocationRepository{db}
}

func (g *gormLocationRepository) GetByID(id uint) (*entity.Map, error) {
	panic("")
}
