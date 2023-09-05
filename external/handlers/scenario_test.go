package handlers_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"gitlab.com/laboct2021/pnl-game/external/handlers"
	"gitlab.com/laboct2021/pnl-game/internal/models"
)

type HandlersScenarioTestSuite struct {
	HandlersTestSuite
}

func TestHandlersScenarioTestSuite(t *testing.T) {
	suite.Run(t, &HandlersScenarioTestSuite{})
}
func (s *HandlersScenarioTestSuite) TestScenario() {
	s.SingUp()
	token := s.Login()
	scenarioData := models.ScenarioData{
		Managers: []models.Manager{
			{
				ID:      1,
				Name:    "manager 1",
				UsersID: 1,
				Employees: []models.Employee{
					{
						ID:             1,
						Name:           "Employee 1",
						Seniority:      1,
						Salary:         1000,
						ProjectID:      -1,
						EmployeeStatus: "banch",
					},
				},
				Events: []models.Event{
					{
						ID:          1,
						EventTypeID: 1,
						Name:        "Card 1",
						Description: "Employee fired",
					},
				},
			},
		},
		Projects: []models.Project{
			{
				ID:   1,
				Name: "Project 1",
				Rates: []models.Rate{
					{
						ID:                   1,
						ProjectID:            1,
						Type:                 "A1",
						Seniority:            1,
						Rate:                 1,
						IllCompensation:      10,
						VacationCompensation: 10,
					},
				},
			},
		},
	}
	s.Run("Create scenario", func() {

		reqBody := &handlers.RequestScenario{
			OwnerID:    1,
			Name:       "scenario 1",
			StepsCount: 24,
			DeletedAt:  nil,
			Data:       scenarioData,
		}
		response, err := MakeAuthRequest(reqBody, "POST", "/admin/scenarios", nil, token.AccessToken)
		s.Require().NoError(err)
		s.Require().Equal(response.StatusCode, http.StatusCreated)
	})

	s.Run("return all scenario", func() {
		received := &[]models.Scenario{}
		expected := &[]models.Scenario{
			{
				ID:         1,
				OwnerID:    1,
				Name:       "scenario 1",
				StepsCount: 24,
				DeletedAt:  nil,
				Data:       scenarioData,
			},
		}
		response, err := MakeAuthRequest(nil, http.MethodGet, "/admin/scenarios", &received, token.AccessToken)
		s.Require().NoError(err)
		s.Require().Equal(expected, received)
		s.Require().Equal(response.StatusCode, http.StatusOK)
	})

	s.Run("delete scenario", func() {
		response, err := MakeAuthRequest(nil, http.MethodDelete, "/admin/scenarios/1", nil, token.AccessToken)
		s.Require().NoError(err)
		s.Require().Equal(response.StatusCode, http.StatusOK)
	})

	s.Run("get empty scenarios list", func() {
		received := &[]models.Scenario{}
		response, err := MakeAuthRequest(nil, http.MethodGet, "/admin/scenarios", &received, token.AccessToken)
		s.Require().NoError(err)
		s.Require().Equal(response.StatusCode, http.StatusOK)
		s.Require().Empty(received)
	})
}
