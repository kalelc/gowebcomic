package main

import (
	"flag"
	"log"
	"os"
	"strconv"
)

var number = flag.Int("number", 1, "Is the number of commic to downloads")

func main() {
	flag.Parse()

	num := *number

	var status bool
	var comic *Comic
	var err error

	comic, status = SearchFile(strconv.Itoa(num))

	HandlerError(err)

	if !status {
		comic, err = Get(num)
		HandlerError(err)
	}

	if err := templ.Execute(os.Stdout, comic); err != nil {
		log.Fatal(err)
	}
}
