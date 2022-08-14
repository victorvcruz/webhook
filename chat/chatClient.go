package chat

import (
	"webhooks-chat/controllers/request"
	"webhooks-chat/controllers/response"
)

type ChatClient interface {
	SendMessage(body request.DataToChat, typePullRequest response.PullRequestType, threadID string) (map[string]interface{}, error)
	getMessage(action response.PullRequestType, data request.DataToChat) string
}
