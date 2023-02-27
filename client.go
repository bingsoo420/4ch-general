package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var client *http.Client = nil

func newClient() {
	if client != nil {
		return
	}

	client = &http.Client{}
}

func GetBoardCatalog(board string) ([]Page, error) {
	newClient()

    response, err := client.Get(fmt.Sprintf("https://a.4cdn.org/%s/catalog.json", board))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

    var catalog []Page
    err = json.NewDecoder(response.Body).Decode(&catalog)
    if err != nil {
        return nil, err
    }

	return catalog, nil
}
