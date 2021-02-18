package transport

import (
	"net/http"

	"github.com/dsuhinin/suhinin-backend-1/core/http/response"
)

/// GetMembers handles `GET /members` route.
func (h *Transport) GetMembers(r *http.Request) response.Provider {
	resp, err := h.serviceController.GetMembers()
	if err != nil {
		return response.New(err)
	}

	return resp
}
