package main

import (
	"encoding/json"
	"net/http"
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

	w.Header().Set("Content-Type", "application/json")
	payload, err := json.Marshal(comic)

	HandlerError(err)

	w.Write(payload)
}
