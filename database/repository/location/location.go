package location

import (
	"main/database/entity"

	"gorm.io/gorm"
)

type LocationRepository interface {
	FindEventByLocation(lati, longi float64) (*entity.Map, error)
	FindDexById(Id uint) (res uint, err error)
	SaveHistories(location []entity.Map) (err error)
}

type gormLocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &gormLocationRepository{db}
}

// 사건 조회
func (g *gormLocationRepository) FindEventByLocation(lati, longi float64) (res *entity.Map, err error) {
	// 1. 쿼리작성
	// 	select id, latitude , longitude
	//   from "map"
	//  where id = 1;

	// 2. gorm 적용
	tx := g.db
	err = tx.Model(&entity.Map{}).Select("id", "latitude", "longitude").Where("latitude = ? AND longitude = ?", lati, longi).First(&res).Error

	if err != nil {
		return
	}

	return

}

// 사건 조회 : isCollect 사건 보유 여부
func (g *gormLocationRepository) FindDexById(userId uint) (res uint, err error) {
	var collectCount int64
	tx := g.db
	err = tx.Model(&entity.Map{}).Where("id = ?", userId).Count(&collectCount).Error
	if err != nil {
		return
	}

	return
}

// 사건 위치 저장
func (g *gormLocationRepository) SaveHistories(location []entity.Map) (err error) {
	// 	INSERT
	//    INTO "map"(name, address, address_road , longitude , latitude)
	//  VALUES('우림이집', '수원', '세류로39', 127.00711632852955,  37.26275448320069);

	err = g.db.Create(location).Error

	return
}
