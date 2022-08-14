package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strings"
	"webhooks-chat/chat"
	"webhooks-chat/controllers/request"
	"webhooks-chat/controllers/response"
	"webhooks-chat/database"
)

type Controller struct {
	MongoDB *mongo.Collection
	Chat    chat.ChatClient
}

func (a *Controller) Post(c *fiber.Ctx) error {
	webhookData := make(map[string]interface{})
	err := json.Unmarshal(c.Body(), &webhookData)

	id := webhookData["pull_request"].(map[string]interface{})["id"].(float64)

	threadID := ""
	dataByDb, err := database.Find(id, a.MongoDB)
	if dataByDb != nil {
		threadID = dataByDb["thread"].(string)
	}

	typeMessage := response.GetType(webhookData)
	res, err := a.Chat.SendMessage(a.parseWebhookToDataChat(webhookData), typeMessage, threadID)

	if dataByDb == nil {
		err = database.Insert(id, res["thread"].(map[string]interface{})["name"].(string), a.MongoDB)
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(`{"message": "Server Exception"}`)
	}
	return c.Status(http.StatusOK).JSON(`{"message": "OK"}`)
}

func (a *Controller) parseWebhookToDataChat(data map[string]interface{}) request.DataToChat {
	return request.DataToChat{
		User:           data["sender"].(map[string]interface{})["login"].(string),
		RepositoryUrl:  data["repository"].(map[string]interface{})["clone_url"].(string),
		RepositoryName: data["repository"].(map[string]interface{})["full_name"].(string),
		PullRequestUrl: strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(data["pull_request"].(map[string]interface{})["url"].(string), "api.", ""), "repos/", ""), "/pulls", "/pull"),
	}
}
