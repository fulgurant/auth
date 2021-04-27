package auth

import (
	"context"

	"github.com/gin-gonic/gin"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type Handler struct {
	options      *Options
	oidcProvider *oidc.Provider
	oauth2Config *oauth2.Config
}

func New(options *Options) (*Handler, error) {
	oidcProvider, err := oidc.NewProvider(context.Background(), options.config.ProviderUrl)
	if err != nil {
		return nil, err
	}

	oauth2Config := &oauth2.Config{
		ClientID:     options.config.ClientID,
		ClientSecret: options.config.ClientSecret,
		RedirectURL:  options.config.LoginCallbackUrl,
		Endpoint:     oidcProvider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Handler{
		options:      options,
		oidcProvider: oidcProvider,
		oauth2Config: oauth2Config,
	}, nil
}

func (h *Handler) RegisterEndpoint(r gin.IRouter) {
	r.GET("/login", h.ginLogin)
	r.GET("/callback", h.ginCallback)

	if h.options.logger != nil {
		h.options.logger.Info("endpoints registered")
	}
}
