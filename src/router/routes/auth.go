package routes

import (
	"net/http"

	kitHTTP "github.com/dsuhinin/suhinin-backend-1/core/http"

	"github.com/dsuhinin/suhinin-backend-1/core/http/response"
	"github.com/dsuhinin/suhinin-backend-1/src/dao/auth"
	"github.com/dsuhinin/suhinin-backend-1/src/middleware"
	"github.com/dsuhinin/suhinin-backend-1/src/transport"
)

// Available service routes.
const (
	SigninRoute  = "/auth/signin"
	SignupRoute  = "/auth/signup"
	SignoutRoute = "/auth/signout"
)

// InitAuthRouteList makes an initialization of /auth routes.
func InitAuthRouteList(
	router kitHTTP.RouterProvider,
	transport *transport.Transport,
	authRepository auth.RepositoryProvider,
) {
	router.Post(SigninRoute, func(req *http.Request) response.Provider {
		return transport.Signin(req)
	})
	router.Post(SignupRoute, func(req *http.Request) response.Provider {
		return transport.Signup(req)
	})
	router.Get(SignoutRoute, func(req *http.Request) response.Provider {
		return middleware.WithAuthorization(
			req,
			transport.Signout,
			authRepository,
		)
	})
}
