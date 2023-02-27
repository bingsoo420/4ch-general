package main

import (
	"encoding/json"
	"fmt"
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

	if j, _ := json.Marshal(mappings); true {
		fmt.Println(string(j))
	}
}
