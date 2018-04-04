package game

import "errors"

type Result struct {
	Message string
}

type Stats struct {
	Player int64 `bson:"player"`
	Mastery int `bson:"mastery"`
	Endurance int `bson:"endurance"`
	Luck int `bson:"luck"`
}

type Act struct {
	Player int64
	Action string
}

var NotFound = errors.New("not found")