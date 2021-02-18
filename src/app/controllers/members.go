package controllers

import (
	coreResponse "github.com/dsuhinin/suhinin-backend-1/core/http/response"
)

// GetMembers handles GEt /members route.
func (c Controller) GetMembers() (coreResponse.Provider, error) {
	return coreResponse.New(nil), nil
}
