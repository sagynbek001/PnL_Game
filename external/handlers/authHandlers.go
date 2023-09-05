package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"gitlab.com/laboct2021/pnl-game/external/lib"
	"gitlab.com/laboct2021/pnl-game/external/service"
	"golang.org/x/crypto/bcrypt"

	"gitlab.com/laboct2021/pnl-game/internal/models"
)

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h Handler) SingIn(w http.ResponseWriter, r *http.Request) {
	login := Login{}
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.ReturnError(w, fmt.Errorf("invalid request: %w", err))
		return
	}

	user, err := h.svcUser.FindByLogin(r.Context(), login.Login)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		lib.ReturnError(w, errors.New("invalid login"))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(login.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		lib.ReturnError(w, errors.New("invalid login"))
		return
	}
	//
	jwt := service.Jwt{}
	token, err := jwt.CreateToken(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.ReturnError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	lib.ReturnJSON(w, token)
}

func (h Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {

	token := models.Token{}
	err := json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.ReturnError(w, fmt.Errorf("bad request: %w", err))
		return
	}

	jwt := service.Jwt{}
	user, err := jwt.ValidateRefreshToken(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.ReturnError(w, errors.New("invalid token"))
		return
	}

	token, err = jwt.CreateToken(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.ReturnError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	lib.ReturnJSON(w, token)
}

func (h Handler) JwtAuthentication(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		token := strings.Split(bearerToken, " ")
		if len(token) != 2 {
			w.WriteHeader(http.StatusInternalServerError)
			lib.ReturnError(w, errors.New("malformed auth token"))
			return
		}
		jwt := service.Jwt{}
		user, err := jwt.ValidateToken(token[1])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			lib.ReturnError(w, errors.New("invalid token"))
			return
		}
		if _, err := h.svcUser.GetByID(r.Context(), user.ID); err != nil {
			w.WriteHeader(http.StatusForbidden)
			lib.ReturnError(w, err)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)
		next(w, r)
	})
}
