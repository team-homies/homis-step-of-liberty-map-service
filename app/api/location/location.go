package patient

import (
	"main/app/api/location/handler"
	"main/constant"
	"main/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetlocationApis(route fiber.Router) {
	h := handler.NewLocationHandler()
	// 사건 조회
	route.Get(constant.GetPath().Location.FindEvent, middleware.JWTMiddleware, h.FindEvent)

}
