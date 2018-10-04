package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/sh9zad/todo-api-go/DataAccess"
	. "github.com/sh9zad/todo-api-go/models"
	"gopkg.in/mgo.v2/bson"
)

var doa = TodosDataAccess{}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

/**/
func GetTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// CreateTodo something
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	todo.ID = bson.NewObjectId()
	if err := doa.Insert(todo); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, todo)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

// respondWithJSON as
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/todo", GetTodos).Methods("GET")

	if err := http.ListenAndServe(":3060", r); err != nil {
		log.Fatal(err)
	}
}
