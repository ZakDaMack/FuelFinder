package scraper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	dateTimeLayout = "02/01/2006 15:04:05"
)

func ReadJsonFrom(url string) ([]StationDataset, error) {
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
	var priceData PriceDataset
	err = json.NewDecoder(resp.Body).Decode(&priceData)
	if err != nil {
		return nil, err
	}

	// parse the time. Can we read it?
	// NOTE - jetlocal use the wrong string format
	priceData.LastUpdated = strings.ReplaceAll(priceData.LastUpdated, "-", "/")
	createdAt, err := time.Parse(dateTimeLayout, priceData.LastUpdated)
	if err != nil {
		return nil, err
	}

	// TODO: There was a data field providing a different price format, where is it?
	// TODO: Sense check the price data, if it looks off, invalidate the json and report it

	// move last updated to each station
	for i := range priceData.Stations {
		priceData.Stations[i].CreatedAt = &createdAt
	}

	return priceData.Stations, nil
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
