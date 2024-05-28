package service

import "main/app/api/location/resource"

type LocationService interface {
	GetPatient(req *resource.GetPatientRequest) (res *resource.GetPatientResponse, err error)
}

func NewLocationService() LocationService {
	return &locationService{
		LocationService: &locationService{},
	}
}

type locationService struct {
	LocationService
}

func (s *locationService) GetPatient(req *resource.GetPatientRequest) (res *resource.GetPatientResponse, err error) {

	return
}
