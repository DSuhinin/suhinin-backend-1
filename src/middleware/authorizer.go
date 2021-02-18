package middleware

import (
	"net/http"
	"strings"

	"github.com/dsuhinin/suhinin-backend-1/core/http/response"

	"github.com/dsuhinin/suhinin-backend-1/src/api"
	"github.com/dsuhinin/suhinin-backend-1/src/dao/auth"
)

// represents Authentication header list.
const (
	ServerAuthorizationKeyHeader = "Authorization"
)

// WithAuthorization makes validation of Authorization header.
func WithAuthorization(
	req *http.Request, callback Callback, authRepository auth.RepositoryProvider,
) response.Provider {

	authorization := req.Header.Get(ServerAuthorizationKeyHeader)
	if authorization == "" {
		return response.New(api.ServerAuthorizationHeaderEmptyError)
	}

	authorizationParts := strings.Split(authorization, " ")
	if len(authorizationParts) != 2 {
		return response.New(api.ServerAuthorizationHeaderEmptyError)
	}

	return callback(req)
}
