package controller

import "github.com/gofiber/fiber/v2"

type NmsController interface {
	SetRoute(app *fiber.App)
	GetByNama(ctx *fiber.Ctx) error
	GetByVendor(ctx *fiber.Ctx) error
	GetByIpServer(ctx *fiber.Ctx) error
}
