package controller

import "github.com/gofiber/fiber/v2"

type Controller interface {
	SetRoute(app *fiber.App)
	GetAll(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
}
