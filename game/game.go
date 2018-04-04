package game

import (
	"math/rand"
	"math"
	"gopkg.in/mgo.v2"
	"fmt"
	"log"
)

type Game struct {
	Db *DB
}

func NewGame(db *DB) Game {
	return Game{db}
}

func (game Game) ProcessMessage(act Act) (Result, error) {
	var result Result
	_, err := game.Db.LoadStats(act.Player)
	if act.Action == "/newgame" || err == mgo.ErrNotFound {
		stats, err := game.StartGame(act.Player)
		if err != nil {
			return result, err
		}
		message := fmt.Sprintf(
`You started a new game
Your Mastery: %d
Your Endurance: %d
Your Luck: %d`, stats.Mastery, stats.Endurance, stats.Luck)
		log.Println(message)
		result = Result{message}
	}
	return result, err
}

func (game Game) StartGame(player int64) (Stats, error) {
	stats := Stats{player, NewMastery(), NewEndurance(), NewLuck()}
	err := game.Db.SaveStats(stats)
	return stats, err
}

func NewMastery() int {
	halfDice := float64(RollTheDice() / 2)
	return int(math.Ceil(halfDice) + 7)
}

func NewEndurance() int {
	return int(RollTheDice() + RollTheDice() + 10)
}

func NewLuck() int {
	return int(rand.Intn(6) + 6)
}

func RollTheDice() int {
	return rand.Intn(6) + 1
}
