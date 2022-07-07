package controller

import "github.com/gofiber/fiber/v2"

type WitelController interface {
	SetRoute(app *fiber.App)
	GetByRegionalId(ctx *fiber.Ctx) error
}
