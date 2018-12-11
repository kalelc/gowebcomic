package main

import "log"

func HandlerError(e error) {
	if e != nil {
		log.Fatalf("%s", e)
		panic(e)
	}
}
