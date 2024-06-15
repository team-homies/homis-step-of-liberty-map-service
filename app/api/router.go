package api

import (
	location "main/app/api/location"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	locationGroup := app.Group("/map")
	location.SetLocationApis(locationGroup)
}
