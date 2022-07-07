package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gtxiqbal/sac24/helper"
	"github.com/gtxiqbal/sac24/model/web"
	"github.com/gtxiqbal/sac24/model/web/telegram/request"
	"github.com/gtxiqbal/sac24/service"
	"runtime"
)

type TelegramControllerImpl struct {
	service.TelegramService
}

func NewTelegramControllerImpl(telegramService service.TelegramService) TelegramController {
	return &TelegramControllerImpl{TelegramService: telegramService}
}

func (controller *TelegramControllerImpl) SetRoute(app *fiber.App) {
	app.Post("/api/telegram", controller.DoPostByTelegram)
	app.Get("/api/telegram", controller.DoCheckSpec)
}

func (controller *TelegramControllerImpl) DoPostByTelegram(ctx *fiber.Ctx) error {
	req := new(request.TelegramRequest)
	err := ctx.BodyParser(req)
	helper.PanicIfError(err)

	response := controller.TelegramService.DoPostByTelegram(ctx.Context(), req)
	return ctx.JSON(response)
}

func (controller *TelegramControllerImpl) DoCheckSpec(ctx *fiber.Ctx) error {
	dataMap := make(map[string]any, 0)
	dataMap["total_cpu"] = runtime.NumCPU()
	dataMap["total_thread"] = runtime.GOMAXPROCS(-1)
	dataMap["total_goroutine"] = runtime.NumGoroutine()

	response := web.Response[map[string]any]{
		Code:    "00",
		Status:  "SUCCESS",
		Message: "Check Spec",
		Data:    dataMap,
	}

	return ctx.JSON(response)
}
