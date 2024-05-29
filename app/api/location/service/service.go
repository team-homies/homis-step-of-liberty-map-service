package service

import (
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

	// 1. isCollect 참거짓 구분
	countDex, err := locationRepository.FindDexById(userId)
	// 2. 만약 값이 0이 아니라면 에러 반환
	if countDex != 0 {
		return nil, err
		// 3. 만약 값이 0이면 만들어놓은 레포지토리를 사용해 데이터를 가져온다
	} else {
		eventFind, err := locationRepository.FindEventByLocation(req.Latitude, req.Longitude)
		if err != nil {
			return nil, err
		}
		// 4. 가져온 데이터를 res에 넣는다
		res = &resource.FindEventResponse{
			Id:        eventFind.ID,
			Latitude:  eventFind.Latitude,
			Longitude: eventFind.Longitude,
			// IsCollect: ,
		}
	}
	return

}
