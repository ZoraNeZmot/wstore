//go:build ignore

package authhandlers

import (
	"database/sql"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
	handlershelp "wstore/internal/handlers/help"
	"wstore/internal/storage/dbusers"
	"wstore/pkg/token"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
)

func (a *AuthHandlers) GoogleRegisterRoutes(mux *http.ServeMux) error {
	mux.HandleFunc("/auth/google/login", a.GoogleLogin)
	mux.HandleFunc("/auth/google/callback", a.GoogleCallback)
	return nil
}

func (a *AuthHandlers) redirectToFrontend(w http.ResponseWriter, r *http.Request, params url.Values) {
	redirectURL, err := url.Parse(a.Config.FrontendURL)
	if err != nil {
		handlershelp.JSONError(w, http.StatusInternalServerError, "Invalid frontend URL")
		a.logger.WithError(err).Error("Failed to parse frontend URL")
		return
	}
	redirectURL.Path = "/auth/callback"
	redirectURL.RawQuery = params.Encode()
	http.Redirect(w, r, redirectURL.String(), http.StatusTemporaryRedirect)
}

func (a *AuthHandlers) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	state := uuid.New().String()
	http.SetCookie(w, &http.Cookie{
		Name:     "state",
		Value:    state,
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	authURL := a.Config.GetGoogleOAuthConfig().AuthCodeURL(state, oauth2.AccessTypeOffline)

	a.logger.WithField("Handler Name", "GoogleLogin").Debugf("Google state: %s", state)
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

func (a *AuthHandlers) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	log := a.logger.WithField("Handler Name", "GoogleCallback")
	state := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")
	log.Debugf("GoogleCallback state: %s, code: %s", state, code)

	fail := func(message string) {
		params := url.Values{}
		params.Set("error", message)
		a.redirectToFrontend(w, r, params)
	}

	if state == "" {
		log.Error("State not found")
		fail("State not found")
		return
	}
	cookie, err := r.Cookie("state")
	if err != nil {
		log.Error("State not found in COOKIE")
		fail("State not found")
		return
	}
	if cookie.Value != state {
		log.Error("State mismatch")
		fail("State mismatch")
		return
	}

	oauthToken, err := a.Config.GetGoogleOAuthConfig().Exchange(
		r.Context(),
		code,
	)
	if err != nil {
		log.WithError(err).Error("Failed to exchange code")
		fail("Failed to exchange code")
		return
	}

	idTokenRaw, ok := oauthToken.Extra("id_token").(string)
	if !ok || idTokenRaw == "" {
		log.Error("Missing id_token in OAuth response")
		fail("Missing id_token")
		return
	}

	payload, err := idtoken.Validate(r.Context(), idTokenRaw, a.Config.GoogleOAuthConfig.ClientID)
	if err != nil {
		log.WithError(err).Error("Failed to validate id_token")
		fail("Invalid id_token")
		return
	}

	email, _ := payload.Claims["email"].(string)
	name, _ := payload.Claims["name"].(string)
	emailVerified, _ := payload.Claims["email_verified"].(bool)
	if email == "" || !emailVerified {
		log.Error("Email missing or not verified in id_token")
		fail("Email not verified")
		return
	}

	email = strings.ToLower(email)
	if name == "" {
		name = email
	}
	name = strings.ToUpper(name)

	userData, err := a.db.GetUserByEmail(r.Context(), dbusers.NewUserWithEmail(email))
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.WithError(err).Error("Failed to get user")
			fail("Failed to get user. Try again later")
			return
		}

		id, createErr := a.db.CreateOAuthUser(r.Context(), dbusers.NewOAuthUser(email, name))
		if createErr != nil {
			log.WithError(createErr).Error("Failed to create OAuth user")
			fail("Failed to create user. Try again later")
			return
		}
		userData = &dbusers.User{ID: id, Email: email, Name: name}
	}

	tokenString, err := token.GenerateToken(token.TokenClaims{
		UserID: userData.ID,
		Email:  userData.Email,
		Name:   userData.Name,
		Exp:    time.Now().Add(time.Hour * 24),
	})
	if err != nil {
		log.WithError(err).Error("Failed to generate token")
		fail("Failed to generate token. Try again later")
		return
	}

	log.Debug("SUCCESS exchange code for token")
	params := url.Values{}
	params.Set("token", tokenString)
	a.redirectToFrontend(w, r, params)
}
