package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strings"
	"time"
	"webhooks-chat/chat"
	"webhooks-chat/controllers/request"
	"webhooks-chat/controllers/response"
	"webhooks-chat/database"
)

type Controller struct {
	MongoDB *mongo.Collection
	Chat    chat.ChatClient
}

var isTimeout bool

func (a *Controller) Post(c *fiber.Ctx) error {
	webhookData := make(map[string]interface{})
	err := json.Unmarshal(c.Body(), &webhookData)

	id := webhookData["pull_request"].(map[string]interface{})["id"].(float64)

	threadID := ""
	dataByDb, err := database.Find(id, a.MongoDB)
	if dataByDb != nil {
		threadID = dataByDb["thread"].(string)
	}

	doneC := make(chan bool, 1)
	timeOut := make(chan bool, 1)
	go func() {
		isTimeout = false
		time.Sleep(20 * time.Second)
		isTimeout = true
		timeOut <- true
	}()
	go func() {
		typeMessage := response.GetType(webhookData)
		threadID, err = a.Chat.SendMessage(a.parseWebhookToDataChat(webhookData), typeMessage, threadID)
		if err != nil {
			c.Status(http.StatusInternalServerError).JSON(`{"message": "Internal Server Error"}`)
			doneC <- true
		}

		if dataByDb == nil {
			err = database.Insert(id, threadID, a.MongoDB)
			if err != nil {
				c.Status(http.StatusInternalServerError).JSON(`{"message": "Internal Server Error"}`)
				doneC <- true
			}
		}

		c.Status(http.StatusOK).JSON(`{"message": "OK"}`)
		doneC <- true
	}()
	select {
	case <-doneC:
		return nil
	case <-timeOut:
		c.Status(http.StatusGatewayTimeout).JSON(`{"message": "Gateway Timeout"}`)
		return nil
	}
}

func (a *Controller) parseWebhookToDataChat(data map[string]interface{}) request.DataToChat {
	return request.DataToChat{
		User:           data["sender"].(map[string]interface{})["login"].(string),
		RepositoryUrl:  data["repository"].(map[string]interface{})["clone_url"].(string),
		RepositoryName: data["repository"].(map[string]interface{})["full_name"].(string),
		PullRequestUrl: strings.NewReplacer("api.", "", "repos/", "", "/pulls", "/pull").Replace(data["pull_request"].(map[string]interface{})["url"].(string)),
	}
}
