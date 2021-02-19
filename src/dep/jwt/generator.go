package jwt

import (
	"strconv"
	"time"

	"github.com/dsuhinin/suhinin-backend-1/core/errors"

	"github.com/dgrijalva/jwt-go"
)

// GeneratorProvider provides an interface to work with JWT Generator.
type GeneratorProvider interface {
	// Generate generates JWT token according to provided data.
	Generate(id int, email string) (string, error)
}

// AuthTokenClaim this is the claim object which gets parsed from the authorization header
type AuthTokenClaim struct {
	*jwt.StandardClaims
}

// Generator represents JWT token generator.
type Generator struct {
	key string
}

// NewGenerator creates new JWT token generator.
func NewGenerator(key string) *Generator {
	return &Generator{
		key: key,
	}
}

// Generate generates JWT token according to provided data.
func (g Generator) Generate(id int, email string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        strconv.Itoa(id),
		Issuer:    email,
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(g.key))
	if err != nil {
		return "", errors.WithMessage(err, "impossible to generate JWT token")
	}

	return tokenString, nil
}
