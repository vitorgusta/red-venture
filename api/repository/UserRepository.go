package repository

import (
	"github.com/labstack/gommon/log"

	"github.com/vitorgusta/red-venture/api/dbconfig"
	"github.com/vitorgusta/red-venture/api/models"
	"gopkg.in/mgo.v2/bson"
)

func CreateUser(user models.User) models.User {
	db := dbconfig.DB{}
	user.Id = bson.NewObjectId()

	s, err := db.DoDial()

	if err != nil {
		log.Fatal(err)
	}

	defer s.Close()

	c := s.DB(db.Name()).C("user")
	err = c.Insert(user)

	if err != nil {
		log.Fatal(err)
	}

	return user
}

func FindUserById(id string) models.User {
	db := dbconfig.DB{}
	user := models.User{}

	s, err := db.DoDial()

	if err != nil {
		log.Fatal(err)
	}

	defer s.Close()

	c := s.DB(db.Name()).C("user")

	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&user)

	if err != nil {
		log.Fatal(err)
	}

	return user

}

func FindAllUsers() models.Users {
	db := dbconfig.DB{}
	listUser := models.Users{}
	s, err := db.DoDial()

	if err != nil {
		log.Fatal(err)
	}

	defer s.Close()

	c := s.DB(db.Name()).C("user")

	err = c.Find(bson.M{}).All(&listUser)

	if err != nil {
		log.Fatal(err)
	}

	return listUser
}
