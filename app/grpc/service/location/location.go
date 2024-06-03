package location

import (
	"context"
	"main/app/grpc/proto/location"
	"main/database/entity"
	"main/database/repository"
	"strconv"

	"google.golang.org/grpc"
)

type server struct {
	location.LocationServiceServer
}

func RegisterLocationService(grpcServer *grpc.Server) {
	location.RegisterLocationServiceServer(grpcServer, &server{})
}

func (s *server) SetHistories(ctx context.Context, in *location.SetHistoryRequest) (*location.HistoryResponse, error) {
	result := &location.HistoryResponse{
		IsSuccess: false,
	}
	req := []entity.Map{}
	for _, location := range in.Histories {
		lat, err := strconv.ParseFloat(location.Latitude, 64)
		if err != nil {
			return result, err
		}
		lon, err := strconv.ParseFloat(location.Longitude, 64)
		if err != nil {
			return result, err
		}
		req = append(req, entity.Map{
			EventId:     uint(location.EventId),
			Name:        location.Name,
			Address:     location.Address,
			AddressRoad: location.AddressRoad,
			Latitude:    lat,
			Longitude:   lon,
		})
	}
	err := repository.NewRepository().SaveHistories(req)
	if err != nil {
		result.IsSuccess = false
		return result, err
	}
	result.IsSuccess = true

	return result, nil
}
