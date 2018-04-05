package game

import "errors"
import s "strings"

type Result struct {
	Message string
}

type Stats struct {
	Player int64 `bson:"player"`
	Mastery int `bson:"mastery"`
	Endurance int `bson:"endurance"`
	Luck int `bson:"luck"`
	Location int `bson:"location"`
}

type Command struct {
	Action string
	Argument string
}

type Act struct {
	Player int64
	Command Command
}

var NotFound = errors.New("not found")

func ParseCommand(commandString string) (Command) {
	var splitted = s.Split(commandString, " ");
	var action = s.Replace(splitted[0], "/", "", 1);
	var argument string;
	if len(splitted) > 1 {
		argument = splitted[1];
	}
	return Command{action, argument};
}
