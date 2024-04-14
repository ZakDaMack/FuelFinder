package main

import (
	"encoding/json"
	"fmt"
	"main/pkg/database"
)

func main() {
	db, _ := database.NewMongoConnection("uri", "ofd")
	locs, err := db.QueryArea(0, 10, 20)
	if err != nil {
		fmt.Println(err)
	}

	res, _ := json.MarshalIndent(locs, "", "  ")
	fmt.Println(string(res))
}
