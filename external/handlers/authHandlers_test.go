package handlers_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"gitlab.com/laboct2021/pnl-game/external/handlers"
	"gitlab.com/laboct2021/pnl-game/internal/models"
)

type HandlersAuthTestSuite struct {
	HandlersTestSuite
}

func TestHandlersAuthTestSuite(t *testing.T) {
	suite.Run(t, &HandlersAuthTestSuite{})
}

func (s *HandlersAuthTestSuite) TestLogin() {
	s.SingUp()
	s.Run("Login with valid request", func() {
		received := &models.Token{}
		reqBody := &handlers.Login{
			Login:    "admin@gamil.com",
			Password: "1234",
		}
		expected := &models.Token{}
		response, err := MakeRequest(reqBody, http.MethodPost, "/signin", &received)
		s.Require().NoError(err)
		s.Require().IsType(expected, received)
		s.Require().Equal(response.StatusCode, http.StatusOK)
	})

	s.Run("Login with invalid password", func() {
		received := &models.ErrorMsg{}
		reqBody := &handlers.Login{
			Login:    "admin@gamil.com",
			Password: "1234sd",
		}
		expected := &models.ErrorMsg{
			Message: "invalid login",
		}
		response, err := MakeRequest(reqBody, http.MethodPost, "/signin", &received)
		s.Require().NoError(err)
		s.Require().Equal(expected, received)
		s.Require().Equal(response.StatusCode, http.StatusUnauthorized)
	})

	s.Run("Login with invalid email", func() {
		received := &models.ErrorMsg{}
		reqBody := &handlers.Login{
			Login:    "xruterxgamil.com",
			Password: "12341234",
		}
		expected := &models.ErrorMsg{
			Message: "invalid login",
		}
		response, err := MakeRequest(reqBody, http.MethodPost, "/signin", &received)
		s.Require().NoError(err)
		s.Require().Equal(expected, received)
		s.Require().Equal(response.StatusCode, http.StatusUnauthorized)
	})
}
