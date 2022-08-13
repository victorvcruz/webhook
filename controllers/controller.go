package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
	"webhooks-chat/lib"
)

type Controller struct{}

func (a *Controller) Post(c *fiber.Ctx) error {
	webhookData := make(map[string]interface{})
	err := c.BodyParser(webhookData)
	if err != nil {
		return fiber.NewError(782, err.Error())
	}

	messageStr := fmt.Sprintf("<users/all> O *%s* SOLICITOU UM PULL REQUEST NO REPOSITÓRIO <%s|%s>!\n\n%s",
		webhookData["sender"].(map[string]interface{})["login"],
		webhookData["repository"].(map[string]interface{})["clone_url"],
		webhookData["repository"].(map[string]interface{})["full_name"],
		strings.ReplaceAll(strings.ReplaceAll(webhookData["pull_request"].(map[string]interface{})["url"].(string), "api.", ""), "repos/", ""))

	thread := make(map[string]string)
	thread["name"] = "spaces/AAAA_0GFvXc/threads/HtlAsm2hP78"

	lib.GoogleChat(messageStr, thread)

	return c.Status(200).JSON(`{"message": "funfou"}`)
}
