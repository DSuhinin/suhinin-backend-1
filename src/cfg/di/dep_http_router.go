package di

import (
	"github.com/dsuhinin/suhinin-backend-1/core/cfg/di"
	"github.com/dsuhinin/suhinin-backend-1/core/http"

	"github.com/dsuhinin/suhinin-backend-1/src/router/routes"
)

// Dependency name.
const (
	DefHTTPRouter = "HTTPRouter"
)

// registerHTTPRouter dependency registrar.
func (c *Container) registerHTTPRouter() error {

	if err := c.RegisterDependency(
		DefHTTPRouter,
		func(ctx di.Context) (interface{}, error) {
			r := http.NewRouter(c.GetLogger())
			routes.InitAuthRouteList(
				r,
				c.GetServiceTransport(),
				c.GetJWTToken(),
				c.GetTokenRepository(),
			)
			routes.InitMembersRouteList(
				r,
				c.GetServiceTransport(),
				c.GetJWTToken(),
				c.GetTokenRepository(),
			)

			return r, nil

		}, nil,
	); nil != err {
		return err
	}

	return nil
}

// GetHTTPRouter dependency retriever.
func (c *Container) GetHTTPRouter() http.RouterProvider {

	return c.Container.Get(DefHTTPRouter).(http.RouterProvider)
}
