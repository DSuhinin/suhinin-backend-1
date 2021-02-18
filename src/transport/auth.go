package transport

import (
	"net/http"

	"github.com/dsuhinin/suhinin-backend-1/core/http/response"

	"github.com/dsuhinin/suhinin-backend-1/src/api/request"
)

// Signin handles `POST /auth/signin` route.
func (h *Transport) Signin(r *http.Request) response.Provider {
	req := request.Signin{}
	if err := unmarshal(r.Body, &req); err != nil {
		return err
	}

	resp, err := h.serviceController.Signin(&req)
	if err != nil {
		return response.New(err)
	}

	return resp
}

// Signup handles `POST /auth/signup` route.
func (h *Transport) Signup(r *http.Request) response.Provider {
	req := request.Signup{}
	if err := unmarshal(r.Body, &req); err != nil {
		return err
	}

	resp, err := h.serviceController.Signup(&req)
	if err != nil {
		return response.New(err)
	}

	return resp
}

// Signout handles `POST /auth/signout` route.
func (h *Transport) Signout(r *http.Request) response.Provider {
	resp, err := h.serviceController.Signout()
	if err != nil {
		return response.New(err)
	}

	return resp
}
