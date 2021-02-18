package transport

import "github.com/dsuhinin/suhinin-backend-1/src/app/controllers"

// Transport is provide abstraction on transport layer.
type Transport struct {
	serviceController controllers.Provider
}

// NewTransport return Transport instance.
func NewTransport(serviceController controllers.Provider) *Transport {
	return &Transport{
		serviceController: serviceController,
	}
}
