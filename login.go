package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// LoginResponse represents the answer to a login request
type LoginResponse struct {
	Session string `json:"session"`
	Url     string `json:"url"`
}

func (h *Handler) login() (*LoginResponse, error) {
	s, err := h.options.sessionHandler.New()
	if err != nil {
		return nil, err
	}

	id := s.Id()

	r := &LoginResponse{
		Session: id,
		Url:     h.oauth2Config.AuthCodeURL(id),
	}
	fmt.Printf("%s\n", r.Url)
	return r, nil
}

func (h *Handler) ginLogin(ctx *gin.Context) {
	r, err := h.login()
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, r)
}
