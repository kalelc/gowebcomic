package main

import (
	"os"
	"strconv"
)

func main() {
	var status bool
	for _, x := range os.Args[1:] {
		status = SearchFile(x)

		if !status {
			num, err := strconv.Atoi(x)
			HandlerError(err)
			Get(num)
		}
	}
}
