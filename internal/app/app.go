package app

import (
	"github.com/dshepett/task-manager/internal/config"
	"github.com/dshepett/task-manager/internal/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type App struct {
	router  *mux.Router
	log     *logrus.Logger
	storage *storage.Storage
	config  *config.Config
}

func New(config *config.Config) *App {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	app := &App{router: mux.NewRouter(), log: logger, storage: storage.New(config)}
	app.addRoutes()
	app.log.Infoln("app created")
	return app
}

func (a *App) Run() {
	a.storage.Open()
	a.log.Infoln("running server...")
	a.log.Fatal(http.ListenAndServe("localhost:8080", a.router))
}
