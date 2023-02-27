package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	dat, err := os.ReadFile("./fixtures/g/catalog.json")
	if err != nil {
		panic(err)
	}
	var catalog []Page

	if err := json.Unmarshal(dat, &catalog); err != nil {
		panic(err)
	}

	created := time.Now().Unix()
	generalMaps := BuildGenerals("g", catalog)

	g := General{
		created,
		"g",
		generalMaps,
	}

    if j, _ := json.Marshal(g); true {
		fmt.Println(string(j))
	}
}
