package health

// ServiceInfoResponse Service Info response.
type ServiceInfoResponse struct {
	Dependencies map[string]*Data `json:"dependencies,omitempty"`
}

// NewServiceInfoResponse creates new instance of Service Info response.
func NewServiceInfoResponse(dependencies map[string]*Data) *ServiceInfoResponse {
	return &ServiceInfoResponse{
		Dependencies: dependencies,
	}
}
