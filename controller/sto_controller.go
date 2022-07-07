package controller

import "github.com/gofiber/fiber/v2"

type StoController interface {
	SetRoute(app *fiber.App)
	GetByWitelId(ctx *fiber.Ctx) error
	GetByRegionalId(ctx *fiber.Ctx) error
}
