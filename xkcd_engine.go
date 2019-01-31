package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const ServerURL = "https://xkcd.com"

func Get(num int) (*Comic, error) {
	url := ServerURL + "/" + strconv.Itoa(num) + "/info.0.json"
	fmt.Println(url)
	resp, err := http.Get(url)

	HandlerError(err)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("element not found: %s", resp.Status)
	}

	var result Comic

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		HandlerError(err)
	}
	resp.Body.Close()

	SaveFile(&result)

	return &result, err
}
