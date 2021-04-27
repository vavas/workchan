package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vavas/workchan/app/config"
	"github.com/vavas/workchan/app/server/endpoints"
	"github.com/vavas/workchan/app/server/middleware"
	"github.com/vavas/workchan/pkg/db"
	"net/http"
	"time"
)

type RouterDeps struct {
	Logger     logrus.FieldLogger
	APIAuthMap gin.Accounts
	DbConns    *db.DB
}

func ConfigureRouter(deps *RouterDeps) (*gin.Engine, error) {
	g := gin.New()
	g.Use(middleware.InjectDBConnections(deps.DbConns))
	g.GET("/healthcheck", endpoints.HealthCheck)
	g.GET("/dbcheck", endpoints.DBCheck)
	return g, nil
}

func Start(router http.Handler, conf *config.Server) error {
	if conf.TCPPort == "" {
		return errors.New("env PORT is undefined, it must be valid port to start the server")
	}
	s := http.Server{
		Addr:         ":" + conf.TCPPort,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s.ListenAndServe()
}
