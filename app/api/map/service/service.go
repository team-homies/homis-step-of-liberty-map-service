package service

import "main/app/api/map/resource"

type MapService interface {
	GetPatient(req *resource.GetPatientRequest) (res *resource.GetPatientResponse, err error)
}

func NewMapService() MapService {
	return &mapService{
		MapService: &mapService{},
	}
}

type mapService struct {
	MapService
}

func (s *mapService) GetPatient(req *resource.GetPatientRequest) (res *resource.GetPatientResponse, err error) {
	res = new(resource.GetPatientResponse)
	res = &resource.GetPatientResponse{
		Patient: resource.MapResource{
			Id:   1,
			Name: "John Doe",
			Age:  25,
		},
	}
	return
}
