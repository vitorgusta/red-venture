package dbconfig

import (
	"os"

	"gopkg.in/mgo.v2"
)

type DB struct {
	Session *mgo.Session
}

func (db *DB) DoDial() (s *mgo.Session, err error) {
	return mgo.Dial(DBUrl())
}

func (db *DB) Name() string {
	return "redventure"
}

func DBUrl() string {
	dburl := os.Getenv("MONGOHQ_URL")

	if dburl == "" {
		dburl = "localhost:27017"
	}

	return dburl
}
