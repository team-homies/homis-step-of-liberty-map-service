package grpc

import (
	"main/app/grpc/service/location"
	"google.golang.org/grpc"
)

func RegisterServices(grpcServer *grpc.Server) {
	location.RegisterLocationService(grpcServer)
}
