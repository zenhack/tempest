package browsermain

// Model for the email login form
type LoginForm struct {
	TokenSent  bool   // Whether we've already sent a token:
	EmailInput string // The email the user has entered
	TokenInput string // The token the user has entered
}
