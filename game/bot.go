package game

import (
	"net/http"
	"fmt"
	"encoding/json"
	"bytes"
	"io"
	"log"
	"strings"
	"io/ioutil"
)

type WebhookRequest struct {
	Url string `json:"url"`
}

type UpdateRequest struct {
	Offset int64 `json:"offset"`
}

type Update struct {
	Id int64
	Text string
	From int64
}

type OutMessage struct {
	Subject int64 `json:"chat_id"`
	Text string `json:"text"`
}

func InitWebhook(url string) {
	reqBody, err := json.Marshal(WebhookRequest{url})
	if err != nil {
		log.Fatal(err)
	}
	sendRequest("setWebhook", bytes.NewReader(reqBody))
	resp, err := sendRequest("getWebhookInfo", strings.NewReader("{}"))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	log.Printf(string(body))
}

func ProcessUpdates(body []byte) Update {

	var updateJson map[string]interface{}

	if err := json.Unmarshal(body, &updateJson); err != nil {
		log.Fatal(err)
	}

	messageSrc := updateJson["message"].(map[string]interface{})
	id := int64(updateJson["update_id"].(float64))
	from := int64(messageSrc["from"].(map[string]interface{})["id"].(float64))
	text := messageSrc["text"].(string)

	return Update{id, text, from}
}

func SendMessage(to int64, msg string) {
	body, err := json.Marshal(OutMessage{to, msg})
	if err != nil {
		log.Fatal(err)
	}
	resp, err := sendRequest("sendMessage", bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
}

func sendRequest(method string, body io.Reader) (resp *http.Response, err error)  {
	url := fmt.Sprintf("https://api.telegram.org/bot589672797:AAFWeN_wUc7v206dIdFceK_6VjmB9C68O6Q/%s", method)
	return http.Post(url, "application/json", body);
}

