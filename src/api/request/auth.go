package request

// Signin represents Request object for POST /auth/signin route.
type Signin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Signup represents Request object for POST /auth/signup route.
type Signup struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
