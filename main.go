package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	. "github.com/sh9zad/todo-api-go/config"
	. "github.com/sh9zad/todo-api-go/dataaccess"
	. "github.com/sh9zad/todo-api-go/models"
	"gopkg.in/mgo.v2/bson"
)

var config = Config{}
var doa = TodosDataAccess{}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// GetTodos to return all the todos in the db
func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := doa.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, todos)
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

// FindTodoEndPoint get the todo
func FindTodoEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := doa.FindByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Todo ID")
		return
	}
	respondWithJSON(w, http.StatusOK, movie)
}

// UpdateTodoEndpoint update
func UpdateTodoEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := doa.UpdateTodo(todo); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Read()

	doa.Server = config.Server
	doa.Database = config.Database
	doa.Connect()
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/todo", GetTodos).Methods("GET")
	r.HandleFunc("/todo", CreateTodo).Methods("POST")
	r.HandleFunc("/todo/{id}", FindTodoEndPoint).Methods("GET")
	r.HandleFunc("/todo", UpdateTodoEndpoint).Methods("PUT")

	//Lines here are to avoid the CORS issues.
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	fmt.Println("Server started on port: 3060")

	if err := http.ListenAndServe(":3060", handlers.CORS(allowedOrigins, allowedMethods)(r)); err != nil {
		log.Fatal(err)
	}
}
