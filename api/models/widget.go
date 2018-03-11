package models

import "gopkg.in/mgo.v2/bson"

type (
	// Widget
	Widget struct {
		Id        bson.ObjectId `json:"id" bson:"_id"`
		Name      string        `json:"name" bson:"name"`
		Price     float32       `json:"price" bson:"price"`
		Color     string        `json:"color" bson:"color"`
		Melts     bool          `json:"melts" bson:"melts"`
		Inventory string        `json:"inventory" bson:"inventory"`
	}
)

type Widgets []Widget
