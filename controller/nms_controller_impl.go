package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gtxiqbal/sac24/service"
)

type NmsControllerImpl struct {
	service.NmsService
}

func NewNmsControllerImpl(nmsService service.NmsService) NmsController {
	return &NmsControllerImpl{NmsService: nmsService}
}

func (controller *NmsControllerImpl) SetRoute(app *fiber.App) {
	app.Get("/api/nms", controller.GetAll)
	app.Get("/api/nms/:id", controller.GetById)
	app.Get("/api/nms/nama/:nama", controller.GetByNama)
	app.Get("/api/nms/vendor/:vendor", controller.GetByVendor)
	app.Get("/api/nms/ip_server/:ipServer", controller.GetByIpServer)
}

func (controller *NmsControllerImpl) GetAll(ctx *fiber.Ctx) error {
	response := controller.FindAll(ctx.Context())
	return ctx.JSON(response)
}

func (controller *NmsControllerImpl) GetById(ctx *fiber.Ctx) error {
	response := controller.FindById(ctx.Context(), ctx.Params("id"))
	return ctx.JSON(response)
}

func (controller *NmsControllerImpl) GetByNama(ctx *fiber.Ctx) error {
	response := controller.FindByNama(ctx.Context(), ctx.Params("nama"))
	return ctx.JSON(response)
}

func (controller *NmsControllerImpl) GetByVendor(ctx *fiber.Ctx) error {
	response := controller.FindByVendor(ctx.Context(), ctx.Params("vendor"))
	return ctx.JSON(response)
}

func (controller *NmsControllerImpl) GetByIpServer(ctx *fiber.Ctx) error {
	response := controller.FindByIpServer(ctx.Context(), ctx.Params("ipServer"))
	return ctx.JSON(response)
}
