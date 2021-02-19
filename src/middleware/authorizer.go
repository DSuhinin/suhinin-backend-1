package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dsuhinin/suhinin-backend-1/core/http/response"

	"github.com/dsuhinin/suhinin-backend-1/src/api"
	"github.com/dsuhinin/suhinin-backend-1/src/dao/auth/repository"
	"github.com/dsuhinin/suhinin-backend-1/src/dep/jwt"
)

// represents Authentication header list.
const (
	ServerAuthorizationKeyHeader = "Authorization"
)

// supported authorization types.
const (
	BearerAuthorization = "Bearer"
)

type ctxKey struct{}

// Context `user_id` key.
var (
	ContextUserIDKey    ctxKey
	ContextUserTokenKey ctxKey
)

// WithAuthorization makes validation of Authorization header.
func WithAuthorization(
	req *http.Request, callback Callback, jwtToken jwt.Provider, tokenRepository repository.TokenRepositoryProvider,
) response.Provider {

	authorization := req.Header.Get(ServerAuthorizationKeyHeader)
	if authorization == "" {
		return response.New(api.ServerAuthorizationHeaderEmptyError)
	}

	authorizationParts := strings.Split(authorization, " ")
	if len(authorizationParts) != 2 {
		return response.New(api.ServerAuthorizationHeaderEmptyError)
	}

	if authorizationParts[0] != BearerAuthorization {
		return response.New(api.ServerAuthorizationHeaderEmptyError)
	}

	userID, err := jwtToken.Verify(authorizationParts[1])
	if err != nil {
		return response.New(
			api.UnauthorizedRequestError.WithMessage("impossible verify JWT token: %+v", err),
		)
	}

	if err := tokenRepository.GetByUserIDAndToken(userID, authorizationParts[1]); err != nil {
		return response.New(
			api.UnauthorizedRequestError.WithMessage("impossible to get JWT token: %+v", err),
		)
	}

	req = req.WithContext(context.WithValue(req.Context(), ContextUserIDKey, userID))
	req = req.WithContext(context.WithValue(req.Context(), ContextUserTokenKey, authorizationParts[1]))

	return callback(req)
}
