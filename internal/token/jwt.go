package token

import (
	"context"
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	expiryTime        = time.Minute * 5    // 5 minutes
	refreshExpiryTime = time.Hour * 24 * 7 // 1 week
)

type Options struct {
	PrivateKey   *rsa.PrivateKey
	IssuerHost   string
	AudienceHost string
}

type JWTProvider struct {
	options *Options
}

func NewJWTProvider(options Options) *JWTProvider {
	return &JWTProvider{
		options: &options,
	}
}

func (p *JWTProvider) CreateJWT(email, firstName, lastName string) (string, error) {
	idFunc, _ := uuid.NewV7()
	now := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, UserClaim{
		FirstName: firstName,
		LastName:  lastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        idFunc.String(),
			Issuer:    p.options.IssuerHost,
			Subject:   email,
			Audience:  jwt.ClaimStrings{p.options.AudienceHost},
			ExpiresAt: &jwt.NumericDate{Time: now.Add(expiryTime)},
			IssuedAt:  &jwt.NumericDate{Time: now},
		},
	})
	s, err := token.SignedString(p.options.PrivateKey)
	return s, err
}

func (p *JWTProvider) CreateRefreshJWT() (string, uuid.UUID, error) {
	tokenID, _ := uuid.NewV7()
	now := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{
		ID:        tokenID.String(),
		Issuer:    p.options.IssuerHost,
		Audience:  jwt.ClaimStrings{p.options.AudienceHost},
		ExpiresAt: &jwt.NumericDate{Time: now.Add(refreshExpiryTime)},
		IssuedAt:  &jwt.NumericDate{Time: now},
	})
	s, err := token.SignedString(p.options.PrivateKey)
	return s, tokenID, err
}

func (p *JWTProvider) ValidateJWT(tokenString string) (*UserClaim, error) {
	ctx := context.Background()
	url := p.options.IssuerHost + "/.well-known/jwks.json"
	k, err := keyfunc.NewDefaultCtx(ctx, []string{url}) // Context is used to end the refresh goroutine.

	token, err := jwt.Parse(tokenString, k.Keyfunc, jwt.WithAudience(p.options.AudienceHost), jwt.WithIssuer(p.options.IssuerHost), jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Alg()}))
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaim); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
