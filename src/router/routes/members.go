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
	GetMembersRoute = "/members"
)

// InitMembersRouteList makes an initialization of /members routes.
func InitMembersRouteList(
	r kitHTTP.RouterProvider, t *transport.Transport, authRepository auth.RepositoryProvider,
) {
	r.Get(GetMembersRoute, func(req *http.Request) response.Provider {
		return middleware.WithAuthorization(
			req,
			t.GetMembers,
			authRepository,
		)
	})
}
