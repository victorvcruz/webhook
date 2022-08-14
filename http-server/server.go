package http_server

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"webhooks-chat/chat"
	"webhooks-chat/controllers"
)

type API struct {
	MongoDB *mongo.Collection
}

func (a *API) Run(port string) {
	google := &chat.Google{Url: os.Getenv("GOOGLE_CHAT_URL")}

	controller := &controllers.Controller{a.MongoDB, google}

	app := fiber.New()
	app.Post("/payload", controller.Post)

	app.Listen(":" + port)
}
