package game

import (
	"math/rand"
	"math"
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
	var result Result;
	var message string;
	_, err := game.Db.LoadStats(act.Player);
	switch act.Command.Action {
	case "start":
		stats, err := game.StartGame(act.Player)
		if err != nil {
			return result, err
		}
		message = fmt.Sprintf(
			`You started a new game
			Your Mastery: %d
			Your Endurance: %d
			Your Luck: %d`, stats.Mastery, stats.Endurance, stats.Luck);
		result = Result{message};
	case "info":
		stats, err := game.Db.LoadStats(act.Player);
		if err != nil {
			return result, err
		}
		message = fmt.Sprintf(
			`Your Mastery: %d
			Your Endurance: %d
			Your Luck: %d`, stats.Mastery, stats.Endurance, stats.Luck);
		result = Result{message};
	case "next":
		stats, err := game.Db.LoadStats(act.Player);
		if err != nil {
			return result, err
		}
		if stats.Location == -1 {
			message = `Your journey has been started!`;
			stats.Location = 0;
			err = game.Db.SaveStats(stats);
		} else {
			message = `Your journey has been already started!!!`;
		}
		result = Result{message};
	}
	return result, err
}

func (game Game) StartGame(player int64) (Stats, error) {
	stats := Stats{player, NewMastery(), NewEndurance(), NewLuck(), -1}
	err := game.Db.SaveStats(stats)
	log.Println(fmt.Sprintf(`New game for player %d`, player));
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
