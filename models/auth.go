package models

type AuthOptions struct {
	UID           string
	Email         string
	Password      string
	DisplayName   string
	PhotoURL      string
	PhoneNumber   string
	Disabled      bool
	EmailVerified bool
	CustomClaims  map[string]interface{}
}
