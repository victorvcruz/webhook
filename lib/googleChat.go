package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Message struct {
	Text   string            `json:"text"`
	Thread map[string]string `json:"thread,omitempty"`
}

func GoogleChat(body string, thread ...map[string]string) {
	webhookURL := "https://chat.googleapis.com/v1/spaces/AAAA_0GFvXc/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=wpBmnGPTjIdWmo-DhXkU4_iNwMvOxn4p-dW0CzfsaMg%3D"

	var data Message
	if len(thread) > 0 {
		data = Message{body, thread[0]}
	} else {
		data = Message{body, nil}
	}

	b, _ := json.Marshal(data)

	client := &http.Client{}

	req, _ := http.NewRequest("POST", webhookURL, bytes.NewBuffer(b))

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	resp, _ := client.Do(req)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	bodyRes := string(b)
	fmt.Println(bodyRes)
}
