package api

import (
	"user-flow/configs"
	userhandler "user-flow/internal/app/domain/handlers/user"
	userrepository "user-flow/internal/app/domain/repositories/user"
	userservice "user-flow/internal/app/domain/services/user"

	"github.com/gin-gonic/gin"
)

var mainRouter *gin.RouterGroup

func (s *Server) initRoutes() error {
	if configs.Config.App.Environment != "local" {
		gin.SetMode(gin.ReleaseMode)
	}
	s.initDomainRoutes()

	return nil
}

func (s *Server) initDomainRoutes() {
	mainRouter = s.router.Group("/api/v1/user-flow/")

	userRepository := userrepository.NewRepository(s.pgdb)
	userService := userservice.NewService(userRepository)
	userHandler := userhandler.NewHandler(userService)

	s.initUserRoutes(userHandler)
}

func (s *Server) initUserRoutes(handler *userhandler.Handler) {
	userRoutes := mainRouter.Group("user/")
	userRoutes.GET("show/:id", handler.Show)
}
