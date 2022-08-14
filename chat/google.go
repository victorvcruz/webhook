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

type Message struct {
	Text   string            `json:"text"`
	Thread map[string]string `json:"thread,omitempty"`
}

func (g *Google) SendMessage(body request.DataToChat, typePullRequest response.PullRequestType, threadID string) (map[string]interface{}, error) {

	threadMap := make(map[string]string)
	threadMap["name"] = threadID

	data := Message{g.getMessage(typePullRequest, body), threadMap}
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", g.Url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyData := make(map[string]interface{})
	err = json.Unmarshal(b, &bodyData)
	if err != nil {
		return nil, err
	}

	return bodyData, nil
}

func (g *Google) getMessage(action response.PullRequestType, data request.DataToChat) string {
	switch action {
	case response.CLOSED_PULL_REQUEST:
		return fmt.Sprintf("<users/all> O *%s* FECHOU UM PULL REQUEST NO REPOSITÓRIO <%s|%s>!\n\n%s",
			data.User,
			data.RepositoryUrl,
			data.RepositoryName,
			data.PullRequestUrl)
	case response.OPEN_PULL_REQUEST:
		return fmt.Sprintf("<users/all> O *%s* SOLICITOU UM PULL REQUEST NO REPOSITÓRIO <%s|%s>!\n\n%s",
			data.User,
			data.RepositoryUrl,
			data.RepositoryName,
			data.PullRequestUrl)
	case response.MERGED_PULL_REQUEST:
		return fmt.Sprintf("<users/all> O *%s* FEZ MERGE DE UM UM PULL REQUEST NO REPOSITÓRIO <%s|%s>!\n\n%s",
			data.User,
			data.RepositoryUrl,
			data.RepositoryName,
			data.PullRequestUrl)
	case response.APPROVED_PULL_REQUEST:
		return fmt.Sprintf("<users/all> O *%s* APROVOU UM UM PULL REQUEST NO REPOSITÓRIO <%s|%s>!\n\n%s",
			data.User,
			data.RepositoryUrl,
			data.RepositoryName,
			data.PullRequestUrl)
	case response.REOPEN_PULL_REQUEST:
		return fmt.Sprintf("<users/all> O *%s* REABRIU UM UM PULL REQUEST NO REPOSITÓRIO <%s|%s>!\n\n%s",
			data.User,
			data.RepositoryUrl,
			data.RepositoryName,
			data.PullRequestUrl)
	}
	return ""
}
