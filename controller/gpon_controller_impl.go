package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gtxiqbal/sac24/service"
)

type GponControllerImpl struct {
	service.GponService
}

func NewGponControllerImpl(gponService service.GponService) GponController {
	return &GponControllerImpl{GponService: gponService}
}

func (controller *GponControllerImpl) SetRoute(app *fiber.App) {
	app.Get("/api/gpon", controller.GetAll)
	app.Get("/api/gpon/:id", controller.GetById)
	app.Get("/api/gpon/hostname/:hostname", controller.GetByHostname)
	app.Get("/api/gpon/ip_address/:ipAddress", controller.GetByIpAddress)
	app.Get("/api/gpon/nms/:nmsId", controller.GetByNmsId)
	app.Get("/api/gpon/nms/ip_server/:nmsIpServer", controller.GetByNmsIpServer)
	app.Get("/api/gpon/sto/:stoId", controller.GetByStoId)
	app.Get("/api/gpon/witel/:witelId", controller.GetByWitelId)
	app.Get("/api/gpon/regional/:regionalId", controller.GetByRegionalId)
}

func (controller *GponControllerImpl) GetAll(ctx *fiber.Ctx) error {
	response := controller.FindAll(ctx.Context())
	return ctx.JSON(response)
}

func (controller *GponControllerImpl) GetById(ctx *fiber.Ctx) error {
	response := controller.FindById(ctx.Context(), ctx.Params("id"))
	return ctx.JSON(response)
}

func (controller *GponControllerImpl) GetByHostname(ctx *fiber.Ctx) error {
	response := controller.FindByHostname(ctx.Context(), ctx.Params("hostname"))
	return ctx.JSON(response)
}

func (controller *GponControllerImpl) GetByIpAddress(ctx *fiber.Ctx) error {
	response := controller.FindByIpAddress(ctx.Context(), ctx.Params("ipAddress"))
	return ctx.JSON(response)
}

func (controller *GponControllerImpl) GetByNmsId(ctx *fiber.Ctx) error {
	response := controller.FindByNmsId(ctx.Context(), ctx.Params("nmsId"))
	return ctx.JSON(response)
}

func (controller *GponControllerImpl) GetByNmsIpServer(ctx *fiber.Ctx) error {
	response := controller.FindByNmsIpServer(ctx.Context(), ctx.Params("nmsIpServer"))
	return ctx.JSON(response)
}

func (controller *GponControllerImpl) GetByStoId(ctx *fiber.Ctx) error {
	response := controller.FindByStoId(ctx.Context(), ctx.Params("stoId"))
	return ctx.JSON(response)
}

func (controller *GponControllerImpl) GetByWitelId(ctx *fiber.Ctx) error {
	response := controller.FindByWitelId(ctx.Context(), ctx.Params("witelId"))
	return ctx.JSON(response)
}

func (controller *GponControllerImpl) GetByRegionalId(ctx *fiber.Ctx) error {
	response := controller.FindByRegionalId(ctx.Context(), ctx.Params("regionalId"))
	return ctx.JSON(response)
}
