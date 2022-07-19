package controller

import "github.com/gofiber/fiber/v2"

type WebController interface {
	SetRoute(app *fiber.App)
	GetIndex(ctx *fiber.Ctx) error
}
