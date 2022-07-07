package controller

import "github.com/gofiber/fiber/v2"

type RegionalController interface {
	SetRoute(app *fiber.App)
}
