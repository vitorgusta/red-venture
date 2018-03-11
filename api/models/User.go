package models

import "gopkg.in/mgo.v2/bson"

type (
	// User
	User struct {
		Id   bson.ObjectId `json:"id" bson:"_id" bson:", omitempty "`
		Name string        `json:"name" bson:"name" bson:"name"`
	}
)

type Users []User
