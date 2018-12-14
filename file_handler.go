package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

const ComicPATH = "data/"

func SearchFile(name string) (comic *Comic, status bool) {
	files, err := ioutil.ReadDir("data/")

	HandlerError(err)

	for _, file := range files {
		if name == file.Name() {
			status = true
			jsonFile, err := os.Open("data/" + name)

			HandlerError(err)

			if err := json.NewDecoder(jsonFile).Decode(&comic); err != nil {
				HandlerError(err)
			}
			break
		}
	}
	return
}

func SaveFile(comic *Comic) {
	data, err := json.Marshal(comic)

	HandlerError(err)

	f, err := os.Create(ComicPATH + strconv.Itoa(comic.Num))

	HandlerError(err)

	f.Write(data)
}
