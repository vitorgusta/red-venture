package repository

import (
	"fmt"

	"github.com/labstack/gommon/log"

	"github.com/vitorgusta/red-venture/api/dbconfig"
	"github.com/vitorgusta/red-venture/api/models"
	"gopkg.in/mgo.v2/bson"
)

func CreateUser(user models.User) (models.User, error) {
	db := dbconfig.DB{}
	fmt.Println("vitorasidmoasidm", db)
	log.Debug("init db", db)
	user.Id = bson.NewObjectId()

	s, err := db.DoDial()

	if err != nil {
		return user, err
	}

	defer s.Close()

	c := s.DB(db.Name()).C("user")
	fmt.Println(c)
	log.Debug("init c", c)
	err = c.Insert(user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func FindUserById(id string) (models.User, error) {
	db := dbconfig.DB{}
	user := models.User{}

	s, err := db.DoDial()

	if err != nil {
		return user, err
	}

	defer s.Close()

	c := s.DB(db.Name()).C("user")

	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&user)

	if err != nil {
		return user, err
	}

	return user, nil

}

func FindAllUsers() (models.Users, error) {
	db := dbconfig.DB{}
	listUser := models.Users{}
	log.Debug("init db", db)
	fmt.Println("vitorasidmoasidm", db)
	s, err := db.DoDial()

	if err != nil {
		return listUser, err
	}

	defer s.Close()

	c := s.DB(db.Name()).C("user")
	fmt.Println("MONGO", c)
	log.Debug("init c", c)
	err = c.Find(bson.M{}).All(&listUser)

	if err != nil {
		return listUser, err
	}

	return listUser, nil
}
