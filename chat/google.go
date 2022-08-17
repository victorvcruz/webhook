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

type Google struct {
	Url string
}

type GoogleMessage struct {
	Text   string            `json:"text"`
	Thread map[string]string `json:"thread,omitempty"`
}

func (g *Google) SendMessage(body request.DataToChat, typePullRequest response.PullRequestType, threadID string) (string, error) {

	threadMap := make(map[string]string)
	threadMap["name"] = threadID

	data := GoogleMessage{g.getMessage(typePullRequest, body), threadMap}
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", g.Url, bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

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

	return bodyData["thread"].(map[string]interface{})["name"].(string), nil
}

func (g *Google) getMessage(action response.PullRequestType, data request.DataToChat) string {
	switch action {
	case response.CLOSED_PULL_REQUEST:
		return fmt.Sprintf("<users/all> *%s* CLOSED A PULL REQUEST IN THE <%s|%s> REPOSITORY!\n\n%s",
			data.User,
			data.RepositoryUrl,
			data.RepositoryName,
			data.PullRequestUrl)
	case response.OPEN_PULL_REQUEST:
		return fmt.Sprintf("<users/all> *%s* CREATED A PULL REQUEST IN THE <%s|%s> REPOSITORY!\n\n%s",
			data.User,
			data.RepositoryUrl,
			data.RepositoryName,
			data.PullRequestUrl)
	case response.MERGED_PULL_REQUEST:
		return fmt.Sprintf("<users/all> *%s* MERGED A PULL REQUEST IN THE <%s|%s> REPOSITORY!\n\n%s",
			data.User,
			data.RepositoryUrl,
			data.RepositoryName,
			data.PullRequestUrl)
	case response.APPROVED_PULL_REQUEST:
		return fmt.Sprintf("<users/all> *%s* APPROVED A PULL REQUEST IN THE <%s|%s> REPOSITORY!\n\n%s",
			data.User,
			data.RepositoryUrl,
			data.RepositoryName,
			data.PullRequestUrl)
	case response.REOPEN_PULL_REQUEST:
		return fmt.Sprintf("<users/all> *%s* REOPENED A PULL REQUEST IN THE <%s|%s> REPOSITORY!\n\n%s",
			data.User,
			data.RepositoryUrl,
			data.RepositoryName,
			data.PullRequestUrl)
	}
	return ""
}
