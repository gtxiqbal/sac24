package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gtxiqbal/sac24/service"
)

type RegionalControllerImpl struct {
	RegionalService service.RegionalService
}

func NewRegionalControllerImpl(regionalService service.RegionalService) RegionalController {
	return &RegionalControllerImpl{RegionalService: regionalService}
}

func (controller *RegionalControllerImpl) SetRoute(app *fiber.App) {
	app.Get("/api/regional", controller.GetAll)
	app.Get("/api/regional/:regionalId", controller.GetById)
}

func (controller *RegionalControllerImpl) GetAll(ctx *fiber.Ctx) error {
	response := controller.RegionalService.FindAll(ctx.Context())
	return ctx.JSON(response)
}

func (controller *RegionalControllerImpl) GetById(ctx *fiber.Ctx) error {
	response := controller.RegionalService.FindById(ctx.Context(), ctx.Params("regionalId"))
	return ctx.JSON(response)
}
