package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gtxiqbal/sac24/config"
	"github.com/gtxiqbal/sac24/controller"
	"github.com/gtxiqbal/sac24/exception"
	"github.com/gtxiqbal/sac24/helper"
	"github.com/gtxiqbal/sac24/repository"
	"github.com/gtxiqbal/sac24/service"
	"github.com/gtxiqbal/sac24/service/tl1"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.CustomHandling,
	})
	app.Use(recover2.New())
	app.Use(logger.New())
	app.Use(requestid.New())

	err := godotenv.Load()
	helper.PanicIfError(err)

	dbDriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	helper.PanicIfError(err)

	dbMaxIdleConn, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))
	helper.PanicIfError(err)

	dbMaxOpenConn, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN"))
	helper.PanicIfError(err)

	dbConnMaxIdleTime, err := strconv.Atoi(os.Getenv("DB_CONN_MAX_IDLE_TIME"))
	helper.PanicIfError(err)

	dbConnMaxLifeTime, err := strconv.Atoi(os.Getenv("DB_CONN_MAX_LIFE_TIME"))
	helper.PanicIfError(err)

	db := config.NewPG(dbDriver, dbHost, dbName, dbUsername, dbPassword,
		dbPort, dbMaxIdleConn, dbMaxOpenConn, dbConnMaxIdleTime, dbConnMaxLifeTime)

	regionalRepository := repository.NewRegionalRepositoryImpl()
	regionalService := service.NewRegionalServiceImpl(db, regionalRepository)
	regionalController := controller.NewRegionalControllerImpl(regionalService)
	regionalController.SetRoute(app)

	witelRepository := repository.NewWitelRepositoryImpl()
	witelService := service.NewWitelServiceImpl(db, witelRepository)
	witelController := controller.NewWitelControllerImpl(witelService)
	witelController.SetRoute(app)

	stoRepository := repository.NewStoRepositoryImpl()
	stoService := service.NewStoServiceImpl(db, stoRepository)
	stoController := controller.NewStoControllerImpl(stoService)
	stoController.SetRoute(app)

	nmsRepository := repository.NewNmsRepositoryImpl()
	nmsService := service.NewNmsServiceImpl(db, nmsRepository)
	nmsController := controller.NewNmsControllerImpl(nmsService)
	nmsController.SetRoute(app)

	gponRepository := repository.NewGponRepositoryImpl()
	gponService := service.NewGponServiceImpl(db, gponRepository)
	gponController := controller.NewGponControllerImpl(gponService)
	gponController.SetRoute(app)

	userRepository := repository.NewUserRepositoryImpl()

	commandCheck := tl1.NewCommandCheck(2, 21, 2)
	commandCheckSSH := tl1.NewCommandCheckSSHImpl()

	botTokenTelegram := os.Getenv("BOT_TOKEN_TELEGRAM")
	configViaTelegramService := service.NewConfigViaTelegramServiceImpl(db, gponRepository, userRepository, commandCheckSSH, commandCheck, botTokenTelegram)

	telegramService := service.NewTelegramServiceImpl(configViaTelegramService)
	telegramController := controller.NewTelegramControllerImpl(telegramService)
	telegramController.SetRoute(app)

	serverPort := os.Getenv("SERVER_PORT")
	log.Fatal(app.Listen(":" + serverPort))
}
