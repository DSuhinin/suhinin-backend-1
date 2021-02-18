package middleware

import (
	"net/http"

	"github.com/dsuhinin/suhinin-backend-1/core/http/response"
)

// Callback route handler callback function.
type Callback func(req *http.Request) response.Provider
