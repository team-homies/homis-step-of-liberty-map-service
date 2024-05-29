package handler

import (
	"main/app/api/location/resource"
	"main/app/api/location/service"
	"main/common/fiberkit"

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

	// 1. middleware에서 user_id 받기
	err := ctx.C.BodyParser(req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	userId := uint(ctx.GetLocalsInt("user_id"))

	// 2. 쿼리파라미터에서 위도경도 받기
	err = ctx.C.QueryParser(req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}

	// 3. 서비스 함수 실행
	res, err := h.service.FindEvent(userId, req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}

	// 4. 리턴
	return ctx.HttpOK(res)
}
