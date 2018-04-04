package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
	"github.com/alex-lenkevich/howl-game/bot"
	"io/ioutil"
)

func main() {

	//var offset int64

	bot.InitWebhook("https://agile-waters-36090.herokuapp.com/updatesHook")

	http.HandleFunc("/", hello)
	http.HandleFunc("/updatesHook", newMessage)
	http.ListenAndServe(":" + os.Getenv("PORT"), nil)

	log.Println("Started!!!")

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello world</h1>")
}

func newMessage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	defer w.Write([]byte{})
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got update %s", string(body))
	update := bot.ProcessUpdates(body)
	if update.Text == "ping" {
		bot.SendMessage(update.From, "pong")
	}
}


