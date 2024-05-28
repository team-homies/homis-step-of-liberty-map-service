package handler

import (
	"main/app/api/map/resource"
	"main/app/api/map/service"
	"main/common/fiberkit"

	"github.com/gofiber/fiber/v2"
)

type handler interface {
	GetPatient(c *fiber.Ctx) error
}

type mapHandler struct {
	service service.MapService
}

func NewMapHandler() handler {
	return &mapHandler{
		service: service.NewMapService(),
	}
}

func (h *mapHandler) GetPatient(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}
	req := new(resource.GetPatientRequest)
	ctx.C.QueryParser(req)

	res, err := h.service.GetPatient(req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(res)
}
