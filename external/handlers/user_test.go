package handlers_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"gitlab.com/laboct2021/pnl-game/external/handlers"
	"gitlab.com/laboct2021/pnl-game/internal/models"
)

type HandlersUserTestSuite struct {
	HandlersTestSuite
}

func TestHandlersUserTestSuite(t *testing.T) {
	suite.Run(t, &HandlersUserTestSuite{})
}
func (s *HandlersUserTestSuite) TestUser() {

	s.Run("Create valid ", func() {
		reqBody := &handlers.RequestUser{
			Login:    "xruterx@gamil.com",
			Password: "1234",
		}
		response, err := MakeRequest(reqBody, http.MethodPost, "/signup", nil)
		s.Require().NoError(err)
		s.Require().Equal(response.StatusCode, http.StatusCreated)
	})

	s.Run("Create exist", func() {
		received := &models.ErrorMsg{}
		reqBody := &handlers.RequestUser{
			Login:    "xruterx@gamil.com",
			Password: "12341234",
		}
		expected := &models.ErrorMsg{
			Message: "user exist",
		}
		response, err := MakeRequest(reqBody, http.MethodPost, "/signup", &received)
		s.Require().NoError(err)
		s.Require().Equal(expected, received)
		s.Require().Equal(response.StatusCode, http.StatusBadRequest)
	})

}
