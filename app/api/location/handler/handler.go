package handler

import (
	"main/app/api/location/resource"
	"main/app/api/location/service"
	"main/common/fiberkit"
	"main/constant/common"

	"github.com/gofiber/fiber/v2"
)

type handler interface {
	FindEvent(c *fiber.Ctx) error
}

type locationHandler struct {
	service service.LocationService
}

func NewLocationHandler() handler {
	return &locationHandler{
		service: service.NewLocationService(),
	}
}

// 사건 조회 query : latitude longitude , header : userId
func (h *locationHandler) FindEvent(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}

	req := new(resource.FindEventRequest)
	ctx.C.BodyParser(req)

	// 1. 쿼리파라미터에서 위도경도 받기
	err := ctx.C.QueryParser(req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}

	// 2. middleware에서 user_id 받기
	req.UserId = uint(ctx.GetLocalsUint(common.LOCALS_USER_ID))

	// 3. 서비스 함수 실행
	res, err := h.service.FindEvent(req.UserId, req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(res)
}
