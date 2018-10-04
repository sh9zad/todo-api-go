package models

import "gopkg.in/mgo.v2/bson"

/**/
type Todo struct {
	ID      bson.ObjectId `bson:"id" json:"id"`
	Title   string        `bson:"title" json:"title"`
	DueDate string        `bson:"due_date" json:"due_date"`
}
