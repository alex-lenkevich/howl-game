package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
	"github.com/alex-lenkevich/howl-game/bot"
)

func main() {

	//var offset int64

	bot.InitWebhook("https://howlbot.herokuapp.com/updates")

	port := os.Getenv("PORT")

	log.Println("Started!!!")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":" + port, nil))

	//for i := 0; i < 100000; i++ {
	//	log.Println("Getting updates...")
	//	updates := bot.GetUpdates(offset)
	//	log.Printf("Got updates %+v", updates)
	//	for _, v := range updates {
	//		offset = v.Id + 1
	//		if v.Text == "ping" {
	//			bot.SendMessage(bot.OutMessage{v.From, "pong"})
	//		}
	//	}
	//	time.Sleep(1000000)
	//}

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello world</h1>")
}


