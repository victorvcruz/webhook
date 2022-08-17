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

type Slack struct {
	Url      string
	Channel  string
	BotToken string
}

type SlackMessage struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
	Thread  string `json:"thread_ts,omitempty"`
}

func (s *Slack) SendMessage(body request.DataToChat, typePullRequest response.PullRequestType, threadID string) (string, error) {

	data := SlackMessage{s.getMessage(typePullRequest, body), s.Channel, threadID}
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", s.Url, bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", "Bearer "+s.BotToken)

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

	return bodyData["ts"].(string), nil
}

func (s *Slack) getMessage(action response.PullRequestType, data request.DataToChat) string {
	switch action {
	case response.CLOSED_PULL_REQUEST:
		return fmt.Sprintf("<!here> *%s* CLOSED A PULL REQUEST IN THE <%s|%s> REPOSITORY!\n\n%s",
			data.User,
			data.RepositoryUrl,
			data.RepositoryName,
			data.PullRequestUrl)
	case response.OPEN_PULL_REQUEST:
		return fmt.Sprintf("<!here> *%s* CREATED A PULL REQUEST IN THE <%s|%s> REPOSITORY!\n\n%s",
			data.User,
			data.RepositoryUrl,
			data.RepositoryName,
			data.PullRequestUrl)
	case response.MERGED_PULL_REQUEST:
		return fmt.Sprintf("<!here> *%s* MERGED DE A PULL REQUEST IN THE <%s|%s> REPOSITORY!\n\n%s",
			data.User,
			data.RepositoryUrl,
			data.RepositoryName,
			data.PullRequestUrl)
	case response.APPROVED_PULL_REQUEST:
		return fmt.Sprintf("<!here> *%s* APPROVED A PULL REQUEST IN THE <%s|%s> REPOSITORY!\n\n%s",
			data.User,
			data.RepositoryUrl,
			data.RepositoryName,
			data.PullRequestUrl)
	case response.REOPEN_PULL_REQUEST:
		return fmt.Sprintf("<!here> *%s* REOPENED A PULL REQUEST IN THE <%s|%s> REPOSITORY!\n\n%s",
			data.User,
			data.RepositoryUrl,
			data.RepositoryName,
			data.PullRequestUrl)
	}
	return ""
}
