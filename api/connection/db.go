package connection

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
)

func GetSession() *mgo.Session {

	url := os.Getenv("MONGOHQ_URL")
	// Connect to our local mongo
	//s, err := mgo.Dial("mongodb://redventure:visA1234@ds163918.mlab.com:63918/redventure")
	s, err := mgo.Dial(url)

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
		fmt.Println("err", err)
	}

	// Deliver session
	return s
}
