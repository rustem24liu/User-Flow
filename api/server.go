package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	pgdb *gorm.DB
}

var handler *gin.Engine

//func NewServer(ctx context.Context) error {
//	s := &Server{}
//
//	if err := s
//}

func (s *Server) initDeps(ctx context.Context) {

}

//
//func (s *Server) initServer(ctx context.Context) error {
//	httpCfg := configs.Config.App.Url
//	httServer := http.
//}
