package repository

import (
	"github.com/vitorgusta/red-venture/api/dbconfig"
	"github.com/vitorgusta/red-venture/api/models"
	"gopkg.in/mgo.v2/bson"
)

func CreateWidget(widget models.Widget) models.Widget {

	db := dbconfig.DB{}
	//widget.Id = bson.NewObjectId()

	s, err := db.DoDial()

	if err != nil {
		//TODO ADD LOG
	}

	defer s.Close()

	c := s.DB(db.Name()).C("widget")

	err = c.Insert(widget)

	if err != nil {
		//TODO ADD LOG
	}

	return widget
}

func FindWidgetById(id string) models.Widget {
	db := dbconfig.DB{}
	widget := models.Widget{}

	s, err := db.DoDial()

	if err != nil {
		//TODO ADD LOG
	}

	defer s.Close()

	c := s.DB(db.Name()).C("widget")

	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&widget)

	if err != nil {
		//TODO ADD LOG
	}

	return widget

}

func FindAllWidget() models.Widgets {
	db := dbconfig.DB{}
	listWidget := models.Widgets{}

	s, err := db.DoDial()

	if err != nil {
		//TODO ADD LOG
	}

	defer s.Close()

	c := s.DB(db.Name()).C("widget")

	err = c.Find(bson.M{}).All(&listWidget)

	if err != nil {
		//TODO ADD LOG
	}

	return listWidget
}

func UpdateWidget(widget models.Widget, id string) (models.Widget, error) {
	db := dbconfig.DB{}

	s, err := db.DoDial()

	if err != nil {
		return widget, err
	}

	defer s.Close()
	c := s.DB(db.Name()).C("widget")

	colQuerier := bson.M{"_id": bson.ObjectIdHex(id)}
	change := bson.M{"$set": bson.M{"name": widget.Name, "price": widget.Price, "color": widget.Color, "melts": widget.Melts, "inventory": widget.Inventory}}

	err = c.Update(colQuerier, change)
	if err != nil {
		return widget, err
	}

	return widget, nil
}
