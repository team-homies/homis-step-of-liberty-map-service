package service

import "main/app/api/location/resource"

type LocationService interface {
	GetPatient(req *resource.FindEventRequest) (res *resource.FindEventResponse, err error)
}

func NewLocationService() LocationService {
	return &locationService{
		LocationService: &locationService{},
	}
}

type locationService struct {
	LocationService
}

func (s *locationService) GetPatient(req *resource.FindEventRequest) (res *resource.FindEventResponse, err error) {

	return
}
