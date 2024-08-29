package patient

import (
	"main/app/api/location/handler"
	"main/constant"
	"main/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetLocationApis(route fiber.Router) {
	h := handler.NewLocationHandler()
	// 사건 조회
	route.Get(constant.LocationPath().Location.FindEvent, middleware.AuthVerificationMiddleware, h.FindEvent)
	// 위치 리스트 조회
	route.Get(constant.LocationPath().Location.FindLocationList, middleware.AuthVerificationMiddleware, h.FindLocationList)
}
