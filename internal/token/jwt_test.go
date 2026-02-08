package token

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateJWT(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	assert.NoError(t, err)

	p := NewJWTProvider(&Options{
		PrivateKey:   key,
		IssuerHost:   "http://localhost:8080",
		AudienceHost: "http://localhost:8080",
	})

	token, err := p.CreateJWT("zak@test.com", "Zak", "Dowsett")
	assert.NoError(t, err)

	claims, err := p.ValidateJWT(token)
	assert.NoError(t, err)

	assert.Equal(t, "Zak", claims.FirstName)
	assert.Equal(t, "Dowsett", claims.LastName)
	assert.Equal(t, "zak@test.com", claims.Subject)
	assert.Equal(t, "http://localhost:8080", claims.Issuer)
	assert.Equal(t, "http://localhost:8080", claims.Audience[0])
	assert.WithinDuration(t, time.Now().Add(5*time.Minute), claims.ExpiresAt.Time, time.Minute)
}
