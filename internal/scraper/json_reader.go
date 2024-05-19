package scraper

import (
	"encoding/json"
	"fmt"
	"main/api/fuelfinder"
	"main/internal/models"
	"main/internal/sanitiser"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	dateTimeLayout = "02/01/2006 15:04:05"
)

// the main must return (interface{}, error)
func ReadJsonFrom(url string) ([]*fuelfinder.StationItem, error) {

	if url == "" {
		return nil, fmt.Errorf("url cannot be empty")
	}

	// get the json vals
	req, _ := createRequest(url)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// close response
	defer resp.Body.Close()

	// Create an empty struct and unmarshal the body
	var priceData models.PriceDataset
	err = json.NewDecoder(resp.Body).Decode(&priceData)
	if err != nil {
		return nil, err
	}

	// parse the time. Can we read it?
	// NOTE - jetlocal use the wrong string format
	priceData.LastUpdated = strings.Replace(priceData.LastUpdated, "-", "/", -1)
	createdAt, err := time.Parse(dateTimeLayout, priceData.LastUpdated)
	if err != nil {
		return nil, err
	}

	// TODO: There was a data field providing a different price format, where is it?
	// TODO: Sense check the price data, if it looks off, invalidate the json and report it

	// convert inputted format to 2D store format
	var convertedData []*fuelfinder.StationItem
	for _, s := range priceData.Stations {
		// append to list
		convertedData = append(convertedData, &fuelfinder.StationItem{
			CreatedAt: createdAt.Unix(),
			SiteId:    s.SiteId,
			Brand:     s.Brand,
			Address:   s.Address,
			Postcode:  s.Postcode,
			Location: &fuelfinder.Location{
				Type: "Point",
				Coordinates: []float32{
					float32(sanitiser.ToFloat(s.Location.Longitude)),
					float32(sanitiser.ToFloat(s.Location.Latitude)),
				},
			},
			E5:  s.Prices.E5,
			E10: s.Prices.E10,
			B7:  s.Prices.B7,
			Sdv: s.Prices.SDV,
		})
	}

	return convertedData, nil
}

func createRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// NOTE: need to add the headers to emulate an actual request
	r := regexp.MustCompile(`http.?:\/\/(.+?)\/`)  // gets host name
	host := r.FindAllStringSubmatch(url, -1)[0][1] // 0 gets first regex match, 1 gets match group within the regex

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Host", host)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:124.0) Gecko/20100101 Firefox/124.0")
	return req, nil
}
