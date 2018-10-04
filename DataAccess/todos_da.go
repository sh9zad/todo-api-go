package DataAccess

import (
	"log"

	. "github.com/sh9zad/todo-api-go/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TodosDataAccess struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "todos"
)

// Connect Establish connection
func (m *TodosDataAccess) Connect() {
	session, err := mgo.Dial(m.Server)

	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(m.Database)
}

//FindAll todos
func (m *TodosDataAccess) FindAll() ([]Todo, error) {
	var todos []Todo
	err := db.C(COLLECTION).Find(bson.M{}).All(&todos)
	return todos, err
}

// Insert todo
func (m *TodosDataAccess) Insert(todo Todo) error {
	err := db.C(COLLECTION).Insert(&todo)
	return err
}
