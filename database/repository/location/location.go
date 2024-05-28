package location

import (
	"main/database/entity"

	"gorm.io/gorm"
)

type LocationRepository interface {
	FindEventByLocation(Id uint) (*entity.Map, error)
	FindDexById(Id uint) (*entity.Map, error)
}

type gormLocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &gormLocationRepository{db}
}

// 사건 조회
func (g *gormLocationRepository) FindEventByLocation(Id uint) (res *entity.Map, err error) {
	// 1. 쿼리작성
	// 	select id, latitude , longitude
	//   from "map"
	//  where id = 1;

	// 2. gorm 적용
	tx := g.db
	err = tx.Model(&entity.Map{}).Select("id", "latitude", "longitude").Where("id = ?", Id).First(&res).Error

	if err != nil {
		return
	}

	return

}

// 사건 조회 : isCollect 사건 보유 여부
func (g *gormLocationRepository) FindDexById(id uint) (*entity.Map, error) {
	panic("")
}
