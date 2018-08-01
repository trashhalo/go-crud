package main

import (
	"github.com/gorilla/mux"
	"github.com/trashhalo/go-crud/db"
	"net/http"
	"strings"
)

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// TODO better error handling
		panic(err)
	}
	form := r.Form
	todo := form["todo"]
	err = db.CreateTodo(strings.Join(todo, ""))
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// TODO better error handling
		panic(err)
	}
	id := mux.Vars(r)["id"]
	form := r.Form
	done := strings.Join(form["done"], "") == "true"
	err = db.UpdateTodo(id, done)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
