package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var status bool
	var comic *Comic

	comic, status = SearchFile(id)

	num, err := strconv.Atoi(vars["id"])

	HandlerError(err)

	if !status {
		comic, err = Get(num)
		HandlerError(err)
	}

	if err := templ.Execute(os.Stdout, comic); err != nil {
		log.Fatal(err)
	}

}
