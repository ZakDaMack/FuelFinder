package token

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePEM(t *testing.T) {
	privatePEM, err := GeneratePEM()
	assert.NoError(t, err)
	assert.Contains(t, privatePEM, "-----BEGIN RSA PRIVATE KEY-----")
}

func TestNewJWKS(t *testing.T) {
	privatePEM, err := GeneratePEM()
	assert.NoError(t, err)

	jwks, err := NewJKWS(privatePEM)
	assert.NoError(t, err)
	assert.Len(t, jwks.Keys, 1)

	assert.Equal(t, "RSA", jwks.Keys[0].Kty)
	assert.Equal(t, "RS256", jwks.Keys[0].Alg)
	assert.Equal(t, "sig", jwks.Keys[0].Use)
	assert.NotEmpty(t, jwks.Keys[0].Kid)
	assert.NotEmpty(t, jwks.Keys[0].N)
	assert.NotEmpty(t, jwks.Keys[0].E)
}

func TestGetJKWS(t *testing.T) {
	// sample JWKS URL
	url := "https://www.googleapis.com/oauth2/v3/certs"
	jwks, err := GetJWKS(url)
	log.Print(jwks)

	assert.NoError(t, err)
	assert.NotNil(t, jwks)
	assert.Greater(t, len(jwks.Keys), 0)
}
