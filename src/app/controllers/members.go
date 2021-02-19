package controllers

import (
	coreResponse "github.com/dsuhinin/suhinin-backend-1/core/http/response"

	"github.com/dsuhinin/suhinin-backend-1/src/api/response"
)

// GetMembers handles GET /members route.
func (c Controller) GetMembers() (coreResponse.Provider, error) {
	return coreResponse.New(
		response.NewMembers(
			"this text could be accessible only in case of success authentication and goes from the service",
		),
	), nil
}
