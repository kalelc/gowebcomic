package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const ComicPATH = "data/"

func ListFiles() {
	files, err := ioutil.ReadDir("./data")

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func SaveFile(comic *Comic) {
	data, err := json.Marshal(comic)

	HandlerError(err)

	f, err := os.Create(ComicPATH + strconv.Itoa(comic.Num))

	HandlerError(err)

	f.Write(data)
}
