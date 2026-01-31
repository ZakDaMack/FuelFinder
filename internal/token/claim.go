package token

import (
	"github.com/golang-jwt/jwt/v5"
)

type UserClaim struct {
	FirstName string
	LastName  string
	jwt.RegisteredClaims
}
