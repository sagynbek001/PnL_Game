package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/suite"
	"gitlab.com/laboct2021/pnl-game/config"
	"gitlab.com/laboct2021/pnl-game/external/handlers"
	"gitlab.com/laboct2021/pnl-game/external/rest"
	"gitlab.com/laboct2021/pnl-game/internal/models"
	"gitlab.com/laboct2021/pnl-game/pkg/db"
)

const (
	hosturl    = "localhost"
	port       = 8080
	pathPrefix = "api/v1"
)

type HandlersTestSuite struct {
	suite.Suite
	m *migrate.Migrate
}

func (s *HandlersTestSuite) SetupSuite() {
	err := s.Prepare("../../config/config.yaml")
	s.Require().NoError(err)
}

func (s *HandlersTestSuite) TearDownSuite() {
	// Run migrate down
	err := s.m.Down()
	s.Require().NoError(err)
}

func (s *HandlersTestSuite) Prepare(path string) error {
	config := config.LoadConfig(path)
	_, err := db.New((*db.Setup)(&config.DB))
	if err != nil {
		return err
	}
	//migrate up
	absPath, err := filepath.Abs("../../db/migration")
	if err != nil {
		return err
	}
	sourcePtr := "file://" + absPath
	dbs := db.Setup(config.DB)
	databasePtr := dbs.String()
	s.m, err = migrate.New(sourcePtr, databasePtr)
	if err != nil {
		return err
	}
	err = s.m.Up()
	if err != nil {
		return err
	}
	go func() error {
		err := rest.Run(path)
		return err
	}()
	if err != nil {
		return err
	}
	time.Sleep(3 * time.Second)
	return nil
}

func (s *HandlersTestSuite) SingUp() {
	reqBodyCreateUser := handlers.RequestUser{
		Login:    "admin@gamil.com",
		Password: "1234",
	}
	response, err := MakeRequest(reqBodyCreateUser, http.MethodPost, "/signup", nil)
	s.Require().NoError(err)
	s.Require().Equal(response.StatusCode, http.StatusCreated)

}

func (s *HandlersTestSuite) Login() *models.Token {
	token := &models.Token{}
	reqBodyLoginUser := handlers.Login{
		Login:    "admin@gamil.com",
		Password: "1234",
	}
	response, err := MakeRequest(reqBodyLoginUser, http.MethodPost, "/signin", token)
	s.Require().NoError(err)
	s.Require().Equal(response.StatusCode, http.StatusOK)
	return token
}

func createRequest(reqBody interface{}, method string, uri string) (*http.Request, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(reqBody)
	if err != nil {
		return nil, err
	}
	uri = path.Join(pathPrefix, uri)
	strUri := fmt.Sprintf("http://%v:%v/%v", hosturl, port, uri)
	req, err := http.NewRequest(method, strUri, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func getResponse(req *http.Request, resp interface{}) (*http.Response, error) {
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer response.Body.Close()
	if resp != nil {
		err = json.NewDecoder(response.Body).Decode(&resp)
		if err != nil {
			return response, err
		}
	}
	return response, nil
}

func MakeRequest(reqBody interface{}, method string, uri string, resp interface{}) (*http.Response, error) {
	req, err := createRequest(reqBody, method, uri)
	if err != nil {
		return nil, err
	}
	response, err := getResponse(req, resp)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func MakeAuthRequest(reqBody interface{}, method string, uri string, resp interface{}, token string) (*http.Response, error) {
	req, err := createRequest(reqBody, method, uri)
	if err != nil {
		return nil, err
	}
	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}
	response, err := getResponse(req, resp)
	if err != nil {
		return nil, err
	}
	return response, nil
}
