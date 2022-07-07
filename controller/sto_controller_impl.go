package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gtxiqbal/sac24/service"
)

type StoControllerImpl struct {
	service.StoService
}

func NewStoControllerImpl(stoService service.StoService) StoController {
	return &StoControllerImpl{StoService: stoService}
}

func (controller *StoControllerImpl) SetRoute(app *fiber.App) {
	app.Get("/api/sto", controller.GetAll)
	app.Get("/api/sto/:id", controller.GetById)
	app.Get("/api/sto/witel/:witelId", controller.GetByWitelId)
	app.Get("/api/sto/regional/:regionalId", controller.GetByRegionalId)
}

func (controller *StoControllerImpl) GetAll(ctx *fiber.Ctx) error {
	response := controller.FindAll(ctx.Context())
	return ctx.JSON(response)
}

func (controller *StoControllerImpl) GetById(ctx *fiber.Ctx) error {
	response := controller.FindById(ctx.Context(), ctx.Params("id"))
	return ctx.JSON(response)
}

func (controller *StoControllerImpl) GetByWitelId(ctx *fiber.Ctx) error {
	response := controller.FindByWitelId(ctx.Context(), ctx.Params("witelId"))
	return ctx.JSON(response)
}

func (controller *StoControllerImpl) GetByRegionalId(ctx *fiber.Ctx) error {
	response := controller.FindByRegionalId(ctx.Context(), ctx.Params("regionalId"))
	return ctx.JSON(response)
}
