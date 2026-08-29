package govapi

type GovGenerateRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type GovGenerateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		AccessToken           string `json:"access_token"`
		TokenType             string `json:"token_type"`
		ExpiresIn             int    `json:"expires_in"` // seconds
		RefreshToken          string `json:"refresh_token"`
		RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"` // seconds
	} `json:"data"`
}

type GovRefreshRequest struct {
	ClientID     string `json:"client_id"`
	RefreshToken string `json:"refresh_token"`
}

type GovRefreshResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"` // seconds
	} `json:"data"`
}
