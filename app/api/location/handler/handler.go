package handler

import (
	"main/app/api/location/service"

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

func (h *locationHandler) FindEvent(c *fiber.Ctx) error

// {
// ctx := fiberkit.FiberKit{C: c}
// req := new(resource.FindEventRequest)
// ctx.C.QueryParser(req)

// res, err := h.service.FindEvent(req)
// if err != nil {
// 	return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
// }
// return ctx.HttpOK(res)
// }
