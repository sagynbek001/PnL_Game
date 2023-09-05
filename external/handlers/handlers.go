package handlers

import (
	"net/http"

	"gitlab.com/laboct2021/pnl-game/pkg/usecase"
)

type Handler struct {
	svcUser     usecase.UserService
	svcScenario usecase.ScenarioService
}

func NewHandler(svcUser usecase.UserService,
	svcScenario usecase.ScenarioService,
) Handler {
	return Handler{
		svcUser:     svcUser,
		svcScenario: svcScenario,
	}
}

func WrapMidlaware(h http.HandlerFunc, midlles ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, mid := range midlles {
		h = mid(h)
	}
	return h
}
