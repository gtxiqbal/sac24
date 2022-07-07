package exception

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gtxiqbal/sac24/model/web"
)

func CustomHandling(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	e, ok := err.(*fiber.Error)
	if ok {
		code = e.Code
	}

	if err != nil {
		fmt.Println(err)
		return ctx.Status(code).JSON(web.Response[any]{
			Code:    "99",
			Status:  "FAILED",
			Message: "Gagal Server API",
			Data:    err.Error(),
		})
	}

	return nil
}
