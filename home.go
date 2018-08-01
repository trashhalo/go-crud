package main

import (
	"github.com/trashhalo/go-crud/db"
	"github.com/trashhalo/go-crud/web"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := db.ListTodos()
	if err != nil {
		panic(err)
	}
	w.Write([]byte(web.Home(todos)))
}
