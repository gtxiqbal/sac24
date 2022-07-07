package controller

import "github.com/gofiber/fiber/v2"

type GponController interface {
	SetRoute(app *fiber.App)
	GetByHostname(ctx *fiber.Ctx) error
	GetByIpAddress(ctx *fiber.Ctx) error
	GetByNmsId(ctx *fiber.Ctx) error
	GetByNmsIpServer(ctx *fiber.Ctx) error
	GetByStoId(ctx *fiber.Ctx) error
	GetByWitelId(ctx *fiber.Ctx) error
	GetByRegionalId(ctx *fiber.Ctx) error
}
