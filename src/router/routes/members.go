package routes

import (
	"net/http"

	coreHTTP "github.com/dsuhinin/suhinin-backend-1/core/http"
	"github.com/dsuhinin/suhinin-backend-1/core/http/response"

	"github.com/dsuhinin/suhinin-backend-1/src/dao/auth/repository"
	"github.com/dsuhinin/suhinin-backend-1/src/dep/jwt"
	"github.com/dsuhinin/suhinin-backend-1/src/middleware"
	"github.com/dsuhinin/suhinin-backend-1/src/transport"
)

// Available service routes.
const (
	GetMembersRoute = "/members"
)

// InitMembersRouteList makes an initialization of /members routes.
func InitMembersRouteList(
	r coreHTTP.RouterProvider,
	t *transport.Transport,
	jwtToken jwt.Provider,
	tokenRepository repository.TokenRepositoryProvider,
) {
	r.Get(GetMembersRoute, func(req *http.Request) response.Provider {
		return middleware.WithAuthorization(
			req,
			t.GetMembers,
			jwtToken,
			tokenRepository,
		)
	})
}
