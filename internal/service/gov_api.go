package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"fuelfinder/internal/model"
)

type GovAPIService interface {
}

type govAPIService struct {
  clientID     string
  clientSecret string
  apiURL       string

  client 	  *http.Client

  accessToken *string
  expiresAt   *time.Time
}

func NewGovAPIService(clientID, clientSecret, apiURL string) GovAPIService {
  return &govAPIService{
    clientID:     clientID,
    clientSecret: clientSecret,
    apiURL:       apiURL,
    client:      &http.Client{},
  }
}

func (s *govAPIService) getValidAccessToken() string {
  // get the current access token or refresh if less than a minute left
  hasExpired := s.expiresAt != nil && s.expiresAt.Before(time.Now().Subtract(time.Minute))

  if s.accessToken == nil || hasExpired  {
	  s.fetchAccessToken()
  }
  
  return *s.accessToken
}

// Implement the logic to get an access token from the government API using client credentials
func (s *govAPIService) fetchAccessToken() {
  resp, err := s.client.Post(
	fmt.Sprintf("%s/oauth/token", s.apiURL), 
	"application/x-www-form-urlencoded", 
	strings.NewReader("grant_type=client_credentials&client_id="+s.clientID+"&client_secret="+s.clientSecret+"&scope=fuelfinder.read"),
  )

  if err != nil {
	  return
  }

  defer resp.Body.Close()
  body, err := io.ReadAll(resp.Body)
  if err != nil {
	  return
  }

  // Parse the response body to extract the access token
  var result model.GovTokenResponse
  if err := json.Unmarshal(body, &result); err != nil {
	  return
  }

  // save the data to the service
  s.accessToken = result.accessToken
  s.expiresAt = time.Now().Add(result.expiresIn * time.Seconds)
}