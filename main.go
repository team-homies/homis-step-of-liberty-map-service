package main

import (
	"log"
	"main/app"
	"main/config"
	"net"

	"github.com/spf13/viper"
)

func init() {
	config.InitConfig()
}

func main() {
	// fiber
	fiber := app.InitFiber()

	port := viper.GetString(viper.GetString(config.APP_PORT))
	go fiber.Listen(port)

	// grpc
	listen, err := net.Listen("tcp", viper.GetString(config.APP_GRPC_PORT))
	if err != nil {
		log.Fatalf("grpc 포트 에러: %v", err)
	}

	// gRPC 초기화
	grpcServer := app.InitGrpc()

	// 서버 구동
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("grpc 서버 구동 실패: %v", err)
	}
}
