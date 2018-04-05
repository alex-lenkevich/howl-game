package game

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type DB struct {
	database *mgo.Database
}

func Connect(url string) (*DB, error)  {
	var db *DB
	session, err := mgo.Dial(url)
	if err != nil {
		return db, err
	}
	info, err := mgo.ParseURL(url)
	if err != nil {
		return db, err
	}
	dbName := info.Database
	database := session.DB(dbName)
	collection := database.C("stats")
	log.Printf("dbName: %s, BD: %+v collection: %+v\n", dbName, db, collection)
	_, err = collection.Find(bson.M{}).Count()
	if err != nil {
		panic(err)
	}
	db = &DB{database}
	return db, err
}

func (db DB) SaveStats(stats Stats) error {
	_, err := db.database.C("playerstats").Upsert(bson.M{"player": stats.Player}, stats)
	return err
}

func (db DB) LoadStats(player int64) (Stats, error) {
	var stats Stats
	err := db.database.C("playerstats").Find(bson.M{"player": player}).One(&stats)
	return stats, err
}
