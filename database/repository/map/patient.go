package patient

import (
	"main/database/entity"

	"gorm.io/gorm"
)

type MapRepository interface {
	GetByID(id uint) (*entity.Patient, error)
}

type gormMapRepository struct {
	db *gorm.DB
}

func NewMapRepository(db *gorm.DB) MapRepository {
	return &gormMapRepository{db}
}

func (g *gormMapRepository) GetByID(id uint) (*entity.Patient, error) {
	panic("")
}
