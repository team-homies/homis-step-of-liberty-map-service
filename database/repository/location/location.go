package location

import (
	"main/app/api/common"
	"main/database/entity"

	"gorm.io/gorm"
)

type LocationRepository interface {
	FindEventByPoint(MaxLatLon, MinLatLon common.Point) (res *entity.Map, err error)
}

type gormLocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &gormLocationRepository{db}
}

// 사건 조회
func (g *gormLocationRepository) FindEventByPoint(MaxLatLon, MinLatLon common.Point) (res *entity.Map, err error) {
	// 	select event_id, latitude , longitude
	//   from "map"
	//  where  latitude BETWEEN 최소위도 AND 최대위도 AND longitude  BETWEEN 최소경도 AND 최대경도

	// 1. gorm 적용
	err = g.db.Model(&entity.Map{}).Select("event_id", "latitude", "longitude").
		Where("latitude BETWEEN ? AND ? AND longitude BETWEEN ? AND ?", MinLatLon.Latitude, MaxLatLon.Latitude, MinLatLon.Longitude, MaxLatLon.Longitude).
		First(&res).Error

	if err != nil {
		return
	}

	return

}
