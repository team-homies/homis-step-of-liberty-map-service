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

// 사건 조회
func (s *locationService) FindEvent(userId uint, req *resource.FindEventRequest) (res *resource.FindEventResponse, err error) {
	locationRepository := repository.NewRepository()

	// 1. 사용자의 위도경도를 담아 최대최소위도경도 계산 함수 호출
	MaxLatLon, MinLatLon := common.CalculateLatLonRange(req.Latitude, req.Longitude, common.RadiusKm)

	// 2. 최대최소위도경도를 담아 범위 내의 event조회 repository함수 호출
	point, err := locationRepository.FindEventByPoint(MaxLatLon, MinLatLon)
	if err != nil {
		return
	}

	// 3. 사용자의 정보를 담아 수집 여부 조회 grpc함수 호출
	userlist, err := common.GetIsCollectGrpc(point.EventId, userId)
	if err != nil {
		return
	}

	// 4. 가져온 데이터를 res에 담는다
	res = &resource.FindEventResponse{
		EventId:   point.EventId,
		Latitude:  point.Latitude,
		Longitude: point.Longitude,
		IsCollect: userlist,
	}

	return

}
