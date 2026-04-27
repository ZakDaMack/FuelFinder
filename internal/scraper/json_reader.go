package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"regexp"
	"time"
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
		body, _ := io.ReadAll(resp.Body)
		slog.Error("failed to decode json", "error", err, "url", url, "status", resp.StatusCode, "body", string(body))
		return nil, err
	}

	// parse the time. Can we read it?
	createdAt, err := handleDateTime(priceData.LastUpdated)
	if err != nil {
		return nil, err
	}

	// TODO: There was a data field providing a different price format, where is it?
	// TODO: Sense check the price data, if it looks off, invalidate the json and report it

	// move last updated to each station
	for i := range priceData.Stations {
		priceData.Stations[i].CreatedAt = createdAt
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

	//NOTE: USE CASE - Karan Retail Ltd uses invalid cert, disable if host is trl
	if host == "devapi.krlpos.com" {
		http.DefaultTransport.(*http.Transport).TLSClientConfig.InsecureSkipVerify = true
	}

	req.Header.Set("Host", host)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:149.0) Gecko/20100101 Firefox/149.0")
	return req, nil
}

func handleDateTime(dateTimeStr string) (*time.Time, error) {
	// type1 - 02/01/2006 15:04:05
	matchOne := regexp.MustCompile(`\d{2}/\d{2}/\d{4} \d{2}:\d{2}:\d{2}`).MatchString(dateTimeStr)
	if matchOne {
		createdAt, err := time.Parse("02/01/2006 15:04:05", dateTimeStr)
		if err != nil {
			return nil, err
		}
		return &createdAt, nil
	}

	// type2 - 02-01-2006 15:04:05
	matchTwo := regexp.MustCompile(`\d{2}-\d{2}-\d{4} \d{2}:\d{2}:\d{2}`).MatchString(dateTimeStr)
	if matchTwo {
		createdAt, err := time.Parse("02-01-2006 15:04:05", dateTimeStr)
		if err != nil {
			return nil, err
		}
		return &createdAt, nil
	}

	return nil, fmt.Errorf("unrecognized date time format: %s", dateTimeStr)
}
