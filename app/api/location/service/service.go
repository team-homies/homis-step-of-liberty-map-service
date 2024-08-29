package service

import (
	"main/app/api/common"
	"main/app/api/location/resource"
	com "main/constant/common"
	"main/database/repository"
)

type LocationService interface {
	FindEvent(userId uint, req *resource.FindEventRequest) (res *resource.FindEventResponse, err error)
	FindLocationList(userId uint, req *resource.FindLocationListRequest) (res []resource.FindLocationListResponse, err error)
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
	MaxLatLon, MinLatLon := common.CalculateLatLonRange10(req.Latitude, req.Longitude, com.RadiusKm)

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

// 위치 리스트 조회
func (s *locationService) FindLocationList(userId uint, req *resource.FindLocationListRequest) (res []resource.FindLocationListResponse, err error) {
	// 1. 사용자의 위도경도를 담아 최대최소위도경도 계산 함수 호출
	MaxLatLon, MinLatLon := common.CalculateLatLonRange1000(req.Latitude, req.Longitude, com.RadiusKm)

	// 2. 최대최소위도경도를 담아 범위 내의 정보 repository함수 호출
	infoList, err := repository.NewRepository().FindLocationListByPoint(MaxLatLon, MinLatLon)
	if err != nil {
		return
	}

	// 3. 반환 리소스에 정보 담기
	for _, info := range infoList {
		// 3-1. 두 포인트 사이의 거리 계산
		distance := common.CalculateDistance(req.Latitude, req.Longitude, info.Latitude, info.Longitude, com.RadiusKm)

		// 3-2. 사용자의 정보를 담아 수집 여부 조회 grpc함수 호출
		var isCollect bool
		isCollect, err = common.GetIsCollectGrpc(info.EventId, userId)
		if err != nil {
			return
		}

		res = append(res, resource.FindLocationListResponse{
			EventId:     info.EventId,
			Latitude:    info.Latitude,
			Longitude:   info.Longitude,
			Distance:    uint(distance),
			Address:     info.Address,
			AddressRoad: info.AddressRoad,
			IsCollect:   isCollect,
		})
	}

	return
}
