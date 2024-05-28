package patient

import (
	"main/app/api/location/handler"
	"main/constant"

	"github.com/gofiber/fiber/v2"
)

func SetlocationApis(route fiber.Router) {
	h := handler.NewLocationHandler()
	// 환자 리스트 조회
	route.Get(constant.GetPath().Patient.GetPatient, h.GetPatient)

}
