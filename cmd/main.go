package main

import (
	"fmt"

	"github.com/alex-lenkevich/howl-game/bot"
	"time"
)

func main() {

	fmt.Printf("hello, world\n")
	var offset int64

	for i := 0; i < 100; i++ {
		updates := bot.GetUpdates(offset)
		fmt.Println(offset)
		fmt.Println(updates)
		for _, v := range updates {
			offset = v.Id + 1
			if v.Text == "ping" {
				bot.SendMessage(bot.OutMessage{v.From, "pong"})
			}
		}
		time.Sleep(1000000)
	}

}


