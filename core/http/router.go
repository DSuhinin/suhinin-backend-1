package http

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/dsuhinin/suhinin-backend-1/core/errors"
	"github.com/dsuhinin/suhinin-backend-1/core/http/response"
	"github.com/dsuhinin/suhinin-backend-1/core/http/response/serializer"
	"github.com/dsuhinin/suhinin-backend-1/core/log"
)

// RouterProvider provides an interface to work with Application Router.
type RouterProvider interface {
	// GetMuxRouter returns the Request handler.
	GetMuxRouter() *mux.Router
	// HandlerResponse returns the Response handler.
	HandleResponse(http.ResponseWriter, *http.Request, Handler)
	// Get function add to router rule to handle specified route only by http Get method.
	Get(string, Handler)
	// Post function add to router rule to handle specified route only by http Post method.
	Post(string, Handler)
	// Put function add to router rule to handle specified route only by http Put method.
	Put(string, Handler)
	// Patch function add to router rule to handle specified route only by http Patch method.
	Patch(string, Handler)
	// Options function add to router rule to handle specified route only by http Options method.
	Options(string, Handler)
	// Delete function add to router rule to handle specified route only by http Delete method.
	Delete(string, Handler)
	// Trace function add to router rule to handle specified route only by http Trace method.
	Trace(string, Handler)
	// Handle function allows you to accept all types of requests to specified route.
	Handle(string, HandlerFunc)
}

// Router is a Application router.
type Router struct {
	logger          log.Logger
	router          *mux.Router
	errorSerializer serializer.Serializer
}

// NewRouter returns a new Router instance.
func NewRouter(logger log.Logger, options ...func(*Router)) *Router {

	r := Router{
		logger: logger,
		router: mux.NewRouter(),
	}

	for _, option := range options {
		option(&r)
	}

	return &r
}

// GetMuxRouter returns the Request handler.
func (r *Router) GetMuxRouter() *mux.Router {

	return r.router
}

// HandleResponse handles the response.
func (r *Router) HandleResponse(w http.ResponseWriter, req *http.Request, h Handler) {

	resp := h(req)
	if resp.IsError() {
		httpError := errors.Cause(resp.GetError(), (*errors.HTTPError)(nil))
		if httpError == nil {
			r.processUnhandledError(w, req, resp.GetError())
			return
		}

		r.processServiceError(w, req, resp)
		return
	}

	data, err := resp.GetData()
	if err != nil {
		r.processDataError(w, req, errors.WithMessage(err, "data serialization error"))
		return
	}

	// Send response headers.
	r.writeResponse(w, resp.GetStatus(), data, resp.GetHeaders())
}

// processDataError makes processing of the error occurred
// because of data serialization or data post processing.
func (r *Router) processDataError(w http.ResponseWriter, req *http.Request, err error) {

	r.logger.Error("%+v", err)

	r.writeResponse(w, http.StatusInternalServerError, nil, http.Header{})
}

// processServiceError makes processing of the error that occurred
// in case of validation or any other service error.
func (r *Router) processServiceError(
	w http.ResponseWriter,
	req *http.Request,
	response response.Provider,
) {

	if response.GetStatus() >= http.StatusInternalServerError {
		r.logger.Error("HTTP 5xx error happened: %v", response.GetError())
	} else {
		r.logger.Debug("%v", response.GetError())
	}

	data, err := response.GetData()
	if err != nil {
		r.processDataError(w, req, errors.WithMessage(err, "data serialization error"))
		return
	}

	r.writeResponse(
		w,
		response.GetStatus(),
		data,
		response.GetHeaders(),
	)
}

// processUnhandledError makes processing of the unhandled service error.
func (r *Router) processUnhandledError(w http.ResponseWriter, req *http.Request, err error) {

	r.logger.Error("internal unhandled error: %+v", err)

	r.writeResponse(w, http.StatusInternalServerError, nil, http.Header{})
}

// writeResponse writes data to the connection.
func (r *Router) writeResponse(
	writer http.ResponseWriter,
	status int,
	data []byte,
	headers http.Header,
) {

	// write request headers.
	for name, value := range headers {
		writer.Header().Set(name, strings.Join(value, " "))
	}

	// write HTTP status code.
	writer.WriteHeader(status)

	// write response data.
	if data != nil {
		_, err := writer.Write(data)
		if err != nil {
			r.logger.Error("data writing error: %+v", err)
		}
	}
}

// handle internal wrapper for original Gorilla MUX Handle method but with some internal filtering.
func (r *Router) handle(route string, h Handler, methods ...string) {

	r.router.Handle(route, http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			r.HandleResponse(w, req, h)
		},
	)).Methods(methods...)
}

// Get function add to router rule to handle specified route only by http Get method.
func (r *Router) Get(path string, h Handler) {
	r.handle(path, h, MethodGet)
}

// Post function add to router rule to handle specified route only by http Post method.
func (r *Router) Post(path string, h Handler) {
	r.handle(path, h, MethodPost)
}

// Put function add to router rule to handle specified route only by http Put method.
func (r *Router) Put(path string, h Handler) {
	r.handle(path, h, MethodPut)
}

// Patch function add to router rule to handle specified route only by http Patch method.
func (r *Router) Patch(path string, h Handler) {
	r.handle(path, h, MethodPatch)
}

// Options function add to router rule to handle specified route only by http Options method.
func (r *Router) Options(path string, h Handler) {
	r.handle(path, h, MethodOptions)
}

// Delete function add to router rule to handle specified route only by http Delete method.
func (r *Router) Delete(path string, h Handler) {
	r.handle(path, h, MethodDelete)
}

// Trace function add to router rule to handle specified route only by http Trace method.
func (r *Router) Trace(path string, h Handler) {
	r.handle(path, h, MethodTrace)
}

// Handle function allows you to accept all types of requests to specified route.
func (r *Router) Handle(path string, h HandlerFunc) {
	r.router.PathPrefix(path).Handler(h)
}
