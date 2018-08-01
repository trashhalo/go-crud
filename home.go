package main

import (
	"net/http"

	"github.com/trashhalo/go-crud/db"
	"github.com/trashhalo/go-crud/web"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := db.ListTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session, _ := store.Get(r, "session")
	var flashes []string
	if session != nil {
		for _, flash := range session.Flashes() {
			flashes = append(flashes, flash.(string))
		}
	}

	data := web.HomeData{
		Todos:   todos,
		Flashes: flashes,
	}
	session.Save(r, w)
	w.Write([]byte(web.Home(data)))
}
