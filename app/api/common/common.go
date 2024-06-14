package common

import (
	"context"
	"main/app/grpc/proto/iscollect"
	"main/config"
	"math"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// isCollect grpc
func GetIsCollectGrpc(userId uint) (result bool, err error) {
	// 0. grpc 연동
	var address string
	if viper.GetString(config.GRPC_HISTORY_HOST) == "localhost" {
		address = viper.GetString(config.GRPC_HISTORY_PORT)
	} else {
		address = viper.GetString(config.GRPC_HISTORY_HOST) + viper.GetString(config.GRPC_HISTORY_PORT)
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return
	}
	defer conn.Close()
	isCollectClient := iscollect.NewIsCollectServiceClient(conn)

	// 1. userId를 이용해서 user의 수집률을 구하고 변수에 담는다
	isCollect, err := isCollectClient.GetIsCollect(context.Background(), &iscollect.IsCollectRequest{
		UserId: uint64(userId),
	})
	if err != nil {
		return
	}

	result = isCollect.IsCollect

	return
}

// Radius of earth in KM
const RadiusKm = 6371.01

type Point struct {
	Latitude  float64
	Longitude float64
}

// 위도경도 계산
func CalculateLatLonRange(lat, lon, radiusKm float64) (Point, Point) {

	// 위도 1도 당 km
	latitudeKm := 1.0 / (radiusKm * (math.Pi / 180.0))
	// 경도 1도 당 km
	longitudeKm := 1.0 / ((radiusKm * (math.Pi / 180.0)) * math.Cos(lat*math.Pi/180.0))

	// 북위도 (최대위도)
	northLat := lat + latitudeKm
	// 남위도 (최소위도)
	southLat := lat - latitudeKm
	// 동경도 (최대경도)
	eastLon := lon + longitudeKm
	// 서경도 (최소경도)
	westLon := lon - longitudeKm

	return Point{northLat, eastLon}, Point{southLat, westLon}

}
