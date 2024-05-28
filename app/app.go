package app

import (
	"main/app/api"

	g "main/app/grpc"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"google.golang.org/grpc"
)

func InitFiber() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	api.InitRoutes(app)

	return app
}

func InitGrpc() *grpc.Server {
	grpcServer := grpc.NewServer()
	g.RegisterServices(grpcServer)
	return grpcServer
}
