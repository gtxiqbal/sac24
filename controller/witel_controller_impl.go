package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gtxiqbal/sac24/service"
)

type WitelControllerImpl struct {
	service.WitelService
}

func NewWitelControllerImpl(witelService service.WitelService) WitelController {
	return &WitelControllerImpl{WitelService: witelService}
}

func (controller *WitelControllerImpl) SetRoute(app *fiber.App) {
	app.Get("/api/witel", controller.GetAll)
	app.Get("/api/witel/:id", controller.GetById)
	app.Get("/api/witel/regional/:regionalId", controller.GetByRegionalId)
}

func (controller *WitelControllerImpl) GetAll(ctx *fiber.Ctx) error {
	response := controller.FindAll(ctx.Context())
	return ctx.JSON(response)
}

func (controller *WitelControllerImpl) GetById(ctx *fiber.Ctx) error {
	response := controller.FindById(ctx.Context(), ctx.Params("id"))
	return ctx.JSON(response)
}

func (controller *WitelControllerImpl) GetByRegionalId(ctx *fiber.Ctx) error {
	response := controller.FindByRegionalId(ctx.Context(), ctx.Params("regionalId"))
	return ctx.JSON(response)
}
