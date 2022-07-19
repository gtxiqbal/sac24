package controller

import (
	"github.com/gofiber/fiber/v2"
)

type WebControllerImpl struct {
}

func NewWebControllerImpl() WebController {
	return &WebControllerImpl{}
}

func (controller *WebControllerImpl) SetRoute(app *fiber.App) {
	app.Get("/", controller.GetIndex)
}

func (controller *WebControllerImpl) GetIndex(ctx *fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{})
}
