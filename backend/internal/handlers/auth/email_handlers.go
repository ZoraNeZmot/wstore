package authhandlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	handlershelp "wstore/internal/handlers/help"
	"wstore/internal/handlers/help/response"
	"wstore/internal/storage/dbtokens"
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
// @Success		201		{object}	response.MessageDTO
// @Failure		400		{object}	response.ErrorDTO
// @Failure		409		{object}	response.ErrorDTO
// @Failure		500		{object}	response.ErrorDTO
// @Router			/auth/email/signup [post]
func (a *AuthHandlers) EmailSignup(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Bad Request", err.Error())
		a.logger.Debugf("EmailSignup: Failed to read body: %v", err)
		return
	}
	defer r.Body.Close()

	var user EmailRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Bad Request", err.Error())
		a.logger.Debugf("EmailSignup: Failed to unmarshal body: %v", err)
		return
	}

	argon := argon2.New(nil, nil)
	hashPassword, err := argon.Hash(user.Password)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Server error", err.Error())
		a.logger.Debugf("EmailSignup: Failed to hash password: %v", err)
		return
	}

	user.Password = string(hashPassword)
	user.Name = handlershelp.FormatPersonName(user.Name)

	userData := dbusers.NewUser(user.Email, user.Password, user.Name)
	exists := a.db.IsUserExists(r.Context(), userData)
	if exists {
		response.Error(w, http.StatusConflict, "Conflict", "User already exists")
		a.logger.Debugf("EmailSignup: User already exists")
		return
	}

	// TODO handle id returning
	id, err := a.db.CreateUser(
		r.Context(),
		userData,
	)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Server error", "Cannot create user")
		a.logger.Debugf("EmailSignup: Failed to add user: %v", err)
		return
	}
	userData.ID = id

	imageUrl := fmt.Sprintf("https://ui-avatars.com/api/?name=%s&size=256", userData.Name)
	image, err := a.getImage(imageUrl)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Server error", "Cannot get image")
		a.logger.Debugf("EmailSignup: Failed to get image: %v", err)
		return
	}

	imagePath := fmt.Sprintf("./images/users/%s.jpg", userData.ID)
	os.MkdirAll(filepath.Dir(imagePath), 0755)
	err = os.WriteFile(imagePath, image, 0644)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Server error", "Cannot save image")
		a.logger.Debugf("EmailSignup: Failed to save image: %v", err)
		return
	}
	err = a.db.CreateUserProfileImage(r.Context(), dbusers.NewUserWithImagePath(userData.ID, imagePath))
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Server error", "Cannot save user profile image path")
		a.logger.Debugf("EmailSignup: Failed to save user profile image path: %v", err)
		return
	}

	response.Success(w, http.StatusCreated, response.MessageResponse{Message: "User created successfully"})

	// TODO: send email verification
}

func (a *AuthHandlers) getImageUrl(name string) string {
	return fmt.Sprintf("https://ui-avatars.com/api/?name=%s&size=256", name)
}

func (a *AuthHandlers) getImage(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return io.ReadAll(response.Body)
}

// @Summary		Email login
// @Description	Login with email and password
// @Tags			auth
// @Accept			application/json
// @Produce		application/json
// @Param			user	body		EmailRequest	true	"User"
// @Success		200		{object}	response.TokenDTO
// @Failure		400		{object}	response.ErrorDTO
// @Failure		401		{object}	response.ErrorDTO
// @Failure		500		{object}	response.ErrorDTO
// @Router			/auth/email/login [post]
func (a *AuthHandlers) EmailLogin(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Bad Request", err.Error())
		a.logger.Debugf("EmailLogin: Failed to read body: %v", err)
		return
	}
	defer r.Body.Close()
	var user EmailRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Bad Request", err.Error())
		a.logger.Debugf("EmailLogin: Failed to unmarshal body: %v", err)
		return
	}

	user.Email = strings.ToLower(user.Email)

	userData, err := a.db.GetUserByEmail(r.Context(), dbusers.NewUserWithEmail(user.Email))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(w, http.StatusUnauthorized, "Unauthorized", "Invalid email or password")
			a.logger.Debugf("EmailLogin: User not found: %v", err)
			return
		}

		response.Error(w, http.StatusInternalServerError, "Server error", "Cannot get user")
		a.logger.Errorf("EmailLogin: Cannot get user: %v", err)
		return
	}

	isValid, err := argon2.Verify(user.Password, userData.Password)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Server error", "Cannot verify password")
		a.logger.Errorf("EmailLogin: Cannot verify password: %v", err)
		return
	}
	if !isValid {
		response.Error(w, http.StatusUnauthorized, "Unauthorized", "Invalid email or password")
		a.logger.Debugf("EmailLogin: Invalid email or password")
		return
	}

	refreshToken, err := token.GenerateToken(token.NewTokenClaims(
		userData.ID,
		userData.Email,
		userData.Name,
		time.Now().Add(time.Hour*30*24)),
	)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Server error", "Cannot generate refresh token")
		a.logger.Debugf("EmailLogin: Failed to generate refresh token: %v", err)
		return
	}
	tokenData := dbtokens.NewToken(refreshToken, userData.ID, time.Now().Add(time.Hour*30*24))
	err = a.db.CreateToken(r.Context(), tokenData)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Server error", "Cannot create access token")
		a.logger.Debugf("EmailLogin: Failed to create access token: %v", err)
		return
	}

	accessToken, err := token.GenerateToken(token.NewTokenClaims(
		userData.ID,
		userData.Email,
		userData.Name,
		time.Now().Add(time.Minute*15),
	))
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Server error", "Cannot generate access token")
		a.logger.Debugf("EmailLogin: Failed to generate token: %v", err)
		return
	}

	w.Header().Set("Set-Cookie", fmt.Sprintf("RefreshToken=%s; HttpOnly; Secure; SameSite=Strict; Max-Age=%d; Path=/api/auth", refreshToken, time.Hour*30*24))
	response.Success(w, http.StatusOK, response.TokenResponse{AccessToken: accessToken})
}
