package main

import (
	"fmt"
	"net/http"
)

const Port = "3001"

func main() {
	router := NewRouter()
	fmt.Println("* Listening on " + Port + " port")
	fmt.Println("* Use Ctrl-C to stop")
	http.ListenAndServe(":"+Port, router)
}
