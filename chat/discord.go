package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"webhooks-chat/controllers/request"
	"webhooks-chat/controllers/response"
)

type Discord struct {
	Url string
}

type DiscordMessage struct {
	Content string `json:"content"`
	Thread  string `json:"thread_ts,omitempty"`
}

func (d *Discord) SendMessage(body request.DataToChat, typePullRequest response.PullRequestType, threadID string) (string, error) {

	data := DiscordMessage{d.getMessage(typePullRequest, body), threadID}
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", d.Url, bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	//req.Header.Add("Authorization", "Bot "+d.BotToken)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	bodyData := make(map[string]interface{})
	err = json.Unmarshal(b, &bodyData)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (d *Discord) getMessage(action response.PullRequestType, data request.DataToChat) string {
	switch action {
	case response.CLOSED_PULL_REQUEST:
		return fmt.Sprintf("@everyone O *%s* FECHOU UM PULL REQUEST NO REPOSITÓRIO [%s](%s)!\n\n%s",
			data.User,
			data.RepositoryName,
			data.RepositoryUrl,
			data.PullRequestUrl)
	case response.OPEN_PULL_REQUEST:
		return fmt.Sprintf("@everyone O *%s* SOLICITOU UM PULL REQUEST NO REPOSITÓRIO [%s](%s)!\n\n%s",
			data.User,
			data.RepositoryName,
			data.RepositoryUrl,
			data.PullRequestUrl)
	case response.MERGED_PULL_REQUEST:
		return fmt.Sprintf("@everyone O *%s* FEZ MERGE DE UM PULL REQUEST NO REPOSITÓRIO [%s](%s)!\n\n%s",
			data.User,
			data.RepositoryName,
			data.RepositoryUrl,
			data.PullRequestUrl)
	case response.APPROVED_PULL_REQUEST:
		return fmt.Sprintf("@everyone O *%s* APROVOU UM PULL REQUEST NO REPOSITÓRIO [%s](%s)!\n\n%s",
			data.User,
			data.RepositoryName,
			data.RepositoryUrl,
			data.PullRequestUrl)
	case response.REOPEN_PULL_REQUEST:
		return fmt.Sprintf("@everyone O *%s* REABRIU UM PULL REQUEST NO REPOSITÓRIO [%s](%s)!\n\n%s",
			data.User,
			data.RepositoryName,
			data.RepositoryUrl,
			data.PullRequestUrl)
	}
	return ""
}
