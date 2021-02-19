package response

// Members represents response of `GET /members` route.
type Members struct {
	Text string `json:"text"`
}

// NewMembers creates new Response instance for `GET /members` endpoint.
func NewMembers(text string) *Members {
	return &Members{
		Text: text,
	}
}
