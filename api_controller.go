package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	queryValues := r.URL.Query()
	tokenString := queryValues["jwt"][0]

	var status bool
	var comic *Comic

	comic, status = SearchFile(id)

	num, err := strconv.Atoi(vars["id"])

	HandlerError(err)

	if !status {
		comic, err = Get(num)
		HandlerError(err)
	}

	fmt.Println(tokenString)

	token, err := Auth(tokenString)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		w.Header().Set("Content-Type", "application/json")
		payload, err := json.Marshal(comic)
		fmt.Println(claims)
		w.Write(payload)
	} else {
		fmt.Println(err)
	}

}
