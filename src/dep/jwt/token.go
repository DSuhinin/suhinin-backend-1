package jwt

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dsuhinin/suhinin-backend-1/core/errors"

	"github.com/dgrijalva/jwt-go"
)

// Provider provides an interface to work with JWT Token.
type Provider interface {
	// Generate generates JWT token according to provided data.
	Generate(id int, email string) (string, error)
	// Verify makes JWT token verification.
	Verify(tokenString string) (int, error)
}

// Token represents JWT token generator.
type Token struct {
	key string
}

// NewToken creates new JWT token generator.
func NewToken(key string) *Token {
	return &Token{
		key: key,
	}
}

// Generate generates JWT token according to provided data.
func (g Token) Generate(id int, email string) (string, error) {

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

// Verify makes JWT token verification.
func (g Token) Verify(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(g.key), nil
	})

	if err != nil {
		return 0, errors.WithMessage(err, "impossible to parse JWT token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.Atoi(claims["jti"].(string))
		if err != nil {
			return 0, errors.WithMessage(err, "impossible to convert `user_id` to int")
		}
		return userID, nil
	}

	return 0, errors.New("provided JWT token is invalid")
}
