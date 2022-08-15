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
	controller := &controllers.Controller{}

	switch os.Getenv("PLATFORM_CHAT") {
	case "GOOGLE":
		google := &chat.Google{Url: os.Getenv("GOOGLE_CHAT_URL")}
		controller = &controllers.Controller{a.MongoDB, google}
	case "SLACK":
		slack := &chat.Slack{Url: os.Getenv("SLACK_CHAT_URL"), Channel: os.Getenv("SLACK_CHANNEL_ID"), BotToken: os.Getenv("SLACK_BOT_TOKEN")}
		controller = &controllers.Controller{a.MongoDB, slack}
	}

	app := fiber.New()
	app.Post("/payload", controller.Post)

	app.Listen(":" + port)
}
