package lib

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.com/laboct2021/pnl-game/internal/models"
)

func IDFromVars(r *http.Request) (uint64, error) {
	idString := mux.Vars(r)["id"]

	i, err := strconv.Atoi(idString)
	if err != nil {
		return 0, err
	}

	return uint64(i), nil
}
func ReturnError(w http.ResponseWriter, err error) {
	errMsg := models.ErrorMsg{
		Message: err.Error(),
	}
	ReturnJSON(w, errMsg)
}

func ReturnErrorBadRequest(w http.ResponseWriter, text string) {
	w.WriteHeader(http.StatusBadRequest)
	errMsg := models.ErrorMsg{
		Message: text,
	}
	ReturnJSON(w, errMsg)
}
func ReturnErrorUnauthorized(w http.ResponseWriter, text string) {
	w.WriteHeader(http.StatusUnauthorized)
	errMsg := models.ErrorMsg{
		Message: text,
	}
	ReturnJSON(w, errMsg)
}
func ReturnErrorForbidden(w http.ResponseWriter, text string) {
	w.WriteHeader(http.StatusForbidden)
	errMsg := models.ErrorMsg{
		Message: text,
	}
	ReturnJSON(w, errMsg)
}

func ReturnJSON(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")

	if _, err := w.Write(b); err != nil {
		log.Fatal(err)
	}
}
func IsValidLogin(login string) bool {
	secure := true
	tests := []string{".{7,}", "[a-z]", "[A-Z]", "[0-9]", "[^\\d\\w]"}
	for _, test := range tests {
		t, _ := regexp.MatchString(test, login)
		if !t {
			secure = false
			break
		}
	}
	return secure
}
