package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gitlab.com/laboct2021/pnl-game/external/lib"
	"gitlab.com/laboct2021/pnl-game/internal/models"
)

func (h Handler) GetScenarios(w http.ResponseWriter, r *http.Request) {
	scenarios, err := h.svcScenario.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.ReturnError(w, fmt.Errorf("get all scenarios: %w", err))
		return
	}
	w.WriteHeader(http.StatusOK)
	lib.ReturnJSON(w, scenarios)
}

type RequestScenario struct {
	OwnerID    int64               `json:"owner_id"`
	Name       string              `json:"name"`
	StepsCount int64               `json:"steps_count"`
	DeletedAt  *time.Time          `json:"deleted_at"`
	Data       models.ScenarioData `json:"data"`
}

func (h Handler) PostScenario(w http.ResponseWriter, r *http.Request) {

	requestScenario := RequestScenario{}
	err := json.NewDecoder(r.Body).Decode(&requestScenario)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.ReturnError(w, fmt.Errorf("bad request: %w", err))
		return
	}

	scenario := models.Scenario{
		OwnerID:    requestScenario.OwnerID,
		Name:       requestScenario.Name,
		StepsCount: requestScenario.StepsCount,
		DeletedAt:  requestScenario.DeletedAt,
		Data:       requestScenario.Data,
	}

	err = h.svcScenario.Create(r.Context(), &scenario)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.ReturnError(w, fmt.Errorf("create scenario: %w", err))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) DeleteScenarioById(w http.ResponseWriter, r *http.Request) {
	id, err := lib.IDFromVars(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.ReturnError(w, fmt.Errorf("get id from vars: %w", err))
		return
	}
	err = h.svcScenario.Delete(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.ReturnError(w, fmt.Errorf("delete scenario: %w", err))
		return
	}
	w.WriteHeader(http.StatusOK)
}
