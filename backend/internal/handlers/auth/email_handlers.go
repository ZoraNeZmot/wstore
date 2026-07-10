package authhandlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
	"wstore/internal/handlers/help/response"
	"wstore/internal/storage/dberrors"
	"wstore/internal/storage/dbusers"
	"wstore/pkg/argon2"
	"wstore/pkg/token"
)

func (a *AuthHandlers) EmailRegisterRoutes(mux *http.ServeMux) error {
	mux.HandleFunc("POST /auth/email/signup", a.EmailSignup)
	mux.HandleFunc("POST /auth/email/login", a.EmailLogin)
	return nil
}

// @Summary		Email signup
// @Description	Signup with email and password
// @Tags			auth
// @Accept			application/json
// @Produce		application/json
// @Param			user	body		EmailRequest	true	"User"
// @Success		201		{object}	response.Response{data=response.MessageResponse,error=interface{}}
// @Failure		400		{object}	map[string]string
// @Failure		500		{object}	map[string]string
// @Router			/auth/email/signup [post]
func (a *AuthHandlers) EmailSignup(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Failed to read body", err.Error())
		a.logger.Debugf("EmailSignup: Failed to read body: %v", err)
		return
	}
	defer r.Body.Close()

	var user EmailRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Failed to parse json body", err.Error())
		a.logger.Debugf("EmailSignup: Failed to unmarshal body: %v", err)
		return
	}

	argon := argon2.New(nil, nil)
	hashPassword, err := argon.Hash(user.Password)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to hash password", err.Error())
		a.logger.Debugf("EmailSignup: Failed to hash password: %v", err)
		return
	}

	user.Password = string(hashPassword)
	user.Email = strings.ToLower(user.Email)
	user.Name = strings.ToUpper(user.Name)

	userData := dbusers.NewUser(user.Email, user.Password, user.Name)
	exists := a.db.IsUserExists(r.Context(), userData)
	if exists {
		response.Error(w, http.StatusBadRequest, dberrors.UserAlreadyExists.Error(), err.Error())
		a.logger.Debugf("EmailSignup: User already exists: %v", err)
		return
	}

	_, err = a.db.CreateUser(
		r.Context(),
		userData,
	)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to add user. Try again later", err.Error())
		a.logger.Debugf("EmailSignup: Failed to add user: %v", err)
		return
	}

	response.Success(w, http.StatusCreated, response.MessageResponse{Message: "User created successfully"})

	// TODO: send email verification
}

// @Summary		Email login
// @Description	Login with email and password
// @Tags			auth
// @Accept			application/json
// @Produce		application/json
// @Param			user	body		EmailRequest	true	"User"
// @Success		200		{object}	map[string]string
// @Failure		400		{object}	map[string]string
// @Failure		500		{object}	map[string]string
// @Router			/auth/email/login [post]
func (a *AuthHandlers) EmailLogin(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Failed to read body", err.Error())
		a.logger.Debugf("EmailLogin: Failed to read body: %v", err)
		return
	}
	defer r.Body.Close()
	var user EmailRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Failed to parse json body", err.Error())
		a.logger.Debugf("EmailLogin: Failed to unmarshal body: %v", err)
		return
	}

	user.Email = strings.ToLower(user.Email)

	userData, err := a.db.GetUserByEmail(r.Context(), dbusers.NewUserWithEmail(user.Email))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(w, http.StatusUnauthorized, "User not found. Check your email or password", err.Error())
			a.logger.Debugf("EmailLogin: User not found: %v", err)
			return
		}

		response.Error(w, http.StatusInternalServerError, "Failed to get user. Try again later", err.Error())
		a.logger.Errorf("EmailLogin: Failed to get user: %v", err)
		return
	}

	isValid, err := argon2.Verify(user.Password, userData.Password)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to verify password. Try again later", err.Error())
		a.logger.Debugf("EmailLogin: Failed to verify password: %v", err)
		return
	}
	if !isValid {
		response.Error(w, http.StatusUnauthorized, "Invalid email or password", nil)
		a.logger.Debugf("EmailLogin: Invalid email or password")
		return
	}

	tokenString, err := token.GenerateToken(token.TokenClaims{
		UserID: userData.ID,
		Name:   userData.Name,
		Exp:    time.Now().Add(time.Hour * 24),
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to generate token. Try again later", err.Error())
		a.logger.Debugf("EmailLogin: Failed to generate token: %v", err)
		return
	}
	response.Success(w, http.StatusOK, map[string]string{"token": tokenString, "message": "Login successful"})
}

func (a *AuthHandlers) EmailLogout(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement should I do that here?
}
