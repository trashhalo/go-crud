package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/trashhalo/go-crud/db"
	"net/http"
)

func main() {
	err := db.OpenDB()
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/todo", CreateTodoHandler).Methods("POST")
	r.HandleFunc("/todo/{id}", UpdateTodoHandler).Methods("PUT")
	http.ListenAndServe(":8080", handlers.HTTPMethodOverrideHandler(r))
}
