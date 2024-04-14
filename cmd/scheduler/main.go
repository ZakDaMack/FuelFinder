package main

import (
	"fmt"
	"main/pkg/database"
	"main/pkg/scraper"
)

func main() {
	links, _ := scraper.GetTableLinks("https://www.gov.uk/guidance/access-fuel-price-data")
	db, _ := database.NewMongoConnection("uri", "ofd")
	for _, val := range links {
		fmt.Println()
		fmt.Println()
		fmt.Println(val)
		fmt.Println("----")
		data, err := scraper.ReadJsonFrom(val)
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}

		fmt.Println("stations:", len(data))
		fmt.Println(data[0])

		err = db.Write(data)
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}
	}
}
