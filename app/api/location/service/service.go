package service

import (
	"main/app/api/common"
	"main/app/api/location/resource"
	"main/database/repository"
)

type LocationService interface {
	FindEvent(userId uint, req *resource.FindEventRequest) (res *resource.FindEventResponse, err error)
}

func NewLocationService() LocationService {
	return &locationService{
		LocationService: &locationService{},
	}
}

type locationService struct {
	LocationService
}

func (s *locationService) FindEvent(userId uint, req *resource.FindEventRequest) (res *resource.FindEventResponse, err error) {
	locationRepository := repository.NewRepository()

	MaxLatLon, MinLatLon := common.CalculateLatLonRange(req.Latitude, req.Longitude, common.RadiusKm)

	point, err := locationRepository.FindEventByPoint(MaxLatLon, MinLatLon)

	userlist, err := common.GetIsCollectGrpc(userId)

	// 4. 가져온 데이터를 res에 넣는다
	res = &resource.FindEventResponse{
		EventId:   point.EventID,
		Latitude:  point.Latitude,
		Longitude: point.Longitude,
		IsCollect: userlist,
	}

	return

}
