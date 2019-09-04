pack]																																																																																																																																																																																																																																																																																																																																																																																																																			age dataaccess

import (
	"log"

	. "github.com/sh9zad/todo-api-go/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TodosDataAccess the struct
type TodosDataAccess struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	//something something.
	COLLECTION = "todos"
)

// Connect Establish connection
func (t *TodosDataAccess) Connect() {
	session, err := mgo.Dial(t.Server)

	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(t.Database)
}

//FindAll todos
func (t *TodosDataAccess) FindAll() ([]Todo, error) {
	var todos []Todo
	err := db.C(COLLECTION).Find(bson.M{}).All(&todos)
	return todos, err
}

// FindByID get the record.
func (t *TodosDataAccess) FindByID(id string) (Todo, error) {
	var todo Todo
	err := db.C(COLLECTION).Find(bson.M{"id": bson.ObjectIdHex(id)}).One(&todo)
	return todo, err
}

// Insert todo
func (t *TodosDataAccess) Insert(todo Todo) error {
	err := db.C(COLLECTION).Insert(&todo)
	return err
}

// UpdateTodo updated
func (t *TodosDataAccess) UpdateTodo(todo Todo) error {
	err := db.C(COLLECTION).Update(bson.M{"id": todo.ID}, &todo)
	return err
}

// DeleteTodo delete
func (t *TodosDataAccess) DeleteTodo(todo Todo) error {
	err := db.C(COLLECTION).Remove(&todo)
	return err
}
