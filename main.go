package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	Sync()
}

func Sync() {
	for i := 1; i <= Limit; i++ {
		result, err := Get(i)
		if err == nil {
			SaveFile(result)
		}
	}
}

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
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
