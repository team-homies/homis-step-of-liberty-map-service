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
func GetIsCollectGrpc(userId uint, eventId uint) (result bool, err error) {
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
		UserId:  uint64(userId),
		EventId: uint64(eventId),
	})
	if err != nil {
		return
	}

	result = isCollect.IsCollect

	return
}

type Point struct {
	Latitude  float64
	Longitude float64
}

// 위도경도 계산
func CalculateLatLonRange(lat, lon, earthRadius float64) (Point, Point) {

	// 조회에 필요한 실제 사용자 위치와 가까워지는 근사치 = 10m

	// 위도 1도 당 km : (R * 2pi) / 360 = R * (pi / 180)
	latitude10M := 0.01 / (earthRadius * (math.Pi / 180.0))
	// 경도 1도 당 km : (R * 2pi) / 360 * cos(위도) = R * (pi / 180) * cos(위도)
	longitude10M := 0.01 / ((earthRadius * (math.Pi / 180.0)) * math.Cos(lat*math.Pi/180.0))

	// 북위도 (최대위도)
	northLat := lat + latitude10M
	// 남위도 (최소위도)
	southLat := lat - latitude10M
	// 동경도 (최대경도)
	eastLon := lon + longitude10M
	// 서경도 (최소경도)
	westLon := lon - longitude10M

	return Point{northLat, eastLon}, Point{southLat, westLon}

}
