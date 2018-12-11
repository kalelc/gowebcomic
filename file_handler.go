package main

import (
	"encoding/json"
	"os"
	"strconv"
)

const ComicPATH = "data/"

func SaveFile(comic *Comic) {
	data, err := json.Marshal(comic)

	HandlerError(err)

	f, err := os.Create(ComicPATH + strconv.Itoa(comic.Num))

	HandlerError(err)

	f.Write(data)
}
