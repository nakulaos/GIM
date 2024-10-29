package xoauth2

import (
	"GIM/pkg/conf"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewGoogleOauthConfig(cfg *conf.GoogleOAuth2) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.ClientId,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectUrl,
		Scopes:       cfg.Scopes,
		Endpoint:     google.Endpoint,
	}
}
