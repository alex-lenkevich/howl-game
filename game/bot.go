package game

import (
	"net/http"
	"fmt"
	"encoding/json"
	"bytes"
	"log"
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

var token = GetConfiguration().String("botToken")

func InitWebhook(url string) {
	reqBody, err := json.Marshal(WebhookRequest{url})
	if err != nil {
		log.Fatal(err)
	}
	var resp *http.Response
	var body []byte
	resp, err = sendRequest("setWebhook", reqBody)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	log.Printf("Response code " + resp.Status + " and body: " + string(body))

	resp, err = sendRequest("getWebhookInfo", []byte("{}"))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	log.Printf("Response: " + string(body))
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
	resp, err := sendRequest("sendMessage", body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
}

func sendRequest(method string, body []byte) (resp *http.Response, err error)  {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, method)
	log.Println("POST " + url + " with body " + string(body))
	return http.Post(url, "application/json", bytes.NewReader(body));
}

