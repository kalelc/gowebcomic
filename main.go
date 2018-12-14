package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	x := os.Args[1:][0]

	var status bool
	var comic *Comic
	var err error

	comic, status = SearchFile(x)

	num, err := strconv.Atoi(x)

	HandlerError(err)

	if !status {
		comic, err = Get(num)
		HandlerError(err)
	}

	if err := templ.Execute(os.Stdout, comic); err != nil {
		log.Fatal(err)
	}
}
