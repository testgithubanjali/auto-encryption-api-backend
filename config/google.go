package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOauthConfig = &oauth2.Config{
	ClientID:     "YOUR_NEW_CLIENT_ID",
	ClientSecret: "YOUR_NEW_CLIENT_SECRET",
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}
