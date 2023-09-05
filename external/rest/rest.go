package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"

	"gitlab.com/laboct2021/pnl-game/config"
	"gitlab.com/laboct2021/pnl-game/external/handlers"
	"gitlab.com/laboct2021/pnl-game/internal/repository"
	"gitlab.com/laboct2021/pnl-game/pkg/db"
	"gitlab.com/laboct2021/pnl-game/pkg/logwrapper"
	"gitlab.com/laboct2021/pnl-game/pkg/usecase"
)

const (
	LoggerFlushDuration = 2
	FatalExitCode       = 1
)

func Run(path string) error {
	logger := logwrapper.New(os.Getenv("DSN_URL_SENTRY"))
	defer logger.Flush(LoggerFlushDuration * time.Second)
	logger.CaptureMessage("application started")

	dbs := db.Setup(config.LoadConfig(path).DB)
	conn, err := db.New(&dbs)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		defer os.Exit(FatalExitCode)
	}
	defer conn.Close(context.Background())

	h := MakeRepo(conn, logger)
	router := MakeRouter(&h)

	log.Println("Starting API server on 8080")

	return http.ListenAndServe(":8080", router)
}

func MakeRepo(conn *pgx.Conn, logger logwrapper.LogWrapper) handlers.Handler {
	userRepo := repository.NewUserRepo(conn)
	scenarioRepo := repository.NewScenarioRepo(conn)
	svcUser := usecase.NewUserService(logger, userRepo)
	svcScenario := usecase.NewScenarioService(logger, scenarioRepo)
	return handlers.NewHandler(svcUser, svcScenario)
}

func MakeRouter(h *handlers.Handler) *mux.Router {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	subRouter.HandleFunc("/signin", handlers.WrapMidlaware(h.SingIn)).Methods(http.MethodPost)
	subRouter.HandleFunc("/refreshtoken", handlers.WrapMidlaware(h.RefreshToken)).Methods(http.MethodPost)
	subRouter.HandleFunc("/signup", handlers.WrapMidlaware(h.SingUp)).Methods(http.MethodPost)

	subRouter.HandleFunc("/admin/scenarios", handlers.WrapMidlaware(h.GetScenarios, h.JwtAuthentication)).Methods(http.MethodGet)
	subRouter.HandleFunc("/admin/scenarios", handlers.WrapMidlaware(h.PostScenario, h.JwtAuthentication)).Methods(http.MethodPost)
	subRouter.HandleFunc("/admin/scenarios/{id}", handlers.WrapMidlaware(h.DeleteScenarioById, h.JwtAuthentication)).Methods(http.MethodDelete)
	return router
}
