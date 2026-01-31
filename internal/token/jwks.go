package token

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"math/big"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type JWK struct {
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	N   string `json:"n"`
	E   string `json:"e"`
}

type JWKS struct {
	Keys []JWK `json:"keys"`
}

const bitSize = 2048

var (
	cachedJWKS  *JWKS
	cacheExpiry time.Time
	cachedKeys  map[string]*rsa.PublicKey
)

func NewJKWS(privateKeyPEM string) (*JWKS, error) {
	block, _ := pem.Decode([]byte(privateKeyPEM))
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pubKey := key.PublicKey
	v7, _ := uuid.NewV7()
	jwk := JWK{
		Kid: v7.String(),
		Kty: "RSA",
		Alg: "RS256",
		Use: "sig",
		N:   base64.URLEncoding.EncodeToString(pubKey.N.Bytes()),
		E:   base64.URLEncoding.EncodeToString(big.NewInt(int64(pubKey.E)).Bytes()),
	}

	return &JWKS{
		Keys: []JWK{jwk},
	}, nil
}

func GeneratePEM() (string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return "", err
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return "", err
	}

	// Get ASN.1 DER format, convert to PEM
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)
	return string(privatePEM), nil
}

func GetJWKS(url string) (*JWKS, error) {
	if cachedJWKS != nil && time.Now().Before(cacheExpiry) {
		return cachedJWKS, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&cachedJWKS)
	if err != nil {
		return nil, err
	}

	cacheExpiry = time.Now().Add(24 * time.Hour)
	return cachedJWKS, nil
}
