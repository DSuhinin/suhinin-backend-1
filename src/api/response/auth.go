package response

// Signin represents response of `POST /auth/signin` route.
type Signin struct {
	Token string `json:"token"`
}

// NewSignin creates new Response instance for `POST /auth/signin` endpoint.
func NewSignin(token string) *Signin {
	return &Signin{
		Token: token,
	}
}
