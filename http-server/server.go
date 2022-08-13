package http_server

import (
	"github.com/gofiber/fiber/v2"
	"webhooks-chat/controllers"
)

type API struct{}

func (a *API) Run() {

	controller := controllers.Controller{}

	app := fiber.New()

	app.Post("/payload", controller.Post)

	app.Listen(":8080")
}
