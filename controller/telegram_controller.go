package controller

import "github.com/gofiber/fiber/v2"

type TelegramController interface {
	SetRoute(app *fiber.App)
	DoPostByTelegram(ctx *fiber.Ctx) error
	DoCheckSpec(ctx *fiber.Ctx) error
}
