package service

import (
	"errors"
	"log/slog"
	"sync"
	"time"

	"main/internal/govapi"
)

type GovAPIService interface {
	GetAllStations(effectiveStartTimestamp *time.Time) ([]govapi.StationNode, error)
	GetAllFuelPrices(effectiveStartTimestamp *time.Time) ([]govapi.NodeFuelPrice, error)
}

type govAPIService struct {
	clientID     string
	clientSecret string
	apiURL       string

	mu               sync.Mutex
	accessToken      *string
	expiresAt        *time.Time
	refreshToken     *string
	refreshExpiresAt *time.Time
}

func NewGovAPIService(clientID, clientSecret, apiURL string) GovAPIService {
	return &govAPIService{
		clientID:     clientID,
		clientSecret: clientSecret,
		apiURL:       apiURL,
	}
}

func (s *govAPIService) GetAllStations(effectiveStartTimestamp *time.Time) ([]govapi.StationNode, error) {
	token, err := s.getValidAccessToken()
	if err != nil {
		return nil, err
	}

	stations := make([]govapi.StationNode, 0)
	for batchNo := 1; ; batchNo++ {
		res, err := govapi.GetPFSInfo(s.apiURL, token, batchNo, toString(effectiveStartTimestamp))
		// a 404 means either past the last batch, or nothing changed since the timestamp
		if errors.Is(err, govapi.ErrNoMoreData) {
			break
		}
		if err != nil {
			return nil, err
		}

		stations = append(stations, res...)

		// each batch holds up to 500 forecourts, so a short batch is the last
		if len(res) < govapi.BatchSize {
			break
		}
	}

	return stations, nil
}

func (s *govAPIService) GetAllFuelPrices(effectiveStartTimestamp *time.Time) ([]govapi.NodeFuelPrice, error) {
	token, err := s.getValidAccessToken()
	if err != nil {
		return nil, err
	}

	prices := make([]govapi.NodeFuelPrice, 0)
	for batchNo := 1; ; batchNo++ {
		res, err := govapi.GetPFSFuelPrices(s.apiURL, token, batchNo, toString(effectiveStartTimestamp))
		if errors.Is(err, govapi.ErrNoMoreData) {
			break
		}
		if err != nil {
			return nil, err
		}

		prices = append(prices, res...)

		if len(res) < govapi.BatchSize {
			break
		}
	}

	return prices, nil
}

// SECTION: private functions for maintaining tokens

func (s *govAPIService) getValidAccessToken() (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// treat tokens as expired a minute early so one isn't used mid-request expiry
	now := time.Now().Add(time.Minute)

	// CASE 1 - access token is still valid
	if s.accessToken != nil && s.expiresAt != nil && now.Before(*s.expiresAt) {
		return *s.accessToken, nil
	}

	// CASE 2 - refresh token is still valid, regenerate the access token.
	// The API issues no new refresh token here, so the existing one is kept.
	if s.refreshToken != nil && s.refreshExpiresAt != nil && now.Before(*s.refreshExpiresAt) {
		res, err := govapi.RefreshToken(s.apiURL, s.clientID, *s.refreshToken)
		if err == nil {
			s.accessToken = &res.Data.AccessToken
			expiresAt := time.Now().Add(time.Duration(res.Data.ExpiresIn) * time.Second)
			s.expiresAt = &expiresAt
			return *s.accessToken, nil
		}
		slog.Warn("token refresh failed, requesting a new token", "error", err)
	}

	// CASE 3 - both nil or invalid, generate a fresh token pair
	res, err := govapi.GetToken(s.apiURL, s.clientID, s.clientSecret)
	if err != nil {
		return "", err
	}

	// assign vals
	s.accessToken = &res.Data.AccessToken
	expiresAt := time.Now().Add(time.Duration(res.Data.ExpiresIn) * time.Second)
	s.expiresAt = &expiresAt
	s.refreshToken = &res.Data.RefreshToken
	refreshExpiresAt := time.Now().Add(time.Duration(res.Data.RefreshTokenExpiresIn) * time.Second)
	s.refreshExpiresAt = &refreshExpiresAt

	return *s.accessToken, nil
}

func toString(t *time.Time) *string {
	if t == nil {
		return nil
	}

	res := t.UTC().Format("2006-01-02")
	return &res
}
