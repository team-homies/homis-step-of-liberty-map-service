package handler

import (
	"main/app/api/location/resource"
	"main/app/api/location/service"
	"main/common/fiberkit"

	"github.com/gofiber/fiber/v2"
)

type handler interface {
	GetPatient(c *fiber.Ctx) error
}

type locationHandler struct {
	service service.LocationService
}

func NewLocationHandler() handler {
	return &locationHandler{
		service: service.NewLocationService(),
	}
}

func (h *locationHandler) GetPatient(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}
	req := new(resource.GetPatientRequest)
	ctx.C.QueryParser(req)

	res, err := h.service.GetPatient(req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(res)
}
