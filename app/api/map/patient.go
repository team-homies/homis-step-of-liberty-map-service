package patient

import (
	"main/app/api/map/handler"
	"main/constant"

	"github.com/gofiber/fiber/v2"
)

func SetMapApis(route fiber.Router) {
	h := handler.NewMapHandler()
	// 환자 리스트 조회
	route.Get(constant.GetPath().Patient.GetPatient, h.GetPatient)

}
