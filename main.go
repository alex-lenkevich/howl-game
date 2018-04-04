package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
	"github.com/alex-lenkevich/howl-game/game"
	"io/ioutil"
)

var gm game.Game

func main() {

	dbUrl, found := os.LookupEnv("MONGODB_URI")
	if !found {
		dbUrl = "mongodb://howl:howl@localhost:27017/howl"
	}
	port, found := os.LookupEnv("PORT")
	if !found {
		port = "8080"
	}
	db, e := game.Connect(dbUrl)
	if e != nil {
		panic(e)
	}

	gm = game.NewGame(db)
	//var offset int64

	game.InitWebhook("https://howlbot.herokuapp.com/updatesHook")
	//game.InitWebhook("https://oleksandrlinkevich.fwd.wf/updatesHook")

	http.HandleFunc("/", healthcheck)
	http.HandleFunc("/updatesHook", newMessage)
	http.ListenAndServe(":" + port, nil)

	log.Println("Started!!!")

}

func healthcheck(w http.ResponseWriter, r *http.Request) {
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
	update := game.ProcessUpdates(body)
	if update.Text == "ping" {
		game.SendMessage(update.From, "pong")
		return
	}
	result, err := gm.ProcessMessage(game.Act{Player: update.From, Action: update.Text})
	if err != nil {
		game.SendMessage(update.From, "ERROR")
		log.Fatal(err)
		return
	}
	game.SendMessage(update.From, result.Message)
}


