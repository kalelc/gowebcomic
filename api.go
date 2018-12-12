package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Get(num int) {
	url := ServerURL + "/" + strconv.Itoa(num) + "/info.0.json"
	fmt.Println(url)
	resp, err := http.Get(url)

	HandlerError(err)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Errorf("element not found: %s", resp.Status)
	}

	var result Comic
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		HandlerError(err)
	}
	resp.Body.Close()

	SaveFile(&result)
}
