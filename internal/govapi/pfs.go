package govapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const BatchSize = 500

var ErrNoMoreData = errors.New("no data for the requested batch")

func GetPFSInfo(hostURL, accessToken string, batchNumber int, effectiveStartTimestamp *string) ([]StationNode, error) {
	var resBody []StationNode
	err := getHandler(hostURL, "/api/v1/pfs", accessToken, batchNumber, effectiveStartTimestamp, &resBody)
	return resBody, err
}

func GetPFSFuelPrices(hostURL, accessToken string, batchNumber int, effectiveStartTimestamp *string) ([]NodeFuelPrice, error) {
	var resBody []NodeFuelPrice
	err := getHandler(hostURL, "/api/v1/pfs/fuel-prices", accessToken, batchNumber, effectiveStartTimestamp, &resBody)
	return resBody, err
}

func getHandler(hostURL, path, accessToken string, batchNumber int, effectiveStartTimestamp *string, responseBody any) error {
	query := url.Values{}
	query.Set("batch-number", strconv.Itoa(batchNumber))
	if effectiveStartTimestamp != nil {
		query.Set("effective-start-timestamp", *effectiveStartTimestamp)
	}

	url := strings.TrimRight(hostURL, "/") + path + "?" + query.Encode()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return ErrNoMoreData
	}

	// if error, report it
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf(
			"request failed: status=%d body=%s",
			res.StatusCode,
			string(body),
		)
	}

	return json.NewDecoder(res.Body).Decode(&responseBody)
}
