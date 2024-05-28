package api

import (
	patient "main/app/api/map"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	mapGroup := app.Group("/map")
	patient.SetMapApis(mapGroup)
}
