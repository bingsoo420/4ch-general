package main

import (
	"encoding/json"
	"os"
)

var (
	BOARDS = []string{"g", "fit", "biz"}
)

func main() {
	mappings := make(map[string]interface{})

	for _, b := range BOARDS {
		catalog, err := GetBoardCatalog(b)
		if err != nil {
			panic(err)
		}

		generals := BuildGenerals(b, catalog)
		mappings[b] = generals
	}

	jsonPayload, _ := json.Marshal(mappings)
	os.WriteFile("output/mappings.json", jsonPayload, 0644)
}
