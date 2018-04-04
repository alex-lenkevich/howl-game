package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {

	//var offset int64

	log.Println("Started!!!")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

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
	fmt.Fprint(w, "<h1>Hello world</h1>`")
}


