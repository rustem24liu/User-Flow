package api

import (
	"net/http"
	"time"
	"user-flow/configs"
	"user-flow/pkg/postgres"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	pgdb   *gorm.DB
}

func NewServer() *Server {
	router := gin.Default()
	configs.InitConfig()
	pgdb := postgres.NewClient()

	app := &Server{
		router: router,
		pgdb:   pgdb,
	}

	return app
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:         ":8000",
		Handler:      s.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := s.initRoutes()
	if err != nil {
		return err
	}

	return srv.ListenAndServe()
}
