package govapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetToken(hostURL, clientID, clientSecret string) (*GovGenerateResponse, error) {
	url := strings.TrimRight(hostURL, "/") + "/api/v1/oauth/generate_access_token"
	body := &GovGenerateRequest{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	var resBody GovGenerateResponse
	err := authHandler(url, body, &resBody)
	return &resBody, err
}

func RefreshToken(hostURL, clientID, refreshToken string) (*GovRefreshResponse, error) {
	url := strings.TrimRight(hostURL, "/") + "/api/v1/oauth/regenerate_access_token"
	body := &GovRefreshRequest{
		ClientID:     clientID,
		RefreshToken: refreshToken,
	}

	var resBody GovRefreshResponse
	err := authHandler(url, body, &resBody)
	return &resBody, err
}

func authHandler(url string, requestBody any, responseBody any) error {
	reqString, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(reqString))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

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
