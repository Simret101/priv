package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

// InitOauth initializes the OAuth2 configuration by loading environment variables and setting up the OAuth2 config.
func InitOauth() (*oauth2.Config, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Could not load .env file, falling back to system environment variables. Error: %v", err)
	}

	// Load required environment variables
	clientId := os.Getenv("CLIENTID")
	clientSecret := os.Getenv("CLIENTSECRET")
	redirectUrl := os.Getenv("REDIRECTURL")

	// Validate environment variables
	if clientId == "" {
		return nil, errors.New("missing CLIENTID environment variable")
	}
	if clientSecret == "" {
		return nil, errors.New("missing CLIENTSECRET environment variable")
	}
	if redirectUrl == "" {
		return nil, errors.New("missing REDIRECTURL environment variable")
	}

	// Set up the OAuth2 configuration
	googleOauthConfig := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  redirectUrl,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}

	return googleOauthConfig, nil
}
