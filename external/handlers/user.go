package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"gitlab.com/laboct2021/pnl-game/external/lib"
	"gitlab.com/laboct2021/pnl-game/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type RequestUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h Handler) SingUp(w http.ResponseWriter, r *http.Request) {
	requestUser := RequestUser{}
	err := json.NewDecoder(r.Body).Decode(&requestUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.ReturnError(w, errors.New("bad request"))
		return
	}

	_, err = h.svcUser.FindByLogin(r.Context(), requestUser.Login)
	if err == nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.ReturnError(w, errors.New("user exist"))
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestUser.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.ReturnError(w, err)
		return
	}

	user := models.User{ID: 0,
		Login:        requestUser.Login,
		PasswordHash: string(hashedPassword)}

	err = h.svcUser.Create(r.Context(), &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.ReturnError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
