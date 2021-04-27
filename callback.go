package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Errors
var (
	ErrNoOAuth2State         = errors.New("no oauth2 state provided")
	ErrNoOAuth2Code          = errors.New("no oauth2 code provided")
	ErrCouldNotRetrieveToken = errors.New("could not retrieve token")
	ErrSessionNotFound       = errors.New("session not found")
	ErrNoJWTReturned         = errors.New("no jwt returned")
)

// Session fields
var (
	tokenField = "token"
	jwtPrefix  = "jwt_"
)

// Rather use a dedicated jwt library that can validate the jwt
// signatures. Perhaps github.com/dgrijalva/jwt-go
func parseJwt(jwt string) (map[string]interface{}, error) {
	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		return nil, errors.New("Bad jwt format")
	}

	middlePart64 := parts[1]
	middlePart, err := base64.RawStdEncoding.DecodeString(middlePart64)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}
	err = json.Unmarshal(middlePart, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (h *Handler) ginCallback(ctx *gin.Context) {
	sessionId, found := ctx.GetQuery("state")
	if !found {
		ctx.AbortWithError(http.StatusBadRequest, ErrNoOAuth2State)
		return
	}
	code, found := ctx.GetQuery("code")
	if !found {
		ctx.AbortWithError(http.StatusBadRequest, ErrNoOAuth2Code)
		return
	}
	token, err := h.oauth2Config.Exchange(ctx.Request.Context(), code)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, ErrCouldNotRetrieveToken)
		return
	}
	session, err := h.options.sessionHandler.Get(sessionId)
	if err != nil {
		ctx.AbortWithError(http.StatusForbidden, err)
		return
	}
	session.Set(tokenField, token)
	jwtString, ok := token.Extra("id_token").(string)
	if !ok {
		ctx.AbortWithError(http.StatusInternalServerError, ErrNoJWTReturned)
		return
	}
	jwt, err := parseJwt(jwtString)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Save the jwt elements in the sesison as "jwt_<key> = <value>"
	for k, v := range jwt {
		session.Set(jwtPrefix+k, v)
	}
	ctx.JSON(200, "OK")
}
