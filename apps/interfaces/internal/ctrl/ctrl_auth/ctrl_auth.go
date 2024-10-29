package ctrl_auth

import (
	"GIM/apps/interfaces/internal/config"
	"GIM/apps/interfaces/internal/service/svc_auth"
	"GIM/pkg/common/xoauth2"
	"golang.org/x/oauth2"
)

type AuthCtrl struct {
	authService       svc_auth.AuthService
	googleOauthConfig *oauth2.Config
	githubOauthConfig *oauth2.Config
}

func NewAuthCtrl(authService svc_auth.AuthService) *AuthCtrl {
	srv := &AuthCtrl{authService: authService}
	conf := config.GetConfig()
	srv.googleOauthConfig = xoauth2.NewGoogleOauthConfig(conf.Google)
	srv.githubOauthConfig = xoauth2.NewGithubOauthConfig(conf.Github)
	return srv
}
